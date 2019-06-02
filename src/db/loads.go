package db

import (
	"database/sql"
	"time"
)

type Load struct {
	Id            int          `json:"id"`
	Discipline    Discipline   `json:"discipline"`
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

func GetLoadForEmployeeByIdSemester(db *sql.DB, idEmployee int, idSemester int) ([]Load, error) {
	var loads []Load
	rows, err := db.Query(getAllLoadsForEmployeeBuIdSemesterScript, idEmployee, idSemester)
	if err != nil {
		return []Load{}, err
	}

	for rows.Next() {
		var l Load
		var IdAssistant sql.NullInt64
		var NameAssistant sql.NullString
		err = rows.Scan(&l.Id, &l.Discipline.Id, &l.Discipline.Name, &l.IdEmployee, &l.NameEmployee, &l.IdGroup, &l.NameGroup, &IdAssistant, &NameAssistant, &l.Semester.Id, &l.Semester.Start, &l.Semester.End, &l.Semester.Name, &l.NumSemester)
		if err != nil {
			return []Load{}, err
		}
		if IdAssistant.Valid {
			l.IdAssistant = int(IdAssistant.Int64)
		}
		if NameAssistant.Valid {
			l.NameAssistant = NameAssistant.String
		}
		loads = append(loads, l)
	}
	return loads, nil
}

func GetLoadSemesterById(db *sql.DB, id int) (LoadSemester, error) {
	var ls LoadSemester
	err := db.QueryRow("SELECT id, start, end, name FROM loads_semester WHERE id = ?", id).Scan(&ls.Id, &ls.Start, &ls.End, &ls.Name)
	if err != nil {
		return LoadSemester{}, err
	}
	return ls, nil
}

func GetAllLoadsSemester(db *sql.DB) ([]LoadSemester, error) {
	var loadSemesters []LoadSemester
	rows, err := db.Query("SELECT id, start, end, name FROM loads_semester ORDER BY end DESC")
	if err != nil {
		return []LoadSemester{}, err
	}
	for rows.Next() {
		var semester LoadSemester
		err = rows.Scan(&semester.Id, &semester.Start, &semester.End, &semester.Name)
		if err != nil {
			return []LoadSemester{}, err
		}
		loadSemesters = append(loadSemesters, semester)
	}
	return loadSemesters, nil
}

func (ls *LoadSemester) Insert(db *sql.DB) (int, error) {
	res, err := db.Exec("INSERT INTO loads_semester(start, end, name) VALUES (?,?,?)", ls.Start, ls.End, ls.Name)
	if err != nil {
		return 0, err
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(lastId), nil
}

func (ls *LoadSemester) Update(db *sql.DB) error {
	_, err := db.Exec("UPDATE loads_semester SET start = ?, end = ?, name = ? WHERE id = ?", ls.Start, ls.End, ls.Name, ls.Id)
	if err != nil {
		return err
	}
	return nil
}

func DeleteLoadSemester(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM loads_semester WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}

func (l *Load) Insert(db *sql.DB) (int, error) {
	res, err := db.Exec("INSERT INTO loads(id_discipline, id_employee, id_group, id_assistant, semester, id_semester) VALUES (?, ?, ?, ?, ?, ?)", l.Discipline.Id, l.IdEmployee, l.IdGroup, checkForeignKey(l.IdAssistant), l.NumSemester, l.Semester.Id)
	if err != nil {
		return 0, err
	}

	lastId, _ := res.LastInsertId()
	return int(lastId), nil
}

func GetAllLoadsForEmployee(db *sql.DB, id int) ([]Load, error) {
	var loads []Load
	rows, err := db.Query(getAllLoadsByIdEmployeeScript, id)
	if err != nil {
		return []Load{}, err
	}

	for rows.Next() {
		var l Load
		var IdAssistant sql.NullInt64
		var NameAssistant sql.NullString
		err = rows.Scan(&l.Id, &l.Discipline.Id, &l.Discipline.Name, &l.IdEmployee, &l.NameEmployee, &l.IdGroup, &l.NameGroup, &IdAssistant, &NameAssistant, &l.Semester.Id, &l.Semester.Start, &l.Semester.End, &l.Semester.Name, &l.NumSemester)
		if err != nil {
			return []Load{}, err
		}
		if IdAssistant.Valid {
			l.IdAssistant = int(IdAssistant.Int64)
		}
		if NameAssistant.Valid {
			l.NameAssistant = NameAssistant.String
		}
		loads = append(loads, l)
	}
	return loads, nil
}

func GetAllLoadsForAssistent(db *sql.DB, id int) ([]Load, error) {
	var loads []Load
	rows, err := db.Query(getAllLoadsByIdAssistantScript, id)
	if err != nil {
		return []Load{}, err
	}

	for rows.Next() {
		var l Load
		var IdAssistant sql.NullInt64
		var NameAssistant sql.NullString
		err = rows.Scan(&l.Id, &l.Discipline.Id, &l.Discipline.Name, &l.IdEmployee, &l.NameEmployee, &l.IdGroup, &l.NameGroup, &IdAssistant, &NameAssistant, &l.Semester.Id, &l.Semester.Start, &l.Semester.End, &l.Semester.Name, &l.NumSemester)
		if err != nil {
			return []Load{}, err
		}
		if IdAssistant.Valid {
			l.IdAssistant = int(IdAssistant.Int64)
		}
		if NameAssistant.Valid {
			l.NameAssistant = NameAssistant.String
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
		var IdAssistant sql.NullInt64
		var NameAssistant sql.NullString
		err = rows.Scan(&l.Id, &l.Discipline.Id, &l.Discipline.Name, &l.IdEmployee, &l.NameEmployee, &l.IdGroup, &l.NameGroup, &IdAssistant, &NameAssistant, &l.Semester.Id, &l.Semester.Start, &l.Semester.End, &l.Semester.Name, &l.NumSemester)
		if err != nil {
			return []Load{}, err
		}
		if IdAssistant.Valid {
			l.IdAssistant = int(IdAssistant.Int64)
		}
		if NameAssistant.Valid {
			l.NameAssistant = NameAssistant.String
		}
		loads = append(loads, l)
	}
	return loads, nil
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
		var IdAssistant sql.NullInt64
		var NameAssistant sql.NullString
		err = rows.Scan(&l.Id, &l.Discipline.Id, &l.Discipline.Name, &l.IdEmployee, &l.NameEmployee, &l.IdGroup, &l.NameGroup, &IdAssistant, &NameAssistant, &l.Semester.Id, &l.Semester.Start, &l.Semester.End, &l.Semester.Name, &l.NumSemester)
		if err != nil {
			return []Load{}, err
		}
		if IdAssistant.Valid {
			l.IdAssistant = int(IdAssistant.Int64)
		}
		if NameAssistant.Valid {
			l.NameAssistant = NameAssistant.String
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

func GetLoadsById(db *sql.DB, id int) (Load, error) {
	var l Load
	var IdAssistant sql.NullInt64
	var NameAssistant sql.NullString
	err := db.QueryRow(getAllLoadsByIdScript, id).Scan(&l.Id, &l.Discipline.Id, &l.Discipline.Name, &l.IdEmployee, &l.NameEmployee, &l.IdGroup, &l.NameGroup, &IdAssistant, &NameAssistant, &l.Semester.Id, &l.Semester.Start, &l.Semester.End, &l.Semester.Name, &l.NumSemester)
	if err != nil {
		return Load{}, err
	}
	if IdAssistant.Valid {
		l.IdAssistant = int(IdAssistant.Int64)
	}
	if NameAssistant.Valid {
		l.NameAssistant = NameAssistant.String
	}
	return l, nil
}

func UpdateLoadsById(db *sql.DB, l *Load) error {
	_, err := db.Exec(updateLoadsById, l.Discipline.Id, l.IdEmployee, l.IdGroup, l.IdAssistant, l.NumSemester, l.Semester.Id, l.Id)
	if err != nil {
		return err
	}
	return nil
}
