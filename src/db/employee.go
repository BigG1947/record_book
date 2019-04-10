package db

import "database/sql"

type Employee struct {
	IdPeople   int      `json:"id_people"`
	DateInvite string   `json:"date_invite"`
	Rank       Rank     `json:"rank"`
	Group      []Group  `json:"group"`
	Cathedra   Cathedra `json:"cathedra"`
}

func (empl *Employee) insert(tx *sql.Tx) error {
	_, err := tx.Exec("INSERT INTO employee(id_people, date_invite, id_rank, id_cathedra) VALUES (?, ?, ?, ?)", empl.IdPeople, empl.DateInvite, empl.Rank.Id, empl.Cathedra.Id)
	return err
}

func (empl *Employee) getById(db *sql.DB, id int) error {
	err := db.QueryRow("SELECT id_people, date_invite, id_rank, id_cathedra FROM employee WHERE id_people = (?)", id).Scan(&empl.IdPeople, &empl.DateInvite, &empl.Rank.Id, &empl.Cathedra.Id)
	if err != nil {
		return err
	}
	err = empl.Cathedra.GetCathedraById(db, empl.Cathedra.Id)
	if err != nil {
		return err
	}
	err = empl.Rank.GetRankById(db, empl.Rank.Id)
	if err != nil {
		return err
	}
	empl.Group, err = GetGroupByIdEmployee(db, empl.IdPeople)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	return nil
}
