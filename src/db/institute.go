package db

import "database/sql"

type Institute struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (i *Institute) GetInstituteById(db *sql.DB, id int) error {
	err := db.QueryRow("SELECT id, name FROM institute WHERE id = ?", id).Scan(&i.Id, &i.Name)
	return err
}

func GetAllInstitute(db *sql.DB) ([]Institute, error){
	var institute []Institute
	rows, err := db.Query("SELECT id, name FROM institute")
	if err != nil{
		return []Institute{}, err
	}

	for rows.Next(){
		var i Institute
		err = rows.Scan(&i.Id, &i.Name)
		if err != nil{
			return []Institute{}, nil
		}
		institute = append(institute, i)
	}
	return institute, nil
}
