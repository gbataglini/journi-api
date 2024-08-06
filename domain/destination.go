package domain

type Destination struct {
	ID string `json:"id"`
	City string `json:"city"`
	Country string `json:"country"`
	Visited string `json:"visited"`
}

type DestinationService interface {
	ListDestinations() ([]Destination, error)
	GetDestinationByID(destinationID string) (Destination, error)
	AddDestination(destination Destination) error
	DeleteDestination(destinationID string) ([]Destination, error)
}

type DestinationStore interface {
	DestinationService
}