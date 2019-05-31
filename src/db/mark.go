package db

import "database/sql"

type Mark struct {
	Id            int    `json:"id"`
	IdStudent     int    `json:"id_student"`
	IdDiscipline  int    `json:"id_discipline"`
	IdEmployee    int    `json:"id_employee"`
	Value         int    `json:"value"`
	NationalValue string `json:"national_value"`
	IsExam        bool   `json:"is_exam"`
	Semester      int    `json:"semester"`
	Date          string `json:"date"`
}

func (m *Mark) Insert(db *sql.DB) (int, error) {
	res, err := db.Exec("INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", m.IdStudent, m.IdDiscipline, m.IdEmployee, m.Value, m.NationalValue, m.IsExam, m.Semester, m.Date)
	if err != nil {
		return 0, err
	}
	lastId, _ := res.LastInsertId()
	return int(lastId), err
}

func GetMarksByStudentId(db *sql.DB, id int) ([]Mark, error) {
	var marks []Mark
	rows, err := db.Query("SELECT id, id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date FROM marks WHERE id_student = ?", id)
	if err != nil {
		return []Mark{}, err
	}

	for rows.Next() {
		var m Mark
		err = rows.Scan(&m.Id, &m.IdStudent, &m.IdDiscipline, &m.IdEmployee, &m.Value, &m.NationalValue, &m.IsExam, &m.Semester, &m.Date)
		if err != nil {
			return []Mark{}, err
		}

		marks = append(marks, m)
	}
	return marks, nil
}

func GetMarksByStudentEmail(db *sql.DB, email string) ([]Mark, error) {
	var marks []Mark
	rows, err := db.Query("SELECT M.id, M.id_student, M.id_discipline, M.id_employee, M.value, M.national_value, M.is_exam, M.semester, M.date FROM marks M, student S WHERE S.email = ? AND M.id_student = S.id", email)
	if err != nil {
		return []Mark{}, err
	}

	for rows.Next() {
		var m Mark
		err = rows.Scan(&m.Id, &m.IdStudent, &m.IdDiscipline, &m.IdEmployee, &m.Value, &m.NationalValue, &m.IsExam, &m.Semester, &m.Date)
		if err != nil {
			return []Mark{}, err
		}

		marks = append(marks, m)
	}
	return marks, nil
}

func (m *Mark) Update(db *sql.DB) error {
	_, err := db.Exec("UPDATE marks SET id_student = (?), id_discipline = (?), is_exam = (?), id_employee = (?), value = (?), national_value = (?), semester = (?), date = (?) WHERE id = (?)", m.IdStudent, m.IdDiscipline, m.IsExam, m.IdEmployee, m.Value, m.NationalValue, m.Semester, m.Date, m.Id)
	return err
}

func (m *Mark) Delete(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM marks WHERE id = ?", m.Id)
	return err
}

func GetMarksByDiscipline(db *sql.DB, idDiscipline int) ([]Mark, error) {
	var marks []Mark
	rows, err := db.Query("SELECT id, id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date FROM marks WHERE id_discipline = ?", idDiscipline)
	if err != nil {
		return []Mark{}, err
	}

	for rows.Next() {
		var m Mark
		err = rows.Scan(&m.Id, &m.IdStudent, &m.IdDiscipline, &m.IdEmployee, &m.Value, &m.NationalValue, &m.IsExam, &m.Semester, &m.Date)
		if err != nil {
			return []Mark{}, err
		}

		marks = append(marks, m)
	}
	return marks, nil
}

func GetStudentMarksByDiscipline(db *sql.DB, idStudent int, idDiscipline int) ([]Mark, error) {
	var marks []Mark
	rows, err := db.Query("SELECT id, id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date FROM marks WHERE id_student = ? AND id_discipline = ?", idStudent, idDiscipline)
	if err != nil {
		return []Mark{}, err
	}

	for rows.Next() {
		var m Mark
		err = rows.Scan(&m.Id, &m.IdStudent, &m.IdDiscipline, &m.IdEmployee, &m.Value, &m.NationalValue, &m.IsExam, &m.Semester, &m.Date)
		if err != nil {
			return []Mark{}, err
		}

		marks = append(marks, m)
	}
	return marks, nil
}

func GetStudentDisciplinesByMark(db *sql.DB, idStudent int) ([]Discipline, error) {
	var disciplines []Discipline
	rows, err := db.Query(getStudentDisciplinesByMarkScript, idStudent)
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
