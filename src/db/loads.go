package db

import "database/sql"

type Load struct {
	Id           int `json:"id"`
	IdDiscipline int `json:"id_discipline"`
	IdEmployee   int `json:"id_employee"`
	IdGroup      int `json:"id_group"`
	Semester     int `json:"semester"`
}

func (l *Load) Insert(db *sql.DB) (int, error) {
	res, err := db.Exec("INSERT INTO loads(id_discipline, id_employee, id_group, semester) VALUES (?, ?, ?, ?, ?)", l.IdDiscipline, l.IdEmployee, l.IdGroup, l.Semester)
	if err != nil {
		return 0, err
	}

	lastId, _ := res.LastInsertId()
	return int(lastId), nil
}

func GetAllLoadsForEmployee(db *sql.DB, id int) ([]Load, error) {
	var loads []Load
	rows, err := db.Query("SELECT L.id, L.id_discipline, L.id_employee, L.id_group, L.semester FROM loads L WHERE id_employee = ?", id)
	if err != nil {
		return []Load{}, err
	}

	for rows.Next() {
		var l Load
		err = rows.Scan(&l.Id, &l.IdDiscipline, &l.IdEmployee, &l.IdGroup, &l.Semester)
		if err != nil {
			return []Load{}, err
		}
	}
	return loads, nil
}
