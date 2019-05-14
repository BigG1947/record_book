package db

import "database/sql"

type Direction struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	IdCathedra int    `json:"id_cathedra"`
}

func (d *Direction) GetDirectionById(db *sql.DB, id int) error {
	err := db.QueryRow("SELECT id, name, id_cathedra FROM directon WHERE id = ?", id).Scan(&d.Id, &d.Name, &d.IdCathedra)
	return err
}

func GetAllDirection(db *sql.DB) ([]Direction, error) {
	var directions []Direction
	rows, err := db.Query("SELECT id, name, id_cathedra FROM direction")
	if err != nil {
		return []Direction{}, err
	}

	for rows.Next() {
		var d Direction
		err = rows.Scan(&d.Id, &d.Name, &d.IdCathedra)
		if err != nil {
			return []Direction{}, nil
		}
		directions = append(directions, d)
	}
	return directions, nil
}

func GetAllDirectionByCathedra(db *sql.DB, idCathedra int) ([]Direction, error) {
	var directions []Direction
	rows, err := db.Query("SELECT id, name FROM direction WHERE id_cathedra = ?", idCathedra)
	if err != nil {
		return []Direction{}, err
	}

	for rows.Next() {
		var d Direction
		err = rows.Scan(&d.Id, &d.Name)
		d.IdCathedra = idCathedra
		if err != nil {
			return []Direction{}, nil
		}
		directions = append(directions, d)
	}
	return directions, nil
}

func GetAllDirectionByFaculty(db *sql.DB, idFaculty int) ([]Direction, error) {
	var directions []Direction
	rows, err := db.Query("SELECT direction.id, direction.name, direction.id_cathedra FROM direction, cathedra WHERE cathedra.id_faculty = ? AND direction.id_cathedra = cathedra.id", idFaculty)
	if err != nil {
		return []Direction{}, err
	}

	for rows.Next() {
		var d Direction
		err = rows.Scan(&d.Id, &d.Name, &d.IdCathedra)
		if err != nil {
			return []Direction{}, nil
		}
		directions = append(directions, d)
	}
	return directions, nil
}

func GetAllDirectionByInstitute(db *sql.DB, idInstitute int) ([]Direction, error) {
	var directions []Direction
	rows, err := db.Query("SELECT direction.id, direction.name, direction.id_cathedra FROM direction, cathedra, faculty WHERE faculty.id_institute = ? AND cathedra.id_faculty = faculty.id AND direction.id_cathedra = cathedra.id", idInstitute)
	if err != nil {
		return []Direction{}, err
	}

	for rows.Next() {
		var d Direction
		err = rows.Scan(&d.Id, &d.Name, &d.IdCathedra)
		if err != nil {
			return []Direction{}, nil
		}
		directions = append(directions, d)
	}
	return directions, nil
}

func UpdateDirection(db *sql.DB, d *Direction) error {
	_, err := db.Exec(updateDirectionScript, d.Name, d.IdCathedra, d.Id)
	if err != nil {
		return err
	}
	return nil
}

func DeleteDirection(db *sql.DB, id int) error {
	_, err := db.Exec(deleteDirectionScript, id)
	if err != nil {
		return err
	}
	return nil
}

func InsertDirection(db *sql.DB, d *Direction) (int, error) {
	res, err := db.Exec(insertDirectionScript, d.Name, d.IdCathedra)
	if err != nil {
		return 0, err
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(lastId), nil
}
