package db

import "database/sql"

type Student struct {
	IdPeople      int    `json:"id_people"`
	DateAdmission string `json:"date_admission"`
	IsFullTime    bool   `json:"is_full_time"`
	IsCut         bool   `json:"is_cut"`
	Group         Group  `json:"group"`
	Semester      int    `json:"semester"`
	Marks         []Mark `json:"marks"`
}

func (st *Student) insert(tx *sql.Tx) error {
	_, err := tx.Exec("INSERT INTO student(id_people, date_admission, is_full_time, is_cut, id_group, semester) VALUES (?, ?, ?, ?, ?, ?)", st.IdPeople, st.DateAdmission, st.IsFullTime, st.IsCut, st.Group.Id, st.Semester)
	return err
}

func (st *Student) getById(db *sql.DB, id int) error {
	err := db.QueryRow("SELECT id_people, date_admission, is_full_time, is_cut, id_group, semester FROM student WHERE id_people = (?)", id).Scan(&st.IdPeople, &st.DateAdmission, &st.IsFullTime, &st.IsCut, &st.Group.Id, &st.Semester)
	if err != nil {
		return err
	}

	err = st.Group.GetGroupById(db, st.Group.Id)
	if err != nil {
		return err
	}
	st.Marks, err = GetMarksByStudentId(db, st.IdPeople)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	return nil
}

func (st *Student) getByIdV2(db *sql.DB, id int) error {
	err := db.QueryRow("SELECT id_people, date_admission, is_full_time, is_cut, id_group, semester FROM student WHERE id_people = (?)", id).Scan(&st.IdPeople, &st.DateAdmission, &st.IsFullTime, &st.IsCut, &st.Group.Id, &st.Semester)
	if err != nil {
		return err
	}

	err = st.Group.GetGroupById(db, st.Group.Id)
	if err != nil {
		return err
	}
	st.Marks, err = GetMarksByStudentId(db, st.IdPeople)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	return nil
}

func ChangeStudentGroup(db *sql.DB, idPeople int, idGroup int) error {
	_, err := db.Exec(changeStudentGroupScript, idGroup, idPeople)
	if err != nil {
		return err
	}
	return nil
}

func (s *Student) Update(db *sql.DB) error {
	_, err := db.Exec(updateStudentDataScript, s.DateAdmission, s.IsFullTime, s.IsCut, s.Group.Id, s.Semester, s.IdPeople)
	if err != nil {
		return err
	}
	return err
}

func UpdateStudentData(db *sql.DB, s *Student) error {
	_, err := db.Exec(updateStudentDataScript, s.DateAdmission, s.IsFullTime, s.IsCut, s.Group.Id, s.Semester, s.IdPeople)
	if err != nil {
		return err
	}
	return err
}
