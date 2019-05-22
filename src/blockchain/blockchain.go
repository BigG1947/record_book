package blockchain

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type BlockChain struct {
	Tip      []byte
	db       *sql.DB
	length   int
	Status   bool
	listNode []Node
	hash     []byte
}

type Iterator struct {
	currentHash []byte
	db          *sql.DB
}

func (bc *BlockChain) AddBlock(data string, employeeId int, mark int, timestamp int64) error {
	var err error
	//if !bc.Status {
	//	return errors.New("BlockChain status if fall ")
	//}
	//ok, err := CheckNodesLive(bc.listNode)
	//if err != nil {
	//	return err
	//}
	//if !ok {
	//	bc.Status = false
	//	return errors.New("Nodes in network < 3 ")
	//}
	prevFeedBackHash := bc.Tip
	newFeedBack := NewBlock(data, employeeId, mark, prevFeedBackHash, timestamp)
	//ok, err = bc.compareBlockWithOtherNodes(employeeId, mark, data, newFeedBack.TimeStamp, newFeedBack.Hash)
	//if err != nil {
	//	return err
	//}
	//if !ok {
	//	return errors.New("Block Hash not match to Hash in other nodes ")
	//}
	//ok, err = bc.sendBlockToOtherNodes(newFeedBack)
	//if err != nil{
	//	return err
	//}
	//if !ok {
	//	return errors.New("Node returns a BadRequest header ")
	//}
	_, err = bc.db.Exec("INSERT INTO feedbacks(hash, prev_hash, nonce, timestamp, data, id_employee, mark) VALUES (?,?,?,?,?,?,?)", newFeedBack.Hash, newFeedBack.PrevFeedBackHash, newFeedBack.Nonce, newFeedBack.TimeStamp, newFeedBack.Data, newFeedBack.EmployeeId, newFeedBack.Mark)
	if err != nil {
		return err
	}
	bc.Tip = newFeedBack.Hash
	bc.length++
	return nil
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
		genesis := NewBlock("Genesis", -1, 0, []byte{}, 1)
		_, err := db.Exec("INSERT INTO feedbacks(hash, prev_hash, nonce, timestamp, data, id_employee, mark) VALUES (?,?,?,?,?,?,?)", genesis.Hash, genesis.PrevFeedBackHash, genesis.Nonce, genesis.TimeStamp, genesis.Data, genesis.EmployeeId, genesis.Mark)
		if err != nil {
			return &BlockChain{}, err
		}
		tip = genesis.Hash
		length = 1
	}
	bc := BlockChain{tip, db, length, false, GetNodeList(), []byte{}}
	bc.hash, err = bc.GetFinalHash()
	if err != nil {
		return &BlockChain{}, err
	}
	//go bc.CheckNodesInNetWork()
	//go bc.CheckBlockChain()
	return &bc, nil
}

func (bc *BlockChain) FillTestFeedBack() {
	bc.AddBlock("Первый", 12, 100, 0)
	bc.AddBlock("Второй", 13, 50, 0)
	bc.AddBlock("Третий", 6, 99, 0)
	bc.AddBlock("Четвертый", 6, 100, 0)
	bc.AddBlock("Пятый", 9, 100, 0)
	bc.AddBlock("Шестой", 9, 99, 0)
	bc.AddBlock("Седьмой", 12, 100, 0)
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

func (bc *BlockChain) GetNodeList() []Node {
	return bc.listNode
}

func (bc *BlockChain) GetFinalHash() ([]byte, error) {
	var finalHash []byte
	bci := bc.Iterator()
	for {
		fb, err := bci.Next()
		if len(fb.Hash) == 0 {
			break
		}
		if err != nil {
			return []byte{}, err
		}

		finalHash = bytes.Join([][]byte{finalHash, fb.Hash}, []byte{})
	}
	return finalHash, nil
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
	for _, node := range bc.listNode {
		if !node.status {
			continue
		}
		stringURL := fmt.Sprintf("http://%s:%d/blockchain/sumBlock", node.ip, node.port)
		req, err := http.NewRequest("POST", stringURL, bytes.NewBuffer(jsonData))
		if err != nil {
			return false, err
		}
		resp, err := client.Do(req)
		if err != nil {
			return false, err
		}
		if resp.StatusCode == http.StatusBadRequest {
			fmt.Printf("BadRequest")
			return false, nil
		}
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Printf("Body request: %x\n", body)
		fmt.Printf("Block hash: %x\n", hash)
		fmt.Printf("%t\n", bytes.Equal(body, hash))
		if bytes.Equal(body, hash) {
			applyNode++
			continue
		}
	}
	if applyNode == len(bc.listNode) {
		return true, nil
	}
	return false, nil
}

func (bc *BlockChain) sendBlockToOtherNodes(block *Block) (bool, error) {
	client := http.Client{}
	jsonBlock, err := json.Marshal(block)
	if err != nil {
		return false, err
	}
	for _, node := range bc.listNode {
		if !node.status {
			continue
		}
		stringURL := fmt.Sprintf("http://%s:%d/blockchain/addBlockWithoutSum", node.ip, node.port)
		req, err := http.NewRequest("POST", stringURL, bytes.NewBuffer(jsonBlock))
		if err != nil {
			return false, err
		}
		resp, err := client.Do(req)
		if err != nil {
			return false, err
		}
		if resp.StatusCode == http.StatusBadRequest {
			fmt.Printf("BadRequest")
			return false, nil
		}
		if resp.StatusCode == http.StatusOK {
			return true, nil
		}
	}
	return false, nil
}

func (bc *BlockChain) CheckNodesInNetWork() {
	for true {
		for true {
			var err error
			bc.Status, err = CheckNodesLive(bc.listNode)
			if err != nil {
				log.Printf("Error in CheckNodesLive:\n%s\n", err)
				return
			}
			if bc.Status == true {
				break
			}
			time.Sleep(10 * time.Second)
		}
		log.Printf("BlockChain is work!")
		time.Sleep(60 * time.Second)
	}
}

func (bc *BlockChain) CheckBlockChain() {
	for true {
		var finalHash []byte
		var ok = true
		bci := bc.Iterator()
		for true {
			fb, err := bci.Next()
			if len(fb.Hash) == 0 {
				break
			}
			if err != nil {
				log.Printf("CheckBlockChain function: %s\n", err)
			}
			pow := NewProofOfWork(fb)
			if !pow.Validate() {
				log.Printf("Broken Block:\nHash: %x\nTimestamp: %d\n", fb.Hash, fb.TimeStamp)
				ok = false
				break
			}
			finalHash = bytes.Join([][]byte{finalHash, fb.Hash}, []byte{})
		}
		if ok {
			log.Printf("BlockChain is Whole!\n")
			bc.hash = finalHash
		} else {
			log.Printf("BlockChain is Broken\n")
		}
		time.Sleep(5 * time.Minute)
	}
}

func (bc *BlockChain) repairBlockChain() {
	log.Printf("Start repair blockchain...\n")
	//lenghtMap := make(map[int][]Node)
}
