package db

import (
	"database/sql"
	"time"
)

type Discipline struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type StudentDiscipline struct {
	Id            int          `json:"id"`
	Discipline    Discipline   `json:"id_discipline"`
	IdEmployee    int          `json:"id_employee"`
	NameEmployee  string       `json:"name_employee"`
	IdGroup       int          `json:"id_group"`
	NameGroup     string       `json:"name_group"`
	IdAssistant   int          `json:"id_assistant"`
	NameAssistant string       `json:"name_assistant"`
	Semester      LoadSemester `json:"semester"`
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

func GetCurrentStudentDiscipline(db *sql.DB, idGroup int) ([]Load, error) {
	var loads []Load
	currentTime := time.Now().Unix()
	rows, err := db.Query(getCurrentStudentDisciplineScript, idGroup, currentTime, currentTime)
	if err != nil {
		return []Load{}, err
	}

	for rows.Next() {
		var l Load
		err = rows.Scan(&l.Id, &l.Discipline.Id, &l.Discipline.Name, &l.IdEmployee, &l.NameEmployee, &l.IdGroup, &l.NameGroup, &l.IdAssistant, &l.NameAssistant, &l.Semester.Id, &l.Semester.Start, &l.Semester.End, &l.Semester.Name, &l.NumSemester)
		if err != nil {
			return []Load{}, err
		}
		loads = append(loads, l)
	}
	return loads, nil
}

func GetAllDisciplineForEmployee(db *sql.DB, id int) ([]Discipline, error) {
	var disciplines []Discipline
	currentTime := time.Now().Unix()
	rows, err := db.Query(getAllDisciplineForEmployeeScripts, id, currentTime, currentTime)
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
