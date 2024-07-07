package models

import "example.com/restapi/db"

type Event struct {
	ID          int    `json:"id"`
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
	defer stmt.Close()
	lastID, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	e.ID = int(lastID)

	return err

	//events = append(events, e)
}

func GetAllEvents() []Event {
	return events
}
