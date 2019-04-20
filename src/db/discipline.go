package db

import "database/sql"

type Discipline struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (d *Discipline) getDisciplineById(db *sql.DB, id int) error {
	err := db.QueryRow("SELECT id, name FROM discipline WHERE id = ?", id).Scan(&d.Id, &d.Name)
	return err
}

func GetAllDiscipline(db *sql.DB) ([]Discipline, error) {
	var disciplines []Discipline

	rows, err := db.Query("SELECT id, name FROM discipline")
	if err != nil {
		return []Discipline{}, err
	}

	for rows.Next() {
		var d Discipline
		err = rows.Scan(&d.Id, &d.Name)
		if err != nil {
			return []Discipline{}, err
		}
		disciplines = append(disciplines, d)
	}
	return disciplines, nil
}

func GetAllDisciplineForEmployee(db *sql.DB, id int) ([]Discipline, error) {
	var disciplines []Discipline

	rows, err := db.Query("SELECT D.id, D.name FROM discipline D WHERE D.id IN (SELECT id_discipline FROM loads WHERE id_employee = ?)", id)
	if err != nil {
		return []Discipline{}, err
	}

	for rows.Next() {
		var d Discipline
		err = rows.Scan(&d.Id, &d.Name)
		if err != nil {
			return []Discipline{}, err
		}
		disciplines = append(disciplines, d)
	}
	return disciplines, nil
}

func (d *Discipline) Insert(db *sql.DB) (int, error) {
	res, err := db.Exec("INSERT INTO discipline(name) VALUES (?)", d.Name)
	if err != nil {
		return 0, err
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		return int(lastId), err
	}
	return int(lastId), nil
}

func GetStudentDisceplines(db *sql.DB, idGroup int) ([]Discipline, error) {
	var disciplines []Discipline

	rows, err := db.Query("SELECT id, name FROM discipline WHERE id IN (SELECT id_discipline FROM loads WHERE id_group = ?)", idGroup)
	if err != nil {
		return []Discipline{}, err
	}

	for rows.Next() {
		var d Discipline
		err = rows.Scan(&d.Id, &d.Name)
		if err != nil {
			return []Discipline{}, err
		}
		disciplines = append(disciplines, d)
	}
	return disciplines, nil
}

func GetStudentDisceplinesBySemester(db *sql.DB, idGroup int, semester int) ([]Discipline, error) {
	var disciplines []Discipline

	rows, err := db.Query("SELECT id, name FROM discipline WHERE id IN (SELECT id_discipline FROM loads WHERE id_group = ? AND semester = ?)", idGroup, semester)
	if err != nil {
		return []Discipline{}, err
	}

	for rows.Next() {
		var d Discipline
		err = rows.Scan(&d.Id, &d.Name)
		if err != nil {
			return []Discipline{}, err
		}
		disciplines = append(disciplines, d)
	}
	return disciplines, nil
}
