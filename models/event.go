package models

import "calendly/db"

type Event struct {
	ID     string `json:"event_id,omitempty"`
	UserId string `json:"user_id"`
	Name   string `json:"name"`
}

func (e Event) Create(event Event) (*Event, error) {
	db.GetMysql().Create(&event)
	print(event.ID)
	return &event, nil
}

func (e Event) ListAll() []Event {
	db.GetMysql().AutoMigrate(&Event{})
	var events []Event
	db.GetMysql().Find(&events)
	return events
}
