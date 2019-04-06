package models

import (
	"../db"
	"database/sql"
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
	Email         string        `json:"email"`
	Status        string        `json:"status"`
	HaveAccess    bool          `json:"have_access"`
	Student       Student       `json:"student"`
	Employee      Employee      `json:"employee"`
	Accession     Accession     `json:"accession"`
	SensitiveData SensitiveData `json:"sensitive_data"`
}

func (p *People) GetPeopleByEmail(db *db.Db, email string) error {
	if db.Status != true || db.Err != nil {
		log.Printf("People.GetPeopleByEmail: failed connection to db\n%s\n", db.Err)
		return db.Err
	}

	err := db.Connection.QueryRow("SELECT id, fio, birthday, gender, comment, password, phone_number, email, status, have_access FROM people WHERE email = (?)", email).Scan(&p.Id, &p.Fio, &p.Birthday, &p.Gender, &p.Comment, &p.Password, &p.PhoneNumber, &p.Email, &p.Status, &p.HaveAccess)
	if err != nil {
		log.Printf("People.GetPeopleByEmail: failed selected people\n%s\n", err)
		return err
	}

	err = p.Student.getById(db.Connection, p.Id)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("People.GetPeopleByEmail: \n%s\n", err)
		return err
	}

	err = p.Employee.getById(db.Connection, p.Id)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("People.GetPeopleByEmail: \n%s\n", err)
		return err
	}

	err = p.Accession.getById(db.Connection, p.Id)
	if err != nil {
		log.Printf("People.GetPeopleByEmail: \n%s\n", err)
		return err
	}

	err = p.SensitiveData.getById(db.Connection, p.Id)
	if err != nil {
		log.Printf("People.GetPeopleByEmail: \n%s\n", err)
		return err
	}
	return nil
}

func GetAllEmployee(db *db.Db) ([]People, error) {
	var students []People

	if db.Status != true || db.Err != nil {
		log.Printf("People.GetAllEmployee: failed connection to db\n%s\n", db.Err)
		return students, db.Err
	}

	rows, err := db.Connection.Query("SELECT id, fio, birthday, gender, comment, password, phone_number, email, status, have_access FROM people WHERE people.id IN (SELECT id_people FROM employee)")
	if err != nil {
		log.Printf("People.GetAllEmployee: failed selected people\n%s\n", err)
		return students, err
	}
	defer rows.Close()

	for rows.Next() {
		var p People
		err := rows.Scan(&p.Id, &p.Fio, &p.Birthday, &p.Gender, &p.Comment, &p.Password, &p.PhoneNumber, &p.Email, &p.Status, &p.HaveAccess)
		if err != nil {
			log.Printf("People.GetAllEmployee: \n%s\n", err)
			return students, err
		}

		var empl Employee
		err = empl.getById(db.Connection, p.Id)
		if err != nil {
			log.Printf("People.GetAllEmployee: \n%s\n", err)
			return students, err
		}

		var ac Accession
		err = ac.getById(db.Connection, p.Id)
		if err != nil {
			log.Printf("People.GetAllEmployee: \n%s\n", err)
			return students, err
		}

		var sd SensitiveData
		err = sd.getById(db.Connection, p.Id)
		if err != nil {
			log.Printf("People.GetAllEmployee: \n%s\n", err)
			return students, err
		}

		p.Employee = empl
		p.Accession = ac
		p.SensitiveData = sd
		students = append(students, p)
	}

	log.Printf("Success GetAllEmployee!!\n")
	return students, nil
}

func GetAllStudent(db *db.Db) ([]People, error) {
	var students []People

	if db.Status != true || db.Err != nil {
		log.Printf("People.GetAllStudent: failed connection to db\n%s\n", db.Err)
		return students, db.Err
	}

	rows, err := db.Connection.Query("SELECT id, fio, birthday, gender, comment, password, phone_number, email, status, have_access FROM people WHERE people.id IN (SELECT id_people FROM student)")
	if err != nil {
		log.Printf("People.GetAllStudent: failed selected people\n%s\n", db.Err)
		return students, err
	}
	defer rows.Close()

	for rows.Next() {
		var p People
		err := rows.Scan(&p.Id, &p.Fio, &p.Birthday, &p.Gender, &p.Comment, &p.Password, &p.PhoneNumber, &p.Email, &p.Status, &p.HaveAccess)
		if err != nil {
			log.Printf("People.GetAllStudent: \n%s\n", err)
			return students, err
		}

		var st Student
		err = st.getById(db.Connection, p.Id)
		if err != nil {
			log.Printf("People.GetAllStudent: \n%s\n", err)
			return students, err
		}

		var ac Accession
		err = ac.getById(db.Connection, p.Id)
		if err != nil {
			log.Printf("People.GetAllStudent: \n%s\n", err)
			return students, err
		}

		var sd SensitiveData
		err = sd.getById(db.Connection, p.Id)
		if err != nil {
			log.Printf("People.GetAllStudent: \n%s\n", err)
			return students, err
		}

		p.Student = st
		p.Accession = ac
		p.SensitiveData = sd
		students = append(students, p)
	}

	log.Printf("Success GetAllStudent!!\n")
	return students, nil
}

