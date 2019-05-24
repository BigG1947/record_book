package db

import "database/sql"

type StudentCard struct {
	Id          int    `json:"id"`
	Fio         string `json:"fio"`
	Gender      int    `json:"gender"`
	Photo       string `json:"photo"`
	HaveAccess  bool   `json:"have_access"`
	IdStatus    int    `json:"id_status"`
	StatusName  string `json:"status_name"`
	IdGroup     int    `json:"id_group"`
	GroupName   string `json:"group_name"`
	IdDirection int    `json:"id_direction"`
	IdCathedra  int    `json:"id_cathedra"`
	IdFaculty   int    `json:"id_faculty"`
	IdInstitute int    `json:"id_institute"`
}

func GetStudentCards(db *sql.DB) ([]StudentCard, error) {
	var studentCards []StudentCard
	rows, err := db.Query("SELECT people.id, people.fio, people.gender, people.img, people.have_access, people.id_status, status.name, student.id_group, groups.name, groups.id_direction, direction.id_cathedra, cathedra.id_faculty, faculty.id_institute FROM people, student, groups, direction, cathedra, faculty, status WHERE people.id_status NOT IN (2, 7) AND people.id = student.id_people AND status.id = people.id_status AND groups.id = student.id_group AND direction.id = groups.id_direction AND cathedra.id = direction.id_cathedra AND faculty.id = cathedra.id_faculty ORDER BY have_access DESC, id_status ASC, groups.id DESC")
	if err != nil {
		return []StudentCard{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var card StudentCard
		err := rows.Scan(&card.Id, &card.Fio, &card.Gender, &card.Photo, &card.HaveAccess, &card.IdStatus, &card.StatusName, &card.IdGroup, &card.GroupName, &card.IdDirection, &card.IdCathedra, &card.IdFaculty, &card.IdInstitute)
		if err != nil {
			return []StudentCard{}, err
		}
		studentCards = append(studentCards, card)
	}
	return studentCards, nil
}
