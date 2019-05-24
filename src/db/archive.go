package db

import "database/sql"

func GetArchiveEmployeeCard(db *sql.DB) ([]EmployeeCard, error) {
	var employeeCards []EmployeeCard
	rows, err := db.Query("SELECT people.id, people.fio, people.gender, people.img, people.have_access, people.id_status, status.name, employee.id_rank, ranks.name, employee.id_cathedra, cathedra.id_faculty, faculty.id_institute FROM people, employee, ranks, cathedra, faculty, status WHERE people.id_status = 9 AND people.id = employee.id_people AND status.id = people.id_status AND ranks.id = employee.id_rank AND cathedra.id = employee.id_cathedra AND faculty.id = cathedra.id_faculty ORDER BY have_access DESC, id_status ASC, id_rank ASC")
	if err != nil {
		return []EmployeeCard{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var card EmployeeCard
		err := rows.Scan(&card.Id, &card.Fio, &card.Gender, &card.Photo, &card.HaveAccess, &card.IdStatus, &card.StatusName, &card.IdRank, &card.RankName, &card.IdCathedra, &card.IdFaculty, &card.IdInstitute)
		if err != nil {
			return []EmployeeCard{}, err
		}
		employeeCards = append(employeeCards, card)
	}
	return employeeCards, nil
}

func GetArchiveEmployee(db *sql.DB, id int) (People, error) {
	var employee People

	err := db.QueryRow(getArchiveEmployeeByIdScript, id).Scan(&employee.Id, &employee.Fio, &employee.Birthday, &employee.Gender, &employee.Img, &employee.Comment, &employee.Password, &employee.PhoneNumber, &employee.Email, &employee.Sensitive, &employee.HaveAccess,
		&employee.Employee.IdPeople, &employee.Employee.DateInvite, &employee.Employee.Rank.Id, &employee.Employee.Rank.Name, &employee.Employee.Cathedra.Id, &employee.Employee.Cathedra.Name, &employee.Employee.Cathedra.IdFaculty,
		&employee.Accession.EditAccess, &employee.Accession.SetAbsence, &employee.Accession.GetAbsence, &employee.Accession.SetMark, &employee.Accession.SetEvent, &employee.Accession.GetSensitive, &employee.Accession.SetSensitive, &employee.Accession.ManageLoad, &employee.Accession.ManageAcadem,
		&employee.Status.Id, &employee.Status.Name)
	if err != nil && err != sql.ErrNoRows {
		return People{}, err
	}
	return employee, nil
}

func GetArchiveStudentCard(db *sql.DB) ([]StudentCard, error) {
	var studentCards []StudentCard
	rows, err := db.Query(getArchiveStudentsCardScripts)
	if err != nil {
		return []StudentCard{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var card StudentCard
		err := rows.Scan(&card.Id, &card.Fio, &card.Gender, &card.Photo, &card.HaveAccess, &card.IdStatus, &card.StatusName)
		if err != nil {
			return []StudentCard{}, err
		}
		studentCards = append(studentCards, card)
	}
	return studentCards, nil
}

func GetArchiveStudent(db *sql.DB, id int) (People, error) {
	var student People
	var employeeId sql.NullInt64
	var groupId sql.NullInt64
	var groupName sql.NullString
	var groupIdDirection sql.NullInt64
	err := db.QueryRow(getArchiveStudentByIdScript, id).Scan(&student.Id, &student.Fio, &student.Birthday, &student.Gender, &student.Img, &student.Comment, &student.Password, &student.PhoneNumber, &student.Email, &student.Sensitive, &student.HaveAccess,
		&student.Student.IdPeople, &student.Student.DateAdmission, &student.Student.IsFullTime, &student.Student.IsCut, &groupId, &employeeId, &groupName, &groupIdDirection, &student.Student.Semester,
		&student.Accession.EditAccess, &student.Accession.SetAbsence, &student.Accession.GetAbsence, &student.Accession.SetMark, &student.Accession.SetEvent, &student.Accession.GetSensitive, &student.Accession.SetSensitive, &student.Accession.ManageLoad, &student.Accession.ManageAcadem,
		&student.Status.Id, &student.Status.Name)
	if groupId.Valid {
		student.Student.Group.Id = int(groupId.Int64)
	}
	if employeeId.Valid {
		student.Student.Group.IdEmployee = int(employeeId.Int64)
	}
	if groupName.Valid {
		student.Student.Group.Name = groupName.String
	}
	if groupIdDirection.Valid {
		student.Student.Group.IdDirection = int(groupIdDirection.Int64)
	}

	if err != nil && err != sql.ErrNoRows {
		return People{}, err
	}
	return student, nil
}
