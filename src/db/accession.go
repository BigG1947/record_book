package db

import (
	"database/sql"
	"errors"
)

const (
	SET_ACCESS    string = "edit_access"
	SET_SKIP      string = "set_absence"
	GET_SKIP      string = "get_absence"
	SET_MARK      string = "set_mark"
	SET_EVENT     string = "set_event"
	GET_SENSITIVE string = "get_sensitive"
	SET_SENSITIVE string = "set_sensitive"
	MANAGE        string = "menage_academ"
	MANAGE_LOAD   string = "manage_load"
)

type Accession struct {
	IdPeople     int  `json:"id_people"`
	EditAccess   bool `json:"edit_access"`
	SetAbsence   bool `json:"set_absence"`
	GetAbsence   bool `json:"get_absence"`
	SetMark      bool `json:"set_mark"`
	SetEvent     bool `json:"set_event"`
	GetSensitive bool `json:"get_sensitive"`
	SetSensitive bool `json:"set_sensitive"`
	ManageLoad   bool `json:"manage_load"`
	ManageAcadem bool `json:"manage_academ"`
}

func (ac *Accession) insert(tx *sql.Tx) error {
	_, err := tx.Exec("INSERT INTO accession(id_people, edit_access, set_absence, get_absence, set_mark, set_event, get_sensitive, set_sensitive, manage_load, manage_academ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", ac.IdPeople, ac.EditAccess, ac.SetAbsence, ac.GetAbsence, ac.SetMark, ac.SetEvent, ac.GetSensitive, ac.SetSensitive, ac.ManageLoad, ac.ManageAcadem)
	return err
}

func (ac *Accession) getById(db *sql.DB, id int) error {
	err := db.QueryRow("SELECT id_people, edit_access, set_absence, get_absence, set_mark, set_event, get_sensitive, set_sensitive, manage_load, manage_academ FROM accession WHERE id_people = (?)", id).Scan(&ac.IdPeople, &ac.EditAccess, &ac.SetAbsence, &ac.GetAbsence, &ac.SetMark, &ac.SetEvent, &ac.GetSensitive, &ac.SetSensitive, &ac.ManageLoad, &ac.ManageAcadem)
	return err
}

func (ac *Accession) Update(db *sql.DB) error {
	_, err := db.Exec(updateAccessionDataScript, ac.EditAccess, ac.SetAbsence, ac.GetAbsence, ac.SetMark, ac.SetEvent, ac.GetSensitive, ac.SetSensitive, ac.ManageLoad, ac.ManageAcadem, ac.IdPeople)
	if err != nil {
		return err
	}
	return nil
}

func UpdateAccession(db *sql.DB, ac *Accession) error {
	_, err := db.Exec(updateAccessionDataScript, ac.EditAccess, ac.SetAbsence, ac.GetAbsence, ac.SetMark, ac.SetEvent, ac.GetSensitive, ac.SetSensitive, ac.ManageLoad, ac.ManageAcadem, ac.IdPeople)
	if err != nil {
		return err
	}
	return nil
}

func (ac *Accession) getAccession() (result map[string]bool) {
	result = make(map[string]bool)
	result[GET_SENSITIVE] = ac.GetSensitive
	result[SET_SENSITIVE] = ac.SetSensitive
	result[GET_SKIP] = ac.GetAbsence
	result[SET_SKIP] = ac.SetAbsence
	result[MANAGE_LOAD] = ac.ManageLoad
	result[SET_ACCESS] = ac.EditAccess
	result[SET_EVENT] = ac.SetEvent
	result[SET_MARK] = ac.SetMark
	result[MANAGE] = ac.ManageAcadem
	return
}

func (ac *Accession) setAccession(acm map[string]bool) error {
	err := errors.New("not found value in map")
	if b, ok := acm[GET_SENSITIVE]; ok {
		ac.GetSensitive = b
	} else {
		return err
	}
	if b, ok := acm[SET_SENSITIVE]; ok {
		ac.SetSensitive = b
	} else {
		return err
	}
	if b, ok := acm[GET_SKIP]; ok {
		ac.GetAbsence = b
	} else {
		return err
	}

	if b, ok := acm[SET_MARK]; ok {
		ac.SetMark = b
	} else {
		return err
	}
	if b, ok := acm[SET_SKIP]; ok {
		ac.SetAbsence = b
	} else {
		return err
	}
	if b, ok := acm[SET_EVENT]; ok {
		ac.SetEvent = b
	} else {
		return err
	}
	if b, ok := acm[SET_ACCESS]; ok {
		ac.EditAccess = b
	} else {
		return err
	}
	if b, ok := acm[MANAGE]; ok {
		ac.ManageAcadem = b
	} else {
		return err
	}
	if b, ok := acm[MANAGE_LOAD]; ok {
		ac.ManageLoad = b
	} else {
		return err
	}

	return nil
}
