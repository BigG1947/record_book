package db

import "database/sql"

type Cathedra struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	IdFaculty int    `json:"id_faculty"`
}

func (c *Cathedra) GetCathedraById(db *sql.DB, id int) error {
	err := db.QueryRow("SELECT id, name, id_faculty FROM cathedra WHERE id = ?", id).Scan(&c.Id, &c.Name, &c.IdFaculty)
	return err
}

func GetAllCathedra(db *sql.DB) ([]Cathedra, error) {
	var cathedra []Cathedra
	rows, err := db.Query("SELECT id, name, id_faculty FROM cathedra")
	if err != nil {
		return []Cathedra{}, err
	}

	for rows.Next() {
		var c Cathedra
		err = rows.Scan(&c.Id, &c.Name, &c.IdFaculty)
		if err != nil {
			return []Cathedra{}, nil
		}
		cathedra = append(cathedra, c)
	}
	return cathedra, nil
}
func GetCathedraByFaculty(db *sql.DB, idFaculty int) ([]Cathedra, error) {
	var cathedra []Cathedra
	rows, err := db.Query("SELECT id, name FROM cathedra WHERE id_faculty = ?", idFaculty)
	if err != nil {
		return []Cathedra{}, err
	}

	for rows.Next() {
		var c Cathedra
		err = rows.Scan(&c.Id, &c.Name)
		c.IdFaculty = idFaculty
		if err != nil {
			return []Cathedra{}, nil
		}
		cathedra = append(cathedra, c)
	}
	return cathedra, nil
}

func GetCathedraByInstitute(db *sql.DB, idInstitute int) ([]Cathedra, error) {
	var cathedra []Cathedra
	rows, err := db.Query("SELECT cathedra.id, cathedra.name, cathedra.id_faculty FROM cathedra, faculty WHERE faculty.id_institute = ? AND cathedra.id_faculty = faculty.id", idInstitute)
	if err != nil {
		return []Cathedra{}, err
	}

	for rows.Next() {
		var c Cathedra
		err = rows.Scan(&c.Id, &c.Name, &c.IdFaculty)
		if err != nil {
			return []Cathedra{}, nil
		}
		cathedra = append(cathedra, c)
	}
	return cathedra, nil
}

func UpdateCathedra(db *sql.DB, c *Cathedra) error {
	_, err := db.Exec(updateCathedraScript, c.Name, c.IdFaculty, c.Id)
	if err != nil {
		return err
	}
	return nil
}

func DeleteCathedra(db *sql.DB, id int) error {
	_, err := db.Exec(deleteCathedraScript, id)
	if err != nil {
		return err
	}
	return nil
}

func InsertCathedra(db *sql.DB, c *Cathedra) (int, error) {
	res, err := db.Exec(insertCathedraScript, c.Name, c.IdFaculty)
	if err != nil {
		return 0, err
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(lastId), nil
}
