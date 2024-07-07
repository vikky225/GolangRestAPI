package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB = InitDB()

func InitDB() *sql.DB {
	DB, err := sql.Open("sqlite3", "api.db")
	if err != nil {
		panic(err)
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	defer DB.Close()

	return DB
}
