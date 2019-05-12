package db

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
)

type People struct {
	Id            int           `json:"id"`
	Fio           string        `json:"fio"`
	Birthday      string        `json:"birthday"`
	Gender        int           `json:"gender"`
	Comment       string        `json:"comment"`
	Password      string        `json:"password"`
	PhoneNumber   string        `json:"phone_number"`
	Img           string        `json:"img"`
	Email         string        `json:"email"`
	Status        Status        `json:"status"`
	HaveAccess    bool          `json:"have_access"`
	Student       Student       `json:"student"`
	Employee      Employee      `json:"employee"`
	Accession     Accession     `json:"accession"`
	SensitiveData SensitiveData `json:"sensitive_data"`
}

func (p *People) GetPeopleByEmail(db *sql.DB, email string) error {

	err := db.QueryRow("SELECT id, fio, birthday, gender, img, comment, password, phone_number, email, id_status, have_access FROM people WHERE email = (?)", email).Scan(&p.Id, &p.Fio, &p.Birthday, &p.Gender, &p.Img, &p.Comment, &p.Password, &p.PhoneNumber, &p.Email, &p.Status.Id, &p.HaveAccess)
	if err != nil {
		log.Printf("People.GetPeopleByEmail: failed selected people\n%s\n", err)
		return err
	}

	err = p.Student.getById(db, p.Id)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("People.GetPeopleByEmail: \n%s\n", err)
		return err
	}

	err = p.Employee.getById(db, p.Id)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("People.GetPeopleByEmail: \n%s\n", err)
		return err
	}

	err = p.Accession.getById(db, p.Id)
	if err != nil {
		log.Printf("People.GetPeopleByEmail: \n%s\n", err)
		return err
	}

	err = p.SensitiveData.getById(db, p.Id)
	if err != nil {
		log.Printf("People.GetPeopleByEmail: \n%s\n", err)
		return err
	}

	err = p.Status.getStatusById(db, p.Status.Id)
	if err != nil {
		return err
	}
	return nil
}

func (p *People) GetPeopleById(db *sql.DB, id int) error {

	err := db.QueryRow("SELECT id, fio, birthday, gender, img, comment, password, phone_number, email, id_status, have_access FROM people WHERE id = (?)", id).Scan(&p.Id, &p.Fio, &p.Birthday, &p.Gender, &p.Img, &p.Comment, &p.Password, &p.PhoneNumber, &p.Email, &p.Status.Id, &p.HaveAccess)
	if err != nil {
		return err
	}

	err = p.Student.getById(db, p.Id)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	err = p.Employee.getById(db, p.Id)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	err = p.Accession.getById(db, p.Id)
	if err != nil {
		return err
	}

	err = p.SensitiveData.getById(db, p.Id)
	if err != nil {
		return err
	}

	err = p.Status.getStatusById(db, p.Status.Id)
	return nil
}

