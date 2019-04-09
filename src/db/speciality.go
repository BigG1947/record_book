package db

import "database/sql"

type Speciality struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	IdDirection int    `json:"id_direction"`
}

func (s *Speciality) GetSpecialityById(db *sql.DB, id int) error{
	err := db.QueryRow("SELECT id, name, id_direction FROM speciality WHERE id = (?)", id).Scan(&s.Id, &s.Name, &s.IdDirection)
	return err
}

func GetAllSpeciality(db *sql.DB) ([]Speciality, error) {
	var speciality []Speciality
	rows, err := db.Query("SELECT id, name, id_direction FROM speciality")
	if err != nil {
		return speciality, err
	}
	defer rows.Close()

	for rows.Next() {
		var s Speciality
		err := rows.Scan(&s.Id, &s.Name, &s.IdDirection)
		if err != nil {
			return speciality, err
		}
		speciality = append(speciality, s)
	}

	return speciality, nil
}