func (p *People) InsertEmployee(db *db.Db) {
	if db.Status != true || db.Err != nil {
		log.Printf("People.InsertEmployee: failed connection to db\n%s\n", db.Err)
		return
	}

	tx, err := db.Connection.Begin()
	if err != nil {
		log.Printf("People.InsertEmployee: failed start transaction\n%s\n", err)
		tx.Rollback()
		return
	}

	idPeople, err := p.insert(tx)
	if err != nil {
		log.Printf("People.InsertEmployee: failed insert to table people\n%s\n", err)
		tx.Rollback()
		return
	}

	p.Employee.IdPeople = idPeople
	p.SensitiveData.IdPeople = idPeople
	p.Accession.IdPeople = idPeople

	err = p.Employee.insert(tx)
	if err != nil {
		log.Printf("People.InsertEmployee: failed insert to table employee\n%s\n", err)
		tx.Rollback()
		return
	}

	err = p.Accession.insert(tx)
	if err != nil {
		log.Printf("People.InsertEmployee: failed insert to table accession\n%s\n", err)
		tx.Rollback()
		return
	}

	err = p.SensitiveData.insert(tx)
	if err != nil {
		log.Printf("People.InsertEmployee: failed insert to table sensitive_data\n%s\n", err)
		tx.Rollback()
		return
	}

	err = tx.Commit()
	if err != nil {
		log.Printf("People.InsertEmployee: failed commit transaction\n%s\n", err)
		tx.Rollback()
		return
	}

	log.Printf("People.InsertEmployee: Success insert employee\n")
}

func (p *People) InsertStudent(db *db.Db) {
	if db.Status != true || db.Err != nil {
		log.Printf("People.InsertStudent: failed connection to db\n%s\n", db.Err)
		return
	}

	tx, err := db.Connection.Begin()
	if err != nil {
		log.Printf("People.InsertStudent: failed start transaction\n%s\n", err)
		tx.Rollback()
		return
	}

	idPeople, err := p.insert(tx)
	if err != nil {
		log.Printf("People.InsertStudent: failed insert to table people\n%s\n", err)
		tx.Rollback()
		return
	}

	p.Student.IdPeople = idPeople
	p.SensitiveData.IdPeople = idPeople
	p.Accession.IdPeople = idPeople

	err = p.Student.insert(tx)
	if err != nil {
		log.Printf("People.InsertStudent: failed insert to table student\n%s\n", err)
		tx.Rollback()
		return
	}

	err = p.Accession.insert(tx)
	if err != nil {
		log.Printf("People.InsertStudent: failed insert to table accession\n%s\n", err)
		tx.Rollback()
		return
	}

	err = p.SensitiveData.insert(tx)
	if err != nil {
		log.Printf("People.InsertStudent: failed insert to table sensitive_data\n%s\n", err)
		tx.Rollback()
		return
	}

	err = tx.Commit()
	if err != nil {
		log.Printf("People.InsertStudent: failed commit transaction\n%s\n", err)
		tx.Rollback()
		return
	}

	log.Printf("People.InsertStudent: Success insert student\n")
}

func GetEmailPasswordMap(db *db.Db) map[string]string {
	var result = make(map[string]string)
	if db.Status != true || db.Err != nil {
		log.Printf("People.GetEmailPasswordMap: failed connection to db\n%s\n", db.Err)
		return result
	}

	rows, err := db.Connection.Query("SELECT email, password FROM people")
	if err != nil {
		log.Printf("People.GetEmailPasswordMap: failed select email, password\n%s\n", db.Err)
		return result
	}
	defer rows.Close()

	for rows.Next() {
		var email string
		var password string
		err := rows.Scan(&email, &password)
		if err != nil {
			log.Printf("People.GetEmailPasswordMap: failed select email, password\n%s\n", db.Err)
			return result
		}
		result[email] = password
	}

	return result
}

func (p *People) insert(tx *sql.Tx) (int, error) {
	res, err := tx.Exec("INSERT INTO people(fio, birthday, gender, comment, password, phone_number, email, status, have_access) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)", p.Fio, p.Birthday, p.Gender, p.Comment, p.Password, p.PhoneNumber, p.Email, p.Status, p.HaveAccess)
	if err != nil {
		return 0, err
	}
	lastId, _ := res.LastInsertId()
	return int(lastId), err
}
