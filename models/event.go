package models

import (
	"example.com/restapi/db"
)

type Event struct {
	ID          int64  `json:"id"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Location    string `json:"location"`
	DateTime    string `json:"dateTime"`
	UserID      int    `json:"userID"`
}

var events = []Event{}

func Save(e Event) error {
	//laer to in deb
	query := `INSERT INTO events (name, description, location, dateTime, userID) 
	VALUES (?, ?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		panic(err)
	}
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		panic(err)
	}
	//defer stmt.Close()
	lastID, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	e.ID = lastID

	return err

	//events = append(events, e)
}

func GetAllEvents() ([]Event, error) {
	query := `SELECT * FROM events`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var e Event
		err := rows.Scan(&e.ID, &e.Name, &e.Description, &e.Location, &e.DateTime, &e.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, e)
	}
	return events, nil
}
