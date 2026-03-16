package postgres

import "database/sql"

var db *sql.DB

func DB() *sql.DB {
	if db == nil {
		panic("postgres not initialized")
	}
	return db
}
