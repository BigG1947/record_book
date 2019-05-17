package db

import "database/sql"

type Load struct {
	Id            int          `json:"id"`
	Discipline    Discipline   `json:"id_discipline"`
	IdEmployee    int          `json:"id_employee"`
	NameEmployee  string       `json:"name_employee"`
	IdGroup       int          `json:"id_group"`
	NameGroup     string       `json:"name_group"`
	IdAssistant   int          `json:"id_assistant"`
	NameAssistant string       `json:"name_assistant"`
	NumSemester   int          `json:"num_semester"`
	Semester      LoadSemester `json:"semester"`
}

type LoadSemester struct {
	Id    int    `json:"id"`
	Start int64  `json:"start"`
	End   int64  `json:"end"`
	Name  string `json:"name"`
}

func (l *Load) Insert(db *sql.DB) (int, error) {
	res, err := db.Exec("INSERT INTO loads(id_discipline, id_employee, id_group, id_assistant, semester, id_semester) VALUES (?, ?, ?, ?, ?, ?)", l.Discipline.Id, l.IdEmployee, l.IdGroup, l.IdAssistant, l.NumSemester, l.Semester.Id)
	if err != nil {
		return 0, err
	}

	lastId, _ := res.LastInsertId()
	return int(lastId), nil
}

func GetAllLoadsForEmployee(db *sql.DB, id int) ([]Load, error) {
	var loads []Load
	rows, err := db.Query(getAllLoadsByIdAssistantScript, id)
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

func GetAllLoadsForAssistent(db *sql.DB, id int) ([]Load, error) {
	var loads []Load
	rows, err := db.Query(getAllLoadsByIdEmployeeScript, id)
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

func GetAllLoadsByIdGroup(db *sql.DB, id int) ([]Load, error) {
	var loads []Load
	rows, err := db.Query(getAllLoadsByIdGroupScript, id)
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

func DeleteLoadsById(db *sql.DB, id int) error {
	_, err := db.Exec(deleteLoadsByIdScript, id)
	if err != nil {
		return err
	}
	return nil
}

func DeleteLoadsSemesterById(db *sql.DB, id int) error {
	_, err := db.Exec(deleteLoadsSemesterByIdScript, id)
	if err != nil {
		return err
	}
	return nil
}

func UpdateLoadsById(db *sql.DB, l *Load) error {
	_, err := db.Exec(updateLoadsById, l.Discipline.Id, l.IdEmployee, l.IdGroup, l.IdAssistant, l.NumSemester, l.Semester.Id, l.Id)
	if err != nil {
		return err
	}
	return nil
}
