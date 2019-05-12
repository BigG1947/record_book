package db

import (
	"database/sql"
)

func checkForeignKey(fk int) sql.NullInt64 {
	if fk > 0 {
		return sql.NullInt64{Int64: int64(fk), Valid: true}
	}
	return sql.NullInt64{
		Int64: 0,
		Valid: false,
	}
}

func GetEmailPasswordMap(db *sql.DB) (map[string]string, error) {
	var result = make(map[string]string)

	rows, err := db.Query("SELECT email, password FROM people")
	if err != nil {
		return result, err
	}
	defer rows.Close()

	for rows.Next() {
		var email string
		var password string
		err := rows.Scan(&email, &password)
		if err != nil {
			return result, err
		}
		result[email] = password
	}

	return result, err
}
