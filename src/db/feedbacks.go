package db

import (
	"../blockchain"
	"database/sql"
)

type AvgFeedback struct {
	IdEmployee   int    `json:"id_employee"`
	NameEmployee string `json:"name_employee"`
	AvgMark      int    `json:"avg_mark"`
}

func GetAllFeedBacks(db *sql.DB) ([]blockchain.Block, error) {
	var feedbacks []blockchain.Block

	rows, err := db.Query("SELECT hash, prev_hash, mark, nonce, id_employee, data, timestamp FROM feedbacks ORDER BY timestamp DESC")
	if err != nil {
		return []blockchain.Block{}, err
	}

	for rows.Next() {
		var f blockchain.Block
		err = rows.Scan(&f.Hash, &f.PrevFeedBackHash, &f.Mark, &f.Nonce, &f.EmployeeId, &f.Data, &f.TimeStamp)
		if err != nil {
			return []blockchain.Block{}, err
		}
		feedbacks = append(feedbacks, f)
	}
	return feedbacks, nil
}

func GetFeedBackByEmployeeId(db *sql.DB, idEmployee int) ([]blockchain.Block, error) {
	var feedbacks []blockchain.Block

	rows, err := db.Query("SELECT hash, prev_hash, mark, nonce, id_employee, data, timestamp FROM feedbacks WHERE id_employee = ? ORDER BY timestamp DESC", idEmployee)
	if err != nil {
		return []blockchain.Block{}, err
	}

	for rows.Next() {
		var f blockchain.Block
		err = rows.Scan(&f.Hash, &f.PrevFeedBackHash, &f.Mark, &f.Nonce, &f.EmployeeId, &f.Data, &f.TimeStamp)
		if err != nil {
			return []blockchain.Block{}, err
		}
		feedbacks = append(feedbacks, f)
	}
	return feedbacks, nil
}

func GetAllFeedBackWithAVG(db *sql.DB) ([]AvgFeedback, error) {
	var avgfeedback []AvgFeedback

	rows, err := db.Query("SELECT P.fio, F.id_employee, AVG(F.mark) AS avg_mark FROM feedbacks F, people P WHERE P.id = F.id_employee GROUP BY F.id_employee ORDER BY avg_mark DESC, P.fio ASC;")
	if err != nil {
		return []AvgFeedback{}, err
	}

	for rows.Next() {
		var af AvgFeedback
		avg := sql.NullFloat64{}
		err = rows.Scan(&af.NameEmployee, &af.IdEmployee, &avg)
		af.AvgMark = int(avg.Float64)
		if err != nil {
			return []AvgFeedback{}, err
		}
		avgfeedback = append(avgfeedback, af)
	}
	return avgfeedback, nil
}

func GetAllFeedBackWithAVGByStudent(db *sql.DB, idGroup int) ([]AvgFeedback, error) {
	var avgfeedback []AvgFeedback

	rows, err := db.Query("SELECT P.fio, F.id_employee, AVG(F.mark) as avg_mark FROM feedbacks F, people P WHERE P.id = F.id_employee AND F.id_employee IN (SELECT id_employee FROM loads WHERE id_group = ?) GROUP BY F.id_employee ORDER BY avg_mark DESC, P.fio ASC", idGroup)
	if err != nil {
		return []AvgFeedback{}, err
	}

	for rows.Next() {
		var af AvgFeedback
		avg := sql.NullFloat64{}
		err = rows.Scan(&af.NameEmployee, &af.IdEmployee, &avg)
		af.AvgMark = int(avg.Float64)
		if err != nil {
			return []AvgFeedback{}, err
		}
		avgfeedback = append(avgfeedback, af)
	}
	return avgfeedback, nil
}
