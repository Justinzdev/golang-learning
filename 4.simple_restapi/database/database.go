package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitializeDB(username, password, hostname, port, database string) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, hostname, port, database)
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("Connect Fail.")
	} else {
		fmt.Println("Connect Success")
	}
}

func GetDB() *sql.DB {
	return db
}
