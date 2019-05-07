package db

import "database/sql"

type Speciality struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	IdDirection int    `json:"id_direction"`
}

func (s *Speciality) GetSpecialityById(db *sql.DB, id int) error {
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

func GetAllSpecialityByDirection(db *sql.DB, idDirection int) ([]Speciality, error) {
	var speciality []Speciality
	rows, err := db.Query("SELECT id, name FROM speciality WHERE id_direction = ?", idDirection)
	if err != nil {
		return speciality, err
	}
	defer rows.Close()

	for rows.Next() {
		var s Speciality
		err := rows.Scan(&s.Id, &s.Name)
		if err != nil {
			return speciality, err
		}
		s.IdDirection = idDirection
		speciality = append(speciality, s)
	}

	return speciality, nil
}

func GetAllSpecialityByCathedra(db *sql.DB, idCathedra int) ([]Speciality, error) {
	var speciality []Speciality
	rows, err := db.Query("SELECT speciality.id, speciality.name, speciality.id_direction FROM speciality, direction WHERE direction.id_cathedra = ? AND speciality.id_direction = direction.id", idCathedra)
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

func GetAllSpecialityByFaculty(db *sql.DB, idFaculty int) ([]Speciality, error) {
	var speciality []Speciality
	rows, err := db.Query("SELECT speciality.id, speciality.name, speciality.id_direction FROM speciality, direction, cathedra WHERE cathedra.id_faculty = ? AND direction.id_cathedra = cathedra.id AND speciality.id_direction = direction.id", idFaculty)
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

func GetAllSpecialityByInstitute(db *sql.DB, idInstitute int) ([]Speciality, error) {
	var speciality []Speciality
	rows, err := db.Query("SELECT speciality.id, speciality.name, speciality.id_direction FROM speciality, direction, cathedra, faculty WHERE faculty.id_institute = ? AND cathedra.id_faculty = faculty.id AND direction.id_cathedra = cathedra.id AND speciality.id_direction = direction.id", idInstitute)
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
