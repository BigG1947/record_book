package db

import "database/sql"

type Group struct {
	Id   int
	Name string
	IdDirection int
}

func (g *Group) GetGroupById(db *sql.DB, id int) error {
	err := db.QueryRow("SELECT id, name, id_direction FROM groups WHERE id = ?;", id).Scan(&g.Id, &g.Name, &g.IdDirection)
	return err
}

func (g *Group) Insert(db *sql.DB) (int, error){
	res, err := db.Exec("INSERT INTO groups(name, id_direction) VALUES (?, ?);", g.Name, &g.IdDirection)
	if err != nil{
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil{
		return 0, err
	}
	return int(id), err
}

func GetAllGroup(db *sql.DB) ([]Group, error) {
	var groups []Group
	rows, err := db.Query("SELECT id, name, id_direction FROM groups")
	if err != nil {
		return groups, err
	}
	defer rows.Close()

	for rows.Next() {
		var g Group
		err := rows.Scan(&g.Id, &g.Name, &g.IdDirection)
		if err != nil {
			return groups, err
		}
		groups = append(groups, g)
	}

	return groups, nil
}