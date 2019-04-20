package db

import "database/sql"

type Load struct {
	Id         int        `json:"id"`
	Discipline Discipline `json:"id_discipline"`
	Employee   People     `json:"id_employee"`
	Group      Group      `json:"id_group"`
	Assistant  People     `json:"id_assistant"`
	Semester   int        `json:"semester"`
}

func (l *Load) Insert(db *sql.DB) (int, error) {
	res, err := db.Exec("INSERT INTO loads(id_discipline, id_employee, id_group, id_assistant, semester) VALUES (?, ?, ?, ?, ?, ?)", l.Discipline.Id, l.Employee.Id, l.Group.Id, l.Assistant.Id, l.Semester)
	if err != nil {
		return 0, err
	}

	lastId, _ := res.LastInsertId()
	return int(lastId), nil
}

func GetAllLoadsForEmployee(db *sql.DB, id int) ([]Load, error) {
	var loads []Load
	rows, err := db.Query("SELECT L.id, L.id_discipline, L.id_employee, L.id_group, L.id_assistant, L.semester FROM loads L WHERE id_employee = ?", id)
	if err != nil {
		return []Load{}, err
	}

	for rows.Next() {
		var l Load
		err = rows.Scan(&l.Id, &l.Discipline.Id, &l.Employee.Id, &l.Group.Id, &l.Assistant.Id, &l.Semester)
		if err != nil {
			return []Load{}, err
		}
		err = l.Discipline.getDisciplineById(db, l.Discipline.Id)
		if err != nil {
			return []Load{}, err
		}
		err = l.Group.GetGroupById(db, l.Group.Id)
		if err != nil {
			return []Load{}, err
		}
		err = l.Employee.GetPeopleById(db, l.Employee.Id)
		if err != nil {
			return []Load{}, err
		}
		err = l.Assistant.GetPeopleById(db, l.Assistant.Id)
		if err != nil {
			return []Load{}, err
		}
		loads = append(loads, l)
	}
	return loads, nil
}

func GetAllLoadsByIdGroup(db *sql.DB, id int) ([]Load, error) {
	var loads []Load
	rows, err := db.Query("SELECT L.id, L.id_discipline, L.id_employee, L.id_group, L.id_assistant, L.semester FROM loads L WHERE L.id_group = ? ORDER BY semester ASC ", id)
	if err != nil {
		return []Load{}, err
	}

	for rows.Next() {
		var l Load
		err = rows.Scan(&l.Id, &l.Discipline.Id, &l.Employee.Id, &l.Group.Id, &l.Assistant.Id, &l.Semester)
		if err != nil {
			return []Load{}, err
		}
		err = l.Discipline.getDisciplineById(db, l.Discipline.Id)
		if err != nil {
			return []Load{}, err
		}
		err = l.Group.GetGroupById(db, l.Group.Id)
		if err != nil {
			return []Load{}, err
		}
		err = l.Employee.GetPeopleById(db, l.Employee.Id)
		if err != nil {
			return []Load{}, err
		}
		err = l.Assistant.GetPeopleById(db, l.Assistant.Id)
		if err != nil {
			return []Load{}, err
		}
		loads = append(loads, l)
	}
	return loads, nil
}
