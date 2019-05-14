package db

import "database/sql"

type Employee struct {
	IdPeople    int          `json:"id_people"`
	DateInvite  string       `json:"date_invite"`
	Rank        Rank         `json:"rank"`
	Group       []Group      `json:"group"`
	Cathedra    Cathedra     `json:"cathedra"`
	Disciplines []Discipline `json:"disciplines"`
}

func (empl *Employee) insert(tx *sql.Tx) error {
	_, err := tx.Exec("INSERT INTO employee(id_people, date_invite, id_rank, id_cathedra) VALUES (?, ?, ?, ?)", empl.IdPeople, empl.DateInvite, empl.Rank.Id, empl.Cathedra.Id)
	return err
}

func GetLoadEmployees(db *sql.DB) ([]People, error) {
	var employees []People

	rows, err := db.Query("SELECT id, fio, birthday, gender, img, comment, password, phone_number, email, sensitive_data, id_status, have_access FROM people WHERE people.id IN (SELECT id_employee FROM loads)")
	if err != nil {
		return employees, err
	}
	defer rows.Close()

	for rows.Next() {
		var p People
		err := rows.Scan(&p.Id, &p.Fio, &p.Birthday, &p.Gender, &p.Img, &p.Comment, &p.Password, &p.PhoneNumber, &p.Email, &p.Sensitive, &p.Status.Id, &p.HaveAccess)
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

		//var sd SensitiveData
		//err = sd.getById(db, p.Id)
		//if err != nil {
		//	return employees, err
		//}

		err = p.Status.getStatusById(db, p.Status.Id)
		if err != nil {
			return employees, err
		}

		p.Employee = empl
		p.Accession = ac
		//p.SensitiveData = sd
		employees = append(employees, p)
	}

	return employees, nil
}

func (empl *Employee) getById(db *sql.DB, id int) error {
	err := db.QueryRow("SELECT id_people, date_invite, id_rank, id_cathedra FROM employee WHERE id_people = (?)", id).Scan(&empl.IdPeople, &empl.DateInvite, &empl.Rank.Id, &empl.Cathedra.Id)
	if err != nil {
		return err
	}
	err = empl.Cathedra.GetCathedraById(db, empl.Cathedra.Id)
	if err != nil {
		return err
	}
	err = empl.Rank.GetRankById(db, empl.Rank.Id)
	if err != nil {
		return err
	}
	empl.Group, err = GetGroupByIdEmployee(db, empl.IdPeople)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	empl.Disciplines, err = GetAllDisciplineForEmployee(db, empl.IdPeople)
	if err != nil {
		return err
	}
	return nil
}

func (empl *Employee) Update(db *sql.DB) error {
	_, err := db.Exec(updateEmployeeDataScript, empl.DateInvite, empl.Rank.Id, empl.Cathedra.Id, empl.IdPeople)
	if err != nil {
		return err
	}
	return nil
}

func UpdateEmployeeData(db *sql.DB, empl *Employee) error {
	_, err := db.Exec(updateEmployeeDataScript, empl.DateInvite, empl.Rank.Id, empl.Cathedra.Id, empl.IdPeople)
	if err != nil {
		return err
	}
	return nil
}
