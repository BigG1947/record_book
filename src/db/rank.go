package db

import "database/sql"

type Rank struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (r *Rank) GetRankById(db *sql.DB, id int) error{
	err := db.QueryRow("SELECT id, name FROM ranks WHERE id = ?", id).Scan(&r.Id, &r.Name)
	return err
}

func GetAllRanks(db *sql.DB) ([]Rank, error){
	var ranks []Rank
	rows, err := db.Query("SELECT id, name FROM ranks")
	if err != nil{
		return []Rank{}, err
	}

	for rows.Next(){
		var r Rank
		err = rows.Scan(&r.Id, &r.Name)
		if err != nil{
			return []Rank{}, err
		}
		ranks = append(ranks, r)
	}

	return ranks, nil
}