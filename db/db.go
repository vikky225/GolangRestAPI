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

	createUsersTables := `
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			email TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL
		);
	`
	_, err := DB.Exec(createUsersTables)
	if err != nil {
		panic(err)
	}

	sql := `
		CREATE TABLE IF NOT EXISTS events (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT,
			description TEXT,
			location TEXT,
			dateTime TEXT,
			userID INTEGER,
			FOREIGN KEY(userID) REFERENCES users(id)
		);
	`
	_, err = DB.Exec(sql)
	if err != nil {
		panic(err)
	}

	createRegisrationTables := `
		CREATE TABLE IF NOT EXISTS registrations (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			userID INTEGER,
			eventID INTEGER,
			FOREIGN KEY(userID) REFERENCES users(id),
			FOREIGN KEY(eventID) REFERENCES events(id)
		);
	`
	_, err = DB.Exec(createRegisrationTables)
	if err != nil {
		panic(err)
	}
}
