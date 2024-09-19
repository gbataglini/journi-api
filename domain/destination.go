package domain

import (
	"context"
	"errors"
	"time"

	"github.com/lib/pq"
	"googlemaps.github.io/maps"
)

type Country struct {
	Country         string         `json:"country" db:"country"`
	DestinationType pq.StringArray `json:"destinationType" db:"destination_type"`
	Visited         bool           `json:"visited" db:"visited"`
}

type Location struct {
	Lat float64 `json:"lat" db:"lat"`
	Lng float64 `json:"lng" db:"lng"`
}

type Destination struct {
	ID              int        `json:"id" db:"id"`
	GoogleMapsId    string     `json:"googleMapsId" db:"googlemaps_id"`
	UserId          int        `json:"userId" db:"user_id"`
	City            string     `json:"city" db:"city"`
	Country         string     `json:"country" db:"country"`
	Visited         bool       `json:"visited" db:"visited"`
	DestinationType string     `json:"destinationType" db:"destination_type"`
	Location        Location   `json:"location"`
	CreatedAt       time.Time  `json:"createdAt" db:"created_at"`
	ModifiedAt      *time.Time `json:"modifiedAt" db:"modified_at"`
}

type DestinationDAO struct {
	ID              int        `db:"id"`
	GoogleMapsId    string     `db:"googlemaps_id"`
	UserId          int        `db:"user_id"`
	City            string     `db:"city"`
	Country         string     `db:"country"`
	Visited         bool       `db:"visited"`
	DestinationType string     `db:"destination_type"`
	CreatedAt       time.Time  `db:"created_at"`
	ModifiedAt      *time.Time `db:"modified_at"`
	Lat             float64    `db:"lat"`
	Lng             float64    `db:"lng"`
}

func (d DestinationDAO) ToDestination() Destination {
	return Destination{
		ID:              d.ID,
		GoogleMapsId:    d.GoogleMapsId,
		UserId:          d.UserId,
		City:            d.City,
		Country:         d.Country,
		Visited:         d.Visited,
		DestinationType: d.DestinationType,
		CreatedAt:       d.CreatedAt,
		ModifiedAt:      d.ModifiedAt,
		Location: Location{
			Lat: d.Lat,
			Lng: d.Lng,
		},
	}
}

func (d Destination) Validate() error {
	var errs error
	if d.ID != 0 {
		errs = errors.Join(errs, errors.New("id cannot be set"))
	}
	if d.GoogleMapsId == "" {
		errs = errors.Join(errs, errors.New("googleMapsId required"))

	}
	if !d.CreatedAt.IsZero() {
		errs = errors.Join(errs, errors.New("createdAt cannot be set"))
	}
	return errs
}

type DestinationReaderWriter interface {
	ListDestinations(userID int) ([]Destination, error)
	ListCountries(userID int) ([]Country, error)
	GetDestinationByID(destinationID int) (Destination, error)
	AddDestination(destination Destination) (Destination, error)
	DeleteDestination(destinationID int) ([]Destination, error)
}

type GoogleClient interface {
	PlaceAutocomplete(ctx context.Context, r *maps.PlaceAutocompleteRequest) (maps.AutocompleteResponse, error)
	PlaceDetails(ctx context.Context, r *maps.PlaceDetailsRequest) (maps.PlaceDetailsResult, error)
}

type GoogleMapsService interface {
	GooglePlacesSearchSuggestions(input string) (maps.AutocompleteResponse, error)
	GooglePlacesEstablishmentSearch(input string, lat string, lng string) (maps.AutocompleteResponse, error)
	GooglePlacesGetDetails(input string) (maps.PlaceDetailsResult, error)
}

type DestinationService interface {
	DestinationReaderWriter
	GoogleMapsService
}

type DestinationStore interface {
	DestinationReaderWriter
}
