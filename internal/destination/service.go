package destination

import (
	"context"
	"fmt"

	owm "github.com/briandowns/openweathermap"
	"github.com/gbataglini/journi-api/domain"
	"github.com/gbataglini/journi-api/internal/config"
	"googlemaps.github.io/maps"
)

type svc struct {
	store        domain.DestinationStore
	googleClient domain.GoogleClient
}

func NewService(store domain.DestinationStore, googleClient domain.GoogleClient) domain.DestinationService {
	return &svc{
		store:        store,
		googleClient: googleClient,
	}
}

func (s *svc) ListDestinations(userID int) ([]domain.Destination, error) {
	return s.store.ListDestinations(userID)
}

func (s *svc) ListCountries(userID int) ([]domain.Country, error) {
	return s.store.ListCountries(userID)
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

func (s *svc) GooglePlacesGetDetails(input string) (maps.PlaceDetailsResult, error) {
	resp, err := s.googleClient.PlaceDetails(context.Background(), &maps.PlaceDetailsRequest{
		PlaceID: input,
	})
	if err != nil {
		return maps.PlaceDetailsResult{}, err
	}
	return resp, nil
}

func (s *svc) GooglePlacesEstablishmentSearch(input string, lat string, lng string) (maps.AutocompleteResponse, error) {
	latlng := fmt.Sprintf("%s,%s", lat, lng)
	formatLatLng, err := maps.ParseLatLng(latlng)

	if err != nil {
		return maps.AutocompleteResponse{}, err
	}

	resp, err := s.googleClient.PlaceAutocomplete(context.Background(), &maps.PlaceAutocompleteRequest{
		Input:        input,
		Types:        "establishment",
		Location:     &formatLatLng,
		Radius:       80500,
		StrictBounds: true,
	})
	if err != nil {
		return maps.AutocompleteResponse{}, err
	}
	return resp, nil
}

func (s *svc) OpenWeatherGetCurrentWeather(lat float64, lng float64) (*owm.CurrentWeatherData, error) {
	cfg := config.Get()
	resp, err := owm.NewCurrent("C", "EN", cfg.OpenWeatherMapApiKey)
	if err != nil {
		return nil, fmt.Errorf("could not get location weather: %w", err)
	}

	resp.CurrentByCoordinates(&owm.Coordinates{Longitude: lng, Latitude: lat})

	return resp, nil
}
