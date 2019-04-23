package blockchain

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
)

type BlockChain struct {
	Tip []byte
	db  *sql.DB
}

type Iterator struct {
	currentHash []byte
	db          *sql.DB
}

func (bc *BlockChain) AddFeedBack(data string, employeeId int, mark int) {
	prevFeedBackHash := bc.Tip
	newFeedBack := NewFeedBack(data, employeeId, mark, prevFeedBackHash)
	bc.db.Exec("INSERT INTO feedbacks(hash, prev_hash, nonce, timestamp, data, id_employee, mark) VALUES (?,?,?,?,?,?,?)", newFeedBack.Hash, newFeedBack.PrevFeedBackHash, newFeedBack.Nonce, newFeedBack.TimeStamp, newFeedBack.Data, newFeedBack.EmployeeId, newFeedBack.Mark)
	bc.Tip = newFeedBack.Hash
}

func (bc *BlockChain) Iterator() *Iterator {
	return &Iterator{bc.Tip, bc.db}
}

func (bci *Iterator) Next() (*FeedBack, error) {
	var fb FeedBack
	err := bci.db.QueryRow("SELECT hash, prev_hash, nonce, timestamp, data, id_employee, mark FROM feedbacks WHERE hash = ?", bci.currentHash).Scan(&fb.Hash, &fb.PrevFeedBackHash, &fb.Nonce, &fb.TimeStamp, &fb.Data, &fb.EmployeeId, &fb.Mark)
	bci.currentHash = fb.PrevFeedBackHash
	return &fb, err
}

func InitBlockChain(db *sql.DB) (*BlockChain, error) {
	var tip []byte

	_, err := db.Exec("CREATE TABLE IF NOT EXISTS feedbacks(hash VARBINARY(256) PRIMARY KEY, prev_hash VARBINARY(256), nonce INT, timestamp INT, data TEXT, id_employee INT, mark INT);")
	if err != nil {
		errLogs(err)
		return &BlockChain{}, err
	}
	var timestamp int
	db.QueryRow("SELECT hash, timestamp FROM feedbacks WHERE timestamp = (SELECT max(timestamp) FROM blockchain);").Scan(&tip, &timestamp)
	if len(tip) == 0 {
		genesis := NewFeedBack("Genesis", -1, 0, []byte{})
		_, err := db.Exec("INSERT INTO feedbacks(hash, prev_hash, nonce, timestamp, data, id_employee, mark) VALUES (?,?,?,?,?,?,?)", genesis.Hash, genesis.PrevFeedBackHash, genesis.Nonce, genesis.TimeStamp, genesis.Data, genesis.EmployeeId, genesis.Mark)
		if err != nil {
			errLogs(err)
			return &BlockChain{}, err
		}
		tip = genesis.Hash
	}

	return &BlockChain{tip, db}, nil
}

func errLogs(err error) {
	fmt.Printf("err: %s\n", err)
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
