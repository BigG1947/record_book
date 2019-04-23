package blockchain

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type FeedBack struct {
	TimeStamp        int64
	Data             []byte
	EmployeeId       int
	PrevFeedBackHash []byte
	Hash             []byte
	Nonce            int
	Mark             int
}

func NewFeedBack(data string, idEmployee int, mark int, prevFeedBackHash []byte) *FeedBack {
	feedback := &FeedBack{time.Now().Unix(), []byte(data), idEmployee, prevFeedBackHash, []byte{}, 0, mark}
	pow := NewProofOfWork(feedback)
	nonce, hash := pow.Run()

	feedback.Hash = hash
	feedback.Nonce = nonce

	return feedback
}

func GetAllFeedBacks(db *sql.DB) ([]FeedBack, error) {
	var feedbacks []FeedBack

	rows, err := db.Query("SELECT hash, prev_hash, mark, nonce, id_employee, data, timestamp FROM feedbacks ORDER BY timestamp DESC")
	if err != nil {
		return []FeedBack{}, err
	}

	for rows.Next() {
		var f FeedBack
		err = rows.Scan(&f.Hash, &f.PrevFeedBackHash, &f.Mark, &f.Nonce, &f.EmployeeId, &f.Data, &f.TimeStamp)
		if err != nil {
			return []FeedBack{}, err
		}
		feedbacks = append(feedbacks, f)
	}
	return feedbacks, nil
}

func (feedback *FeedBack) printFeedBack(w *http.ResponseWriter) {
	fmt.Fprintf(*w, "timestamp: %v\n", time.Unix(feedback.TimeStamp, 0))
	fmt.Fprintf(*w, "data: %s\n", feedback.Data)
	fmt.Fprintf(*w, "employeeId: %d\n", feedback.EmployeeId)
	fmt.Fprintf(*w, "hash: %x\n", feedback.Hash)
	fmt.Fprintf(*w, "prev_hash: %x\n", feedback.PrevFeedBackHash)
	fmt.Fprintf(*w, "mark: %d\n", feedback.Mark)
	pow := NewProofOfWork(feedback)
	fmt.Fprintf(*w, "pow: %s\n", strconv.FormatBool(pow.Validate()))
	fmt.Fprintf(*w, "\n")
}
