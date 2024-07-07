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
	UserID      int64  `json:"userID"`
}

var events = []Event{}

func Save(e *Event) error {
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

func GetEventByID(id int64) (*Event, error) {

	query := "SELECT * FROM events WHERE id = ?"
	row := db.DB.QueryRow(query, id)

	var e Event
	err := row.Scan(&e.ID, &e.Name, &e.Description, &e.Location, &e.DateTime, &e.UserID)
	if err != nil {
		return nil, err
	}
	return &e, nil
}

func Update(e Event) error {
	query := "UPDATE events SET name = ?, description = ?, location = ?, dateTime = ?, userID = ? WHERE id = ?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID, e.ID)
	if err != nil {
		return err
	}
	return nil
}

func (event Event) Delete() error {
	query := "DELETE FROM events WHERE id = ?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(event.ID)
	return err
}
