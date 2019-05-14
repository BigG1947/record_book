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

func GetAllInstitute(db *sql.DB) ([]Institute, error) {
	var institute []Institute
	rows, err := db.Query("SELECT id, name FROM institute")
	if err != nil {
		return []Institute{}, err
	}

	for rows.Next() {
		var i Institute
		err = rows.Scan(&i.Id, &i.Name)
		if err != nil {
			return []Institute{}, nil
		}
		institute = append(institute, i)
	}
	return institute, nil
}

func UpdateInstitute(db *sql.DB, i *Institute) error {
	_, err := db.Exec(updateInstituteScript, i.Name, i.Id)
	if err != nil {
		return err
	}
	return nil
}

func InsertInstitute(db *sql.DB, i *Institute) (int, error) {
	res, err := db.Exec(insertInstituteScript, i.Name)
	if err != nil {
		return 0, err
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(lastId), nil
}

func DeleteInstitute(db *sql.DB, id int) error {
	_, err := db.Exec(deleteInstituteScript, id)
	if err != nil {
		return err
	}
	return nil
}
