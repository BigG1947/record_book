package db

import "database/sql"

type Direction struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	IdCathedra int    `json:"id_cathedra"`
}

func (d *Direction) GetDirectionById(db *sql.DB, id int) error {
	err := db.QueryRow("SELECT id, name, id_cathedra FROM directon WHERE id = ?", id).Scan(&d.Id, &d.Name, &d.IdCathedra)
	return err
}

func GetAllDirection(db *sql.DB) ([]Direction, error){
	var directions []Direction
	rows, err := db.Query("SELECT id, name, id_cathedra FROM direction")
	if err != nil{
		return []Direction{}, err
	}

	for rows.Next(){
		var d Direction
		err = rows.Scan(&d.Id, &d.Name, &d.IdCathedra)
		if err != nil{
			return []Direction{}, nil
		}
		directions = append(directions, d)
	}
	return directions, nil
}