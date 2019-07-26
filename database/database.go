package database

import (
	"database/sql"
	"util"

	_ "github.com/go-sql-driver/mysql"
)

const (
	DB_Driver = "root:@tcp(127.0.0.1:3306)/test?charset=utf8"
)

func OpenDB() (success bool, db *sql.DB) {
	var isOpen bool
	db, err := sql.Open("mysql", DB_Driver)
	if err != nil {
		isOpen = false
	} else {
		isOpen = true
	}
	util.CheckErr(err)
	return isOpen, db
}
