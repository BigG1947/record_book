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
	rows, err := db.Query(getCurrentStudentCardScript)
	if err != nil {
		return []StudentCard{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var card StudentCard

		var IdGroup sql.NullInt64
		var GroupName sql.NullString
		var IdDirection sql.NullInt64
		var IdCathedra sql.NullInt64
		var IdFaculty sql.NullInt64
		var IdInstitute sql.NullInt64

		err := rows.Scan(&card.Id, &card.Fio, &card.Gender, &card.Photo, &card.HaveAccess, &card.IdStatus, &card.StatusName, &IdGroup, &GroupName, &IdDirection, &IdCathedra, &IdFaculty, &IdInstitute)
		if err != nil {
			return []StudentCard{}, err
		}
		if IdGroup.Valid {
			card.IdGroup = int(IdGroup.Int64)
			if GroupName.Valid {
				card.GroupName = GroupName.String
			}
			if IdDirection.Valid {
				card.IdDirection = int(IdDirection.Int64)
			}
			if IdCathedra.Valid {
				card.IdCathedra = int(IdCathedra.Int64)
			}
			if IdFaculty.Valid {
				card.IdFaculty = int(IdFaculty.Int64)
			}
			if IdInstitute.Valid {
				card.IdInstitute = int(IdInstitute.Int64)
			}
		}
		studentCards = append(studentCards, card)
	}
	return studentCards, nil
}
