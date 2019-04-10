package db

import (
	"database/sql"
)

func checkForeignKey(fk int) sql.NullInt64 {
	if fk > 0 {
		return sql.NullInt64{Int64: int64(fk), Valid: true}
	}
	return sql.NullInt64{
		Int64: 0,
		Valid: false,
	}
}

func getCashData(db *sql.DB) (emailPeople map[string]People, emailPassword map[string]string, emailSensetiveData map[string]SensitiveData, err error) {
	var allUsers []People

	students, err := GetAllStudent(db)
	if err != nil{
		return
	}
	employees, err := GetAllEmployee(db)
	if err != nil{
		return
	}
	allUsers = append(allUsers, students...)
	allUsers = append(allUsers, employees...)

	for _, people := range allUsers{
		emailPeople[people.Email] = people
		emailSensetiveData[people.Email] = people.SensitiveData
		}
	emailPassword, err = GetEmailPasswordMap(db)
	if err != nil{
		return
	}

	err = nil
	return
}
