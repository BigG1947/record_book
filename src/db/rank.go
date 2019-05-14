package db

import "database/sql"

type Rank struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (r *Rank) GetRankById(db *sql.DB, id int) error {
	err := db.QueryRow("SELECT id, name FROM ranks WHERE id = ?", id).Scan(&r.Id, &r.Name)
	return err
}

func GetAllRanks(db *sql.DB) ([]Rank, error) {
	var ranks []Rank
	rows, err := db.Query("SELECT id, name FROM ranks")
	if err != nil {
		return []Rank{}, err
	}

	for rows.Next() {
		var r Rank
		err = rows.Scan(&r.Id, &r.Name)
		if err != nil {
			return []Rank{}, err
		}
		ranks = append(ranks, r)
	}

	return ranks, nil
}

func UpdateRank(db *sql.DB, r *Rank) error {
	_, err := db.Exec(updateRankScript, r.Name, r.Id)
	if err != nil {
		return err
	}
	return nil
}

func DeleteRank(db *sql.DB, id int) error {
	_, err := db.Exec(deleteRankScript, id)
	if err != nil {
		return err
	}
	return nil
}

func InsertRank(db *sql.DB, r *Rank) (int, error) {
	res, err := db.Exec(insertRankScript, r.Name)
	if err != nil {
		return 0, err
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(lastId), nil
}
