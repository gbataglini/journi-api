package domain

import "time"

type Destination struct {
	ID         int        `json:"id" db:"id"`
	UserId     int        `json:"userId" db:"user_id"`
	City       string     `json:"city" db:"city"`
	Country    string     `json:"country" db:"country"`
	Visited    string     `json:"visited" db:"visited"`
	CreatedAt  time.Time  `json:"createdAt" db:"created_at"`
	ModifiedAt *time.Time `json:"modifiedAt" db:"modified_at"`
}

type DestinationService interface {
	ListDestinations() ([]Destination, error)
	GetDestinationByID(destinationID int) (Destination, error)
	AddDestination(destination Destination) (Destination, error)
	DeleteDestination(destinationID int) ([]Destination, error)
}

type DestinationStore interface {
	DestinationService
}
