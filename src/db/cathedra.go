package db

import "database/sql"

type Cathedra struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	IdFaculty int    `json:"id_faculty"`
}

func (c *Cathedra) GetCathedraById(db *sql.DB, id int) error {
	err := db.QueryRow("SELECT id, name, id_faculty FROM cathedra WHERE id = ?", id).Scan(&c.Id, &c.Name, &c.IdFaculty)
	return err
}

func GetAllCathedra(db *sql.DB) ([]Cathedra, error){
	var cathedra []Cathedra
	rows, err := db.Query("SELECT id, name, id_faculty FROM cathedra")
	if err != nil{
		return []Cathedra{}, err
	}

	for rows.Next(){
		var c Cathedra
		err = rows.Scan(&c.Id, &c.Name, &c.IdFaculty)
		if err != nil{
			return []Cathedra{}, nil
		}
		cathedra = append(cathedra, c)
	}
	return cathedra, nil
}
