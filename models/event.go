package models

type Event struct {
	ID          int    `json:"id"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Location    string `json:"location"`
	DateTime    string `json:"dateTime"`
	UserID      int    `json:"userID"`
}

var events = []Event{}

func Save(e Event) {
	//laer to in deb
	events = append(events, e)
}

func GetAllEvents() []Event {
	return events
}
