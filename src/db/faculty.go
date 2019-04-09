package db

import "database/sql"

type Faculty struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	IdInstitute int    `json:"id_institute"`
}

func (f *Faculty) GetFacultyById(db *sql.DB, id int) error {
	err := db.QueryRow("SELECT id, name, id_institute FROM faculty WHERE id = ?", id).Scan(&f.Id, &f.Name, &f.IdInstitute)
	return err
}

func GetAllFaculty(db *sql.DB) ([]Faculty, error){
	var faculty []Faculty
	rows, err := db.Query("SELECT id, name, id_institute FROM faculty")
	if err != nil{
		return []Faculty{}, err
	}

	for rows.Next(){
		var f Faculty
		err = rows.Scan(&f.Id, &f.Name, &f.IdInstitute)
		if err != nil{
			return []Faculty{}, nil
		}
		faculty = append(faculty, f)
	}
	return faculty, nil
}
