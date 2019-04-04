package models

import "database/sql"

type Accession struct {
	IdPeople     int  `json:"id_people"`
	EditAccess   bool `json:"edit_access"`
	SetAbsence   bool `json:"set_absence"`
	GetAbsence   bool `json:"get_absence"`
	SetMark      bool `json:"set_mark"`
	SetEvent     bool `json:"set_event"`
	GetSensitive bool `json:"get_sensitive"`
	SetSensitive bool `json:"set_sensitive"`
	GetYlist     bool `json:"get_ylist"`
	ManageAcadem bool `json:"manage_academ"`
}

func (ac *Accession) insert(tx *sql.Tx) error {
	_, err := tx.Exec("INSERT INTO accession(id_people, edit_access, set_absence, get_absence, set_mark, set_event, get_sensitive, set_sensitive, get_ylist, manage_academ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", ac.IdPeople, ac.EditAccess, ac.SetAbsence, ac.GetAbsence, ac.SetMark, ac.SetEvent, ac.GetSensitive, ac.SetSensitive, ac.GetYlist, ac.ManageAcadem)
	return err
}

func (ac *Accession) getById(db *sql.DB, id int) error {
	err := db.QueryRow("SELECT id_people, edit_access, set_absence, get_absence, set_mark, set_event, get_sensitive, set_sensitive, get_ylist, manage_academ FROM accession WHERE id_people = (?)", id).Scan(&ac.IdPeople, &ac.EditAccess, &ac.SetAbsence, &ac.GetAbsence, &ac.SetMark, &ac.SetEvent, &ac.GetSensitive, &ac.SetSensitive, &ac.GetYlist, &ac.ManageAcadem)
	return err
}
