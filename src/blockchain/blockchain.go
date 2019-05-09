package blockchain

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type BlockChain struct {
	Tip      []byte
	db       *sql.DB
	length   int
	Status   bool
	listNode *[]Node
}

type Iterator struct {
	currentHash []byte
	db          *sql.DB
}

func (bc *BlockChain) AddBlock(data string, employeeId int, mark int, timestamp int64) {
	prevFeedBackHash := bc.Tip
	newFeedBack := NewBlock(data, employeeId, mark, prevFeedBackHash, timestamp)

	_, err := bc.db.Exec("INSERT INTO feedbacks(hash, prev_hash, nonce, timestamp, data, id_employee, mark) VALUES (?,?,?,?,?,?,?)", newFeedBack.Hash, newFeedBack.PrevFeedBackHash, newFeedBack.Nonce, newFeedBack.TimeStamp, newFeedBack.Data, newFeedBack.EmployeeId, newFeedBack.Mark)
	if err != nil {
		log.Printf("Error in add block to blockchain: %s", err)
	}
	bc.Tip = newFeedBack.Hash
	bc.length++
}

func (bc *BlockChain) AddBlockWithOutSum(block *Block) error {
	_, err := bc.db.Exec("INSERT INTO feedbacks(hash, prev_hash, nonce, timestamp, data, id_employee, mark) VALUES (?,?,?,?,?,?,?)", block.Hash, block.PrevFeedBackHash, block.Nonce, block.TimeStamp, block.Data, block.EmployeeId, block.Mark)
	if err != nil {
		return err
	}
	bc.Tip = block.Hash
	bc.length++
	return nil
}

func (bc *BlockChain) Iterator() *Iterator {
	return &Iterator{bc.Tip, bc.db}
}

func (bci *Iterator) Next() (*Block, error) {
	var fb Block
	err := bci.db.QueryRow("SELECT hash, prev_hash, nonce, timestamp, data, id_employee, mark FROM feedbacks WHERE hash = ?", bci.currentHash).Scan(&fb.Hash, &fb.PrevFeedBackHash, &fb.Nonce, &fb.TimeStamp, &fb.Data, &fb.EmployeeId, &fb.Mark)
	bci.currentHash = fb.PrevFeedBackHash
	return &fb, err
}

func InitBlockChain(db *sql.DB) (*BlockChain, error) {
	var tip []byte
	var length int

	_, err := db.Exec("CREATE TABLE IF NOT EXISTS feedbacks(hash VARBINARY(256) PRIMARY KEY, prev_hash VARBINARY(256), nonce INT, timestamp INT, data TEXT, id_employee INT, mark INT);")
	if err != nil {
		return &BlockChain{}, err
	}

	err = db.QueryRow("SELECT hash, COUNT(hash) FROM feedbacks WHERE timestamp = (SELECT max(timestamp) FROM feedbacks);").Scan(&tip, &length)
	if err != nil {
		log.Printf("Error in initialization blockchain: %s", err)
	}
	if len(tip) == 0 {
		genesis := NewBlock("Genesis", -1, 0, []byte{}, 0)
		_, err := db.Exec("INSERT INTO feedbacks(hash, prev_hash, nonce, timestamp, data, id_employee, mark) VALUES (?,?,?,?,?,?,?)", genesis.Hash, genesis.PrevFeedBackHash, genesis.Nonce, genesis.TimeStamp, genesis.Data, genesis.EmployeeId, genesis.Mark)
		if err != nil {
			return &BlockChain{}, err
		}
		tip = genesis.Hash
		length = 1
	}

	return &BlockChain{tip, db, length, false, GetNodeList()}, nil
}

func (bc *BlockChain) FillTestFeedBack() {
	bc.AddBlock("Первый", 3, 100, 0)
	bc.AddBlock("Второй", 3, 87, 0)
	bc.AddBlock("Третий", 6, 100, 0)
	bc.AddBlock("Четвертый", 6, 34, 0)
	bc.AddBlock("Пятый", 9, 100, 0)
	bc.AddBlock("Шестой", 9, 56, 0)
	bc.AddBlock("Седьмой", 12, 12, 0)
}

func (bc *BlockChain) PrintBlockChain(bci *Iterator, w *http.ResponseWriter) {
	for {
		fb, err := bci.Next()
		if len(fb.Hash) == 0 || err != nil {
			break
		}
		fb.printFeedBack(w)
	}
}

func (bc *BlockChain) GetLength() int {
	return bc.length
}

func (bc *BlockChain) compareBlockWithOtherNodes(idEmployee int, mark int, text string, timestamp int64, hash []byte) (bool, error) {
	client := http.Client{}
	var applyNode int
	var data = struct {
		IdEmployee int    `json:"id_employee"`
		Mark       int    `json:"mark"`
		Text       string `json:"text"`
		TimeStamp  int64  `json:"time"`
	}{
		IdEmployee: idEmployee,
		Mark:       mark,
		Text:       text,
		TimeStamp:  timestamp,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return false, err
	}

	for _, node := range *bc.listNode {
		stringURL := fmt.Sprintf("http://%s:%d/blockchain/sumBlock", node.ip, node.port)
		req, err := http.NewRequest("POST", stringURL, bytes.NewBuffer(jsonData))
		if err != nil {
			return false, err
		}
		resp, err := client.Do(req)
		if err != nil {
			continue
		}
		if resp.StatusCode == http.StatusBadRequest {
			continue
		}
		if body, _ := ioutil.ReadAll(resp.Body); bytes.Equal(body, hash) {
			applyNode++
			continue
		}
	}
	return false, nil
}

func (bc *BlockChain) sendBlockToOtherNodes() (bool, error) {
	return false, nil
}

func (bc *BlockChain) CheckBlockChain(bci *Iterator) (bool, error) {
	for {
		fb, err := bci.Next()
		if len(fb.Hash) == 0 {
			break
		}
		if err != nil {
			return false, err
		}
		pow := NewProofOfWork(fb)
		if !pow.Validate() {
			return false, nil
		}
	}
	return true, nil
}
