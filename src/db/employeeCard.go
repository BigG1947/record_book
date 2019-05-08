package db

import "database/sql"

type EmployeeCard struct {
	Id          int    `json:"id"`
	Fio         string `json:"fio"`
	Gender      int    `json:"gender"`
	Photo       string `json:"photo"`
	HaveAccess  bool   `json:"have_access"`
	IdRank      int    `json:"id_rank"`
	RankName    string `json:"group_name"`
	IdCathedra  int    `json:"id_cathedra"`
	IdFaculty   int    `json:"id_faculty"`
	IdInstitute int    `json:"id_institute"`
}

func GetEmployeeCards(db *sql.DB) ([]EmployeeCard, error) {
	var employeeCards []EmployeeCard
	rows, err := db.Query("SELECT people.id, people.fio, people.gender, people.img, people.have_access, employee.id_rank, employee.id_cathedra, ranks.name, cathedra.id_faculty, faculty.id_institute FROM people, employee, ranks, cathedra, faculty WHERE people.id = employee.id_people AND ranks.id = employee.id_rank AND cathedra.id = employee.id_cathedra AND faculty.id = cathedra.id_faculty ORDER BY have_access DESC")
	if err != nil {
		return []EmployeeCard{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var card EmployeeCard
		err := rows.Scan(&card.Id, &card.Fio, &card.Gender, &card.Photo, &card.HaveAccess, &card.IdRank, &card.IdCathedra, &card.RankName, &card.IdFaculty, &card.IdInstitute)
		if err != nil {
			return []EmployeeCard{}, err
		}
		employeeCards = append(employeeCards, card)
	}
	return employeeCards, nil
}
