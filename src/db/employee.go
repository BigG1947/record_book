package db

import "database/sql"

type Employee struct {
	IdPeople   int      `json:"id_people"`
	DateInvite string   `json:"date_invite"`
	Rank       Rank     `json:"rank"`
	Group      Group    `json:"group"`
	Cathedra   Cathedra `json:"cathedra"`
}

func (empl *Employee) insert(tx *sql.Tx) error {
	_, err := tx.Exec("INSERT INTO employee(id_people, date_invite, id_rank, id_group, id_cathedra) VALUES (?, ?, ?, ?, ?)", empl.IdPeople, empl.DateInvite, empl.Rank.Id, checkForeignKey(empl.Group.Id), empl.Cathedra.Id)
	return err
}

func (empl *Employee) getById(db *sql.DB, id int) error {
	var idGroup sql.NullInt64
	err := db.QueryRow("SELECT id_people, date_invite, id_rank, id_group, id_cathedra FROM employee WHERE id_people = (?)", id).Scan(&empl.IdPeople, &empl.DateInvite, &empl.Rank.Id, &idGroup, &empl.Cathedra.Id)
	empl.Group.Id = int(idGroup.Int64)
	if err != nil {
		return err
	}
	err = empl.Cathedra.GetCathedraById(db, empl.Cathedra.Id)
	if err != nil{
		return err
	}
	err = empl.Rank.GetRankById(db, empl.Rank.Id)
	if err != nil{
		return err
	}
	err = empl.Group.GetGroupById(db, 0)
	if err != nil && err != sql.ErrNoRows{
		return err
	}
	return nil
}
