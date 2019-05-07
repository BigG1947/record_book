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
	res, err := db.Exec("INSERT INTO groups(id_employee, name, id_direction) VALUES (? ,?, ?);", checkForeignKey(g.IdEmployee), g.Name, &g.IdDirection)
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

func GetAllGroupByEmployeeAndDiscipline(db *sql.DB, idEmployee int, idDiscipline int) ([]Group, error) {
	var groups []Group
	rows, err := db.Query("SELECT groups.id, groups.id_employee, groups.name, groups.id_direction FROM groups, loads WHERE loads.id_group in (SELECT DISTINCT student.id_group FROM student, people WHERE people.id_status = 1 AND student.id_people = people.id) AND loads.id_employee = ? AND loads.id_discipline = ? AND groups.id = loads.id_group;", idEmployee, idDiscipline)
	if err != nil {
		return []Group{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var g Group
		var idEmployee sql.NullInt64
		err := rows.Scan(&g.Id, &idEmployee, &g.Name, &g.IdDirection)
		g.IdEmployee = int(idEmployee.Int64)
		if err != nil {
			return []Group{}, err
		}
		groups = append(groups, g)
	}

	return groups, nil
}

func GetAllGroupByDirection(db *sql.DB, idDirection int) ([]Group, error) {
	var groups []Group
	rows, err := db.Query("SELECT id, id_employee, name FROM groups WHERE id_direction = ?", idDirection)
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
		g.IdDirection = idDirection
		groups = append(groups, g)
	}

	return groups, nil
}

func GetAllGroupByCathedra(db *sql.DB, idCathedra int) ([]Group, error) {
	var groups []Group
	rows, err := db.Query("SELECT groups.id, groups.id_employee, groups.name, groups.id_direction FROM groups, direction WHERE direction.id_cathedra = ? AND groups.id_direction = direction.id", idCathedra)
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

func GetAllGroupByFaculty(db *sql.DB, idFaculty int) ([]Group, error) {
	var groups []Group
	rows, err := db.Query("SELECT groups.id, groups.id_employee, groups.name, groups.id_direction FROM groups, direction, cathedra WHERE cathedra.id_faculty = ? AND direction.id_cathedra = cathedra.id AND groups.id_direction = direction.id", idFaculty)
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

func GetAllGroupByInstitute(db *sql.DB, idInstitute int) ([]Group, error) {
	var groups []Group
	rows, err := db.Query("SELECT groups.id, groups.id_employee, groups.name, groups.id_direction FROM groups, institute, direction, cathedra, faculty WHERE faculty.id_institute = ? AND cathedra.id_faculty = faculty.id AND direction.id_cathedra = cathedra.id AND groups.id_direction = direction.id", idInstitute)
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
