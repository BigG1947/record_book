package models

import "database/sql"

type Employee struct {
	IdPeople   int    `json:"id_people"`
	DateInvite string `json:"date_invite"`
	IdRank     int    `json:"id_rank"`
	IdGroup    int    `json:"id_group"`
	IdCathedra int    `json:"id_cathedra"`
}

func (empl *Employee) insert(tx *sql.Tx) error {
	_, err := tx.Exec("INSERT INTO employee(id_people, date_invite, id_rank, id_group, id_cathedra) VALUES (?, ?, ?, ?, ?)", empl.IdPeople, empl.DateInvite, empl.IdRank, checkForeignKey(empl.IdGroup), empl.IdCathedra)
	return err
}

func (empl *Employee) getById(db *sql.DB, id int) error {
	var idGroup sql.NullInt64
	err := db.QueryRow("SELECT id_people, date_invite, id_rank, id_group, id_cathedra FROM employee WHERE id_people = (?)", id).Scan(&empl.IdPeople, &empl.DateInvite, &empl.IdRank, &idGroup, &empl.IdCathedra)
	empl.IdGroup = int(idGroup.Int64)
	return err
}
