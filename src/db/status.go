package db

import (
	"database/sql"
	"fmt"
)

type Status struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}


func (s *Status) getStatusById(db *sql.DB, id int) error{
	err := db.QueryRow("SELECT id, name FROM status WHERE id = ?", id).Scan(&s.Id, &s.Name)
	fmt.Printf("%v", s)
	return err
}
