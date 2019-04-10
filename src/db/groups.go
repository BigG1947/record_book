package db

import "database/sql"

type Group struct {
	Id          int    `json:"id"`
	IdEmployee  int    `json:"id_employee"`
	Name        string `json:"name"`
	IdDirection int    `json:"id_direction"`
}

func (g *Group) GetGroupById(db *sql.DB, id int) error {
	var idEmployee sql.NullInt64
	err := db.QueryRow("SELECT id, id_employee, name, id_direction FROM groups WHERE id = ?;", id).Scan(&g.Id, &idEmployee, &g.Name, &g.IdDirection)
	g.IdEmployee = int(idEmployee.Int64)
	return err
}

func GetGroupByIdEmployee(db *sql.DB, id int) ([]Group, error) {
	var groups []Group
	rows, err := db.Query("SELECT id, id_employee, name, id_direction FROM groups WHERE id_employee = ?;", id)
	if err != nil {
		return groups, err
	}
	for rows.Next() {
		var g Group
		var idEmployee sql.NullInt64
		err = rows.Scan(&g.Id, &idEmployee, &g.Name, &g.IdDirection)
		if err != nil && err != sql.ErrNoRows {
			return groups, err
		}
		groups = append(groups, g)
		g.IdEmployee = int(idEmployee.Int64)
	}
	return groups, nil
}

func (g *Group) Insert(db *sql.DB) (int, error) {
	res, err := db.Exec("INSERT INTO groups(id_mployee, name, id_direction) VALUES (? ,?, ?);", checkForeignKey(g.IdEmployee), g.Name, &g.IdDirection)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), err
}

func GetAllGroup(db *sql.DB) ([]Group, error) {
	var groups []Group
	rows, err := db.Query("SELECT id, id_employee, name, id_direction FROM groups")
	if err != nil {
		return groups, err
	}
	defer rows.Close()

	for rows.Next() {
		var g Group
		var idEmployee sql.NullInt64
		err := rows.Scan(&g.Id, &idEmployee, &g.Name, &g.IdDirection)
		g.IdEmployee = int(idEmployee.Int64)
		if err != nil {
			return groups, err
		}
		groups = append(groups, g)
	}

	return groups, nil
}
