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
