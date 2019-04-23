package db

import (
	"../blockchain"
	"database/sql"
)

func GetAllFeedBacks(db *sql.DB) ([]blockchain.FeedBack, error) {
	var feedbacks []blockchain.FeedBack

	rows, err := db.Query("SELECT hash, prev_hash, mark, nonce, id_employee, data, timestamp FROM feedbacks ORDER BY timestamp DESC")
	if err != nil {
		return []blockchain.FeedBack{}, err
	}

	for rows.Next() {
		var f blockchain.FeedBack
		err = rows.Scan(&f.Hash, &f.PrevFeedBackHash, &f.Mark, &f.Nonce, &f.EmployeeId, &f.Data, &f.TimeStamp)
		if err != nil {
			return []blockchain.FeedBack{}, err
		}
		feedbacks = append(feedbacks, f)
	}
	return feedbacks, nil
}
