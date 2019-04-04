package models

import "database/sql"

type SensitiveData struct {
	IdPeople     int    `json:"id_people"`
	PassportCode string `json:"passport_code"`
	Rntrs        string `json:"rntrs"`
	RegAddress   string `json:"reg_address"`
	MilitaryId   string `json:"military_id"`
}

func (sd *SensitiveData) insert(tx *sql.Tx) error  {
	_, err := tx.Exec("INSERT INTO sensitive_data(id_people, passport_code, rntrs, reg_address, military_id) VALUES (?, ?, ?, ?, ?)", sd.IdPeople, sd.PassportCode, sd.Rntrs, sd.RegAddress, sd.MilitaryId)
	return err
}

func (sd *SensitiveData) getById(db *sql.DB, id int) error{
	err := db.QueryRow("SELECT id_people, passport_code, rntrs, reg_address, military_id FROM sensitive_data WHERE id_people = (?)", id).Scan(&sd.IdPeople, &sd.PassportCode, &sd.Rntrs, &sd.RegAddress, &sd.MilitaryId)
	return err
}