func GetAllEmployee(db *sql.DB) ([]People, error) {
	var employees []People

	rows, err := db.Query("SELECT id, fio, birthday, gender, img, comment, password, phone_number, email, id_status, have_access FROM people WHERE people.id IN (SELECT id_people FROM employee)")
	if err != nil {
		return []People{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var p People
		err := rows.Scan(&p.Id, &p.Fio, &p.Birthday, &p.Gender, &p.Img, &p.Comment, &p.Password, &p.PhoneNumber, &p.Email, &p.Status.Id, &p.HaveAccess)
		if err != nil {
			return []People{}, err
		}

		var empl Employee
		err = empl.getById(db, p.Id)
		if err != nil {
			return []People{}, err
		}

		var ac Accession
		err = ac.getById(db, p.Id)
		if err != nil {
			return []People{}, err
		}

		var sd SensitiveData
		err = sd.getById(db, p.Id)
		if err != nil {
			return []People{}, err
		}

		err = p.Status.getStatusById(db, p.Status.Id)
		if err != nil {
			return []People{}, err
		}

		p.Employee = empl
		p.Accession = ac
		p.SensitiveData = sd
		employees = append(employees, p)
	}

	return employees, nil
}

func GetAllStudent(db *sql.DB) ([]People, error) {
	var students []People

	rows, err := db.Query("SELECT id, fio, birthday, gender, img, comment, password, phone_number, email, id_status, have_access FROM people WHERE people.id IN (SELECT id_people FROM student)")
	if err != nil {
		return []People{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var p People
		err := rows.Scan(&p.Id, &p.Fio, &p.Birthday, &p.Gender, &p.Img, &p.Comment, &p.Password, &p.PhoneNumber, &p.Email, &p.Status.Id, &p.HaveAccess)
		if err != nil {
			return []People{}, err
		}

		var st Student
		err = st.getById(db, p.Id)
		if err != nil {
			return []People{}, err
		}

		var ac Accession
		err = ac.getById(db, p.Id)
		if err != nil {
			return []People{}, err
		}

		var sd SensitiveData
		err = sd.getById(db, p.Id)
		if err != nil {
			return []People{}, err
		}

		err = p.Status.getStatusById(db, p.Status.Id)
		if err != nil {
			return []People{}, err
		}

		p.Student = st
		p.Accession = ac
		p.SensitiveData = sd
		students = append(students, p)
	}

	return students, nil
}

func GetStudentFromGroup(db *sql.DB, id int) ([]People, error) {
	var students []People

	rows, err := db.Query("SELECT id, fio, birthday, gender, img, comment, password, phone_number, email, id_status, have_access FROM people WHERE id IN (SELECT id_people FROM student WHERE id_group = ?)", id)
	if err != nil {
		return students, err
	}
	defer rows.Close()

	for rows.Next() {
		var p People
		err := rows.Scan(&p.Id, &p.Fio, &p.Birthday, &p.Gender, &p.Img, &p.Comment, &p.Password, &p.PhoneNumber, &p.Email, &p.Status.Id, &p.HaveAccess)
		if err != nil {
			return students, err
		}

		var st Student
		err = st.getById(db, p.Id)
		if err != nil {
			return students, err
		}

		var ac Accession
		err = ac.getById(db, p.Id)
		if err != nil {
			return students, err
		}

		var sd SensitiveData
		err = sd.getById(db, p.Id)
		if err != nil {
			return students, err
		}

		err = p.Status.getStatusById(db, p.Status.Id)
		if err != nil {
			return students, err
		}

		p.Student = st
		p.Accession = ac
		p.SensitiveData = sd
		students = append(students, p)
	}

	return students, nil
}

func (p *People) InsertEmployee(db *sql.DB) error {

	tx, err := db.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}

	idPeople, err := p.insert(tx)
	if err != nil {
		tx.Rollback()
		return err
	}

	p.Employee.IdPeople = idPeople
	p.SensitiveData.IdPeople = idPeople
	p.Accession.IdPeople = idPeople

	err = p.Employee.insert(tx)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = p.Accession.insert(tx)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = p.SensitiveData.insert(tx)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (p *People) InsertStudent(db *sql.DB) error {

	tx, err := db.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}

	idPeople, err := p.insert(tx)
	if err != nil {
		tx.Rollback()
		return err
	}

	p.Student.IdPeople = idPeople
	p.SensitiveData.IdPeople = idPeople
	p.Accession.IdPeople = idPeople

	err = p.Student.insert(tx)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = p.Accession.insert(tx)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = p.SensitiveData.insert(tx)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (p *People) insert(tx *sql.Tx) (int, error) {
	res, err := tx.Exec("INSERT INTO people(fio, birthday, gender, img, comment, password, phone_number, email, id_status, have_access) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", p.Fio, p.Birthday, p.Gender, p.Img, p.Comment, p.Password, p.PhoneNumber, p.Email, p.Status.Id, p.HaveAccess)
	if err != nil {
		return 0, err
	}
	lastId, _ := res.LastInsertId()
	return int(lastId), err
}

func GetAllEmployeesV2(db *sql.DB) ([]People, error) {
	var students []People

	rows, err := db.Query(getAllEmployeeScript)
	if err != nil {
		return []People{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var p People
		err := rows.Scan(&p.Id, &p.Fio, &p.Birthday, &p.Gender, &p.Img, &p.Comment, &p.Password, &p.PhoneNumber, &p.Email, &p.HaveAccess,
			&p.Employee.IdPeople, &p.Employee.DateInvite, &p.Employee.Rank.Id, &p.Employee.Rank.Name, &p.Employee.Cathedra.Id, &p.Employee.Cathedra.Name, &p.Employee.Cathedra.IdFaculty,
			&p.Accession.EditAccess, &p.Accession.SetAbsence, &p.Accession.GetAbsence, &p.Accession.SetMark, &p.Accession.SetEvent, &p.Accession.GetSensitive, &p.Accession.SetSensitive, &p.Accession.GetYlist, &p.Accession.ManageAcadem,
			&p.SensitiveData.PassportCode, &p.SensitiveData.Rntrs, &p.SensitiveData.RegAddress, &p.SensitiveData.MilitaryId,
			&p.Status.Id, &p.Status.Name)
		if err != nil {
			return []People{}, err
		}
		p.Employee.Group, err = GetGroupByIdEmployee(db, p.Id)
		if err != nil {
			return []People{}, err
		}
		students = append(students, p)
	}

	return students, nil
}

func GetStudentFromGroupV2(db *sql.DB, id int) ([]People, error) {
	var students []People

	rows, err := db.Query(getStudentFromGroupScript, id)
	if err != nil {
		return students, err
	}
	defer rows.Close()

	for rows.Next() {
		var p People
		var employeeId sql.NullInt64
		err := rows.Scan(&p.Id, &p.Fio, &p.Birthday, &p.Gender, &p.Img, &p.Comment, &p.Password, &p.PhoneNumber, &p.Email, &p.HaveAccess,
			&p.Student.IdPeople, &p.Student.DateAdmission, &p.Student.IsFullTime, &p.Student.IsCut, &p.Student.Group.Id, &employeeId, &p.Student.Group.Name, &p.Student.Group.IdDirection, &p.Student.Semester,
			&p.Accession.EditAccess, &p.Accession.SetAbsence, &p.Accession.GetAbsence, &p.Accession.SetMark, &p.Accession.SetEvent, &p.Accession.GetSensitive, &p.Accession.SetSensitive, &p.Accession.GetYlist, &p.Accession.ManageAcadem,
			&p.SensitiveData.PassportCode, &p.SensitiveData.Rntrs, &p.SensitiveData.RegAddress, &p.SensitiveData.MilitaryId,
			&p.Status.Id, &p.Status.Name)
		if employeeId.Valid {
			p.Student.Group.IdEmployee = int(employeeId.Int64)
		}
		if err != nil {
			return []People{}, err
		}

		p.Student.Marks, err = GetMarksByStudentId(db, p.Id)
		if err != nil {
			return []People{}, err
		}

		students = append(students, p)
	}

	return students, nil
}

func GetAllStudentsV2(db *sql.DB) ([]People, error) {
	var students []People

	rows, err := db.Query(getAllStudentsScript)
	if err != nil {
		return []People{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var p People
		var employeeId sql.NullInt64
		err := rows.Scan(&p.Id, &p.Fio, &p.Birthday, &p.Gender, &p.Img, &p.Comment, &p.Password, &p.PhoneNumber, &p.Email, &p.HaveAccess,
			&p.Student.IdPeople, &p.Student.DateAdmission, &p.Student.IsFullTime, &p.Student.IsCut, &p.Student.Group.Id, &employeeId, &p.Student.Group.Name, &p.Student.Group.IdDirection, &p.Student.Semester,
			&p.Accession.EditAccess, &p.Accession.SetAbsence, &p.Accession.GetAbsence, &p.Accession.SetMark, &p.Accession.SetEvent, &p.Accession.GetSensitive, &p.Accession.SetSensitive, &p.Accession.GetYlist, &p.Accession.ManageAcadem,
			&p.SensitiveData.PassportCode, &p.SensitiveData.Rntrs, &p.SensitiveData.RegAddress, &p.SensitiveData.MilitaryId,
			&p.Status.Id, &p.Status.Name)
		if employeeId.Valid {
			p.Student.Group.IdEmployee = int(employeeId.Int64)
		}
		if err != nil {
			return []People{}, err
		}

		p.Student.Marks, err = GetMarksByStudentId(db, p.Id)
		if err != nil {
			return []People{}, err
		}

		students = append(students, p)
	}

	return students, nil
}

func (p *People) GetPeopleByEmailV2(db *sql.DB, email string) error {
	var studentIdPeople sql.NullInt64
	var dateAdmission sql.NullString
	var isFullTime sql.NullBool
	var isCut sql.NullBool
	var studentGroupId sql.NullInt64
	var studentGroupEmployeeId sql.NullInt64
	var studentGroupName sql.NullString
	var studentGroupIdDirection sql.NullInt64
	var studentSemester sql.NullInt64
	var employeeIdPeople sql.NullInt64
	var employeeDateInvite sql.NullString
	var employeeRankId sql.NullInt64
	var employeeRankName sql.NullString
	var employeeCathedraId sql.NullInt64
	var employeeCathedraName sql.NullString
	var employeeCathedraIdFaculty sql.NullInt64
	row := db.QueryRow(getPeopleByEmailScript, email)
	err := row.Scan(&p.Id, &p.Fio, &p.Birthday, &p.Gender, &p.Img, &p.Comment, &p.Password, &p.PhoneNumber, &p.Email, &p.HaveAccess,
		&studentIdPeople, &dateAdmission, &isFullTime, &isCut, &studentGroupId, &studentGroupEmployeeId, &studentGroupName, &studentGroupIdDirection, &studentSemester,
		&employeeIdPeople, &employeeDateInvite, &employeeRankId, &employeeRankName, &employeeCathedraId, &employeeCathedraName, &employeeCathedraIdFaculty,
		&p.Accession.EditAccess, &p.Accession.SetAbsence, &p.Accession.GetAbsence, &p.Accession.SetMark, &p.Accession.SetEvent, &p.Accession.GetSensitive, &p.Accession.SetSensitive, &p.Accession.GetYlist, &p.Accession.ManageAcadem,
		&p.SensitiveData.PassportCode, &p.SensitiveData.Rntrs, &p.SensitiveData.RegAddress, &p.SensitiveData.MilitaryId,
		&p.Status.Id, &p.Status.Name)
	if err != nil {
		return err
	}
	if studentIdPeople.Valid {
		p.Student.IdPeople = int(studentIdPeople.Int64)
		if dateAdmission.Valid {
			p.Student.DateAdmission = dateAdmission.String
		}
		if isFullTime.Valid {
			p.Student.IsFullTime = isFullTime.Bool
		}
		if isCut.Valid {
			p.Student.IsCut = isCut.Bool
		}
		if studentGroupId.Valid {
			p.Student.Group.Id = int(studentGroupId.Int64)
		}
		if studentGroupName.Valid {
			p.Student.Group.Name = studentGroupName.String
		}
		if studentGroupEmployeeId.Valid {
			p.Student.Group.IdEmployee = int(studentGroupEmployeeId.Int64)
		}
		if studentGroupIdDirection.Valid {
			p.Student.Group.IdDirection = int(studentGroupIdDirection.Int64)
		}
		if studentSemester.Valid {
			p.Student.Semester = int(studentSemester.Int64)
		}
		p.Student.Marks, err = GetMarksByStudentId(db, p.Id)
		if err != nil {
			return err
		}
	}
	if employeeIdPeople.Valid {
		p.Employee.IdPeople = int(employeeIdPeople.Int64)
		if employeeDateInvite.Valid {
			p.Employee.DateInvite = employeeDateInvite.String
		}
		if employeeRankId.Valid {
			p.Employee.Rank.Id = int(employeeRankId.Int64)
		}
		if employeeRankName.Valid {
			p.Employee.Rank.Name = employeeRankName.String
		}
		if employeeCathedraId.Valid {
			p.Employee.Cathedra.Id = int(employeeCathedraId.Int64)
		}
		if employeeCathedraName.Valid {
			p.Employee.Cathedra.Name = employeeRankName.String
		}
		if employeeCathedraIdFaculty.Valid {
			p.Employee.Cathedra.IdFaculty = int(employeeCathedraId.Int64)
		}
		p.Employee.Group, err = GetGroupByIdEmployee(db, p.Id)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *People) GetPeopleByIdV2(db *sql.DB, id int) error {
	var studentIdPeople sql.NullInt64
	var dateAdmission sql.NullString
	var isFullTime sql.NullBool
	var isCut sql.NullBool
	var studentGroupId sql.NullInt64
	var studentGroupEmployeeId sql.NullInt64
	var studentGroupName sql.NullString
	var studentGroupIdDirection sql.NullInt64
	var studentSemester sql.NullInt64
	var employeeIdPeople sql.NullInt64
	var employeeDateInvite sql.NullString
	var employeeRankId sql.NullInt64
	var employeeRankName sql.NullString
	var employeeCathedraId sql.NullInt64
	var employeeCathedraName sql.NullString
	var employeeCathedraIdFaculty sql.NullInt64
	row := db.QueryRow(getPeopleByIdScript, id)
	err := row.Scan(&p.Id, &p.Fio, &p.Birthday, &p.Gender, &p.Img, &p.Comment, &p.Password, &p.PhoneNumber, &p.Email, &p.HaveAccess,
		&studentIdPeople, &dateAdmission, &isFullTime, &isCut, &studentGroupId, &studentGroupEmployeeId, &studentGroupName, &studentGroupIdDirection, &studentSemester,
		&employeeIdPeople, &employeeDateInvite, &employeeRankId, &employeeRankName, &employeeCathedraId, &employeeCathedraName, &employeeCathedraIdFaculty,
		&p.Accession.EditAccess, &p.Accession.SetAbsence, &p.Accession.GetAbsence, &p.Accession.SetMark, &p.Accession.SetEvent, &p.Accession.GetSensitive, &p.Accession.SetSensitive, &p.Accession.GetYlist, &p.Accession.ManageAcadem,
		&p.SensitiveData.PassportCode, &p.SensitiveData.Rntrs, &p.SensitiveData.RegAddress, &p.SensitiveData.MilitaryId,
		&p.Status.Id, &p.Status.Name)
	if err != nil {
		return err
	}
	if studentIdPeople.Valid {
		p.Student.IdPeople = int(studentIdPeople.Int64)
		if dateAdmission.Valid {
			p.Student.DateAdmission = dateAdmission.String
		}
		if isFullTime.Valid {
			p.Student.IsFullTime = isFullTime.Bool
		}
		if isCut.Valid {
			p.Student.IsCut = isCut.Bool
		}
		if studentGroupId.Valid {
			p.Student.Group.Id = int(studentGroupId.Int64)
		}
		if studentGroupName.Valid {
			p.Student.Group.Name = studentGroupName.String
		}
		if studentGroupEmployeeId.Valid {
			p.Student.Group.IdEmployee = int(studentGroupEmployeeId.Int64)
		}
		if studentGroupIdDirection.Valid {
			p.Student.Group.IdDirection = int(studentGroupIdDirection.Int64)
		}
		if studentSemester.Valid {
			p.Student.Semester = int(studentSemester.Int64)
		}
		p.Student.Marks, err = GetMarksByStudentId(db, p.Id)
		if err != nil {
			return err
		}
	}
	if employeeIdPeople.Valid {
		p.Employee.IdPeople = int(employeeIdPeople.Int64)
		if employeeDateInvite.Valid {
			p.Employee.DateInvite = employeeDateInvite.String
		}
		if employeeRankId.Valid {
			p.Employee.Rank.Id = int(employeeRankId.Int64)
		}
		if employeeRankName.Valid {
			p.Employee.Rank.Name = employeeRankName.String
		}
		if employeeCathedraId.Valid {
			p.Employee.Cathedra.Id = int(employeeCathedraId.Int64)
		}
		if employeeCathedraName.Valid {
			p.Employee.Cathedra.Name = employeeRankName.String
		}
		if employeeCathedraIdFaculty.Valid {
			p.Employee.Cathedra.IdFaculty = int(employeeCathedraId.Int64)
		}
		p.Employee.Group, err = GetGroupByIdEmployee(db, p.Id)
		if err != nil {
			return err
		}
	}
	return nil
}

func BlockPeopleById(db *sql.DB, id int) error {
	_, err := db.Exec(blockPeopleScript, id)
	if err != nil {
		return err
	}
	return nil
}

func UnblockPeopleById(db *sql.DB, id int) error {
	_, err := db.Exec(unblockPeopleScript, id)
	if err != nil {
		return err
	}
	return nil
}

func UpdatePeopleStatus(db *sql.DB, idPeople int, idStatus int) error {
	_, err := db.Exec(updatePeopleStatusScript, idStatus, idPeople)
	if err != nil {
		return err
	}
	return nil
}

func ChangePeoplePassword(db *sql.DB, idPeople int, password string) error {
	_, err := db.Exec(changePeoplePasswordScript, password, idPeople)
	if err != nil {
		return err
	}
	return nil
}

func UpdatePeopleData(db *sql.DB, idPeople int, updateData map[string]string) error {
	tableName, ok := updateData["table"]
	if !ok {
		return errors.New("Неверный формат мапы, остутсвует поле 'table'")
	}
	delete(updateData, "table")
	lenData := len(updateData)
	sqlScript := fmt.Sprintf("UPDATE people, %s SET", tableName)
	var i int
	for key, value := range updateData {
		if i < lenData-1 {
			sqlScript += fmt.Sprintf(" %s = '%s',", key, value)
		} else if i == lenData-1 {
			sqlScript += fmt.Sprintf(" %s = '%s'", key, value)
		}
		i++
	}
	sqlScript += fmt.Sprintf(" WHERE people.id = ? AND %s.id_people = people.id;", tableName)
	fmt.Printf(sqlScript + "\n")
	_, err := db.Exec(sqlScript, idPeople)
	if err != nil {
		return err
	}
	return nil
}
