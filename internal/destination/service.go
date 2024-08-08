package destination

import (
	"context"

	"github.com/gbataglini/journi-api/domain"
	"googlemaps.github.io/maps"
)

type svc struct {
	store        domain.DestinationStore
	googleClient *maps.Client
}

func NewService(store domain.DestinationStore, googleClient *maps.Client) domain.DestinationService {
	return &svc{
		store:        store,
		googleClient: googleClient,
	}
}

func (s *svc) ListDestinations() ([]domain.Destination, error) {
	return s.store.ListDestinations()
}

func (s *svc) GetDestinationByID(destinationID int) (domain.Destination, error) {
	return s.store.GetDestinationByID(destinationID)
}

func (s *svc) AddDestination(destination domain.Destination) (domain.Destination, error) {
	return s.store.AddDestination(destination)
}

func (s *svc) DeleteDestination(destinationID int) ([]domain.Destination, error) {
	return s.store.DeleteDestination(destinationID)
}

func (s *svc) GooglePlacesSearchSuggestions(input string) (maps.AutocompleteResponse, error) {
	resp, err := s.googleClient.PlaceAutocomplete(context.Background(), &maps.PlaceAutocompleteRequest{
		Input: input,
		Types: "(cities)",
	})
	if err != nil {
		return maps.AutocompleteResponse{}, err
	}
	return resp, nil
}
