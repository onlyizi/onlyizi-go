package postgres

import (
	"database/sql"

	gormDriver "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *sql.DB
var gormDB *gorm.DB

func DB() *sql.DB {
	if db == nil {
		panic("postgres not initialized")
	}
	return db
}

func DBGorm() *gorm.DB {

	if db == nil {
		panic("postgres not initialized")
	}

	if gormDB != nil {
		return gormDB
	}

	gdb, err := gorm.Open(
		gormDriver.New(gormDriver.Config{
			Conn: db,
		}),
		&gorm.Config{},
	)

	if err != nil {
		panic(err)
	}

	gormDB = gdb

	return gormDB
}
