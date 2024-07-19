package models

type Event struct {
	ID               int    `json:"id" db:"id" gorm:"primaryKey""`
	EventName        string `json:"event_name" db:"event_name"`
	EventDate        string `json:"event_date" db:"event_date"`
	EventDescription string `json:"event_description" json:"event_description"`
}
