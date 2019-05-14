package db

import (
	"database/sql"
)

const StatusStudy = 1
const StatusDeducted = 2
const StatusWork = 3
const StatusStudentVacation = 4
const StatusVacation = 5
const StatusMaternityLeave = 6
const StatusEndEducation = 7
const StatusBusinesTrip = 8
const StatusFired = 9

type Status struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (s *Status) getStatusById(db *sql.DB, id int) error {
	err := db.QueryRow("SELECT id, name FROM status WHERE id = ?", id).Scan(&s.Id, &s.Name)
	return err
}

func GetAllStatus(db *sql.DB) ([]Status, error) {
	var status []Status
	rows, err := db.Query("SELECT id, name FROM status")
	if err != nil {
		return status, err
	}
	defer rows.Close()

	for rows.Next() {
		var s Status
		err := rows.Scan(&s.Id, &s.Name)
		if err != nil {
			return status, err
		}
		status = append(status, s)
	}

	return status, nil
}
