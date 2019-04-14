package db

import (
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
		return employees, err
	}
	defer rows.Close()

	for rows.Next() {
		var p People
		err := rows.Scan(&p.Id, &p.Fio, &p.Birthday, &p.Gender, &p.Img, &p.Comment, &p.Password, &p.PhoneNumber, &p.Email, &p.Status.Id, &p.HaveAccess)
		if err != nil {
			return employees, err
		}

		var empl Employee
		err = empl.getById(db, p.Id)
		if err != nil {
			return employees, err
		}

		var ac Accession
		err = ac.getById(db, p.Id)
		if err != nil {
			return employees, err
		}

		var sd SensitiveData
		err = sd.getById(db, p.Id)
		if err != nil {
			return employees, err
		}

		err = p.Status.getStatusById(db, p.Status.Id)
		if err != nil {
			return employees, err
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

func GetEmailPasswordMap(db *sql.DB) (map[string]string, error) {
	var result map[string]string

	rows, err := db.Query("SELECT email, password FROM people")
	if err != nil {
		return result, err
	}
	defer rows.Close()

	for rows.Next() {
		var email string
		var password string
		err := rows.Scan(&email, &password)
		if err != nil {
			return result, err
		}
		result[email] = password
	}

	return result, err
}

func (p *People) insert(tx *sql.Tx) (int, error) {
	res, err := tx.Exec("INSERT INTO people(fio, birthday, gender, img, comment, password, phone_number, email, id_status, have_access) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", p.Fio, p.Birthday, p.Gender, p.Img, p.Comment, p.Password, p.PhoneNumber, p.Email, p.Status.Id, p.HaveAccess)
	if err != nil {
		return 0, err
	}
	lastId, _ := res.LastInsertId()
	return int(lastId), err
}
