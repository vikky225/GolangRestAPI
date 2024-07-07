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

	//defer DB.Close()

	createTables(DB)

	return DB
}

func createTables(DB *sql.DB) {
	sql := `
		CREATE TABLE IF NOT EXISTS events (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT,
			description TEXT,
			location TEXT,
			dateTime TEXT,
			userID INTEGER
		);
	`
	_, err := DB.Exec(sql)
	if err != nil {
		panic(err)
	}
}
