package domain

import (
	"time"

	"googlemaps.github.io/maps"
)

type Destination struct {
	ID         int        `json:"id" db:"id"`
	UserId     int        `json:"userId" db:"user_id"`
	City       string     `json:"city" db:"city"`
	Country    string     `json:"country" db:"country"`
	Visited    string     `json:"visited" db:"visited"`
	CreatedAt  time.Time  `json:"createdAt" db:"created_at"`
	ModifiedAt *time.Time `json:"modifiedAt" db:"modified_at"`
}

type DestinationReaderWriter interface {
	ListDestinations() ([]Destination, error)
	GetDestinationByID(destinationID int) (Destination, error)
	AddDestination(destination Destination) (Destination, error)
	DeleteDestination(destinationID int) ([]Destination, error)
}

type GoogleMapsService interface {
	GooglePlacesSearchSuggestions(input string) (maps.AutocompleteResponse, error)
}

type DestinationService interface {
	DestinationReaderWriter
	GoogleMapsService
}

type DestinationStore interface {
	DestinationReaderWriter
}
