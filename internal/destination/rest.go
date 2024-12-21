package destination

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gbataglini/journi-api/domain"
)

type rest struct {
	svc domain.DestinationService
}

func NewRest(service domain.DestinationService) domain.Router {
	return &rest{
		svc: service,
	}
}

func (re *rest) onError(w http.ResponseWriter, err error) {
	w.WriteHeader(500)
	fmt.Println(err)
}

func (re *rest) Routes(s *http.ServeMux) {
	s.HandleFunc("GET /api/v1/destinations/autocomplete", re.placesAutocompleteAPI)
	s.HandleFunc("GET /api/v1/destinations/autocomplete/", re.placesAutocompleteAPI)

	s.HandleFunc("GET /api/v1/destinations/placesDetails", re.placesGetDetails)
	s.HandleFunc("GET /api/v1/destinations/placesDetails/", re.placesGetDetails)

	s.HandleFunc("GET /api/v1/destinations/establishmentSearch", re.placesGetEstablishments)
	s.HandleFunc("GET /api/v1/destinations/establishmentSearch/", re.placesGetEstablishments)

	s.HandleFunc("GET /api/v1/{userID}/destinations", re.listDestinations)
	s.HandleFunc("GET /api/v1/{userID}/destinations/", re.listDestinations)

	s.HandleFunc("GET /api/v1/{userID}/countries", re.listCountries)
	s.HandleFunc("GET /api/v1/{userID}/countries/", re.listCountries)

	s.HandleFunc("GET /api/v1/{userID}/destinations/{id}", re.getByID)
	s.HandleFunc("GET /api/v1/{userID}/destinations/{id}/", re.getByID)

	s.HandleFunc("POST /api/v1/destinations", re.create)
	s.HandleFunc("POST /api/v1/destinations/", re.create)

	s.HandleFunc("DELETE /api/v1/{userID}/destinations/{id}", re.delete)
	s.HandleFunc("DELETE /api/v1/{userID}/destinations/{id}/", re.delete)

	s.HandleFunc("GET api/v1/destinations/weather/", re.owmGetPlaceWeather)
	s.HandleFunc("GET api/v1/destinations/weather", re.owmGetPlaceWeather)

}

func (re *rest) listDestinations(w http.ResponseWriter, r *http.Request) {
	userID, err :=
		strconv.Atoi(r.PathValue("userID"))
	if err != nil {
		re.onError(w, err)
		return
	}
	dests, err := re.svc.ListDestinations(userID)
	if err != nil {
		re.onError(w, err)
		return
	}
	json.NewEncoder(w).Encode(dests)
}

func (re *rest) listCountries(w http.ResponseWriter, r *http.Request) {
	userID, err :=
		strconv.Atoi(r.PathValue("userID"))

	if err != nil {
		re.onError(w, err)
		return
	}
	dests, err := re.svc.ListCountries(userID)
	if err != nil {
		re.onError(w, err)
		return
	}

	json.NewEncoder(w).Encode(dests)
}

func (re *rest) create(w http.ResponseWriter, r *http.Request) {
	var newDestination domain.Destination
	if err := json.NewDecoder(r.Body).Decode(&newDestination); err != nil {
		re.onError(w, err)
		return
	}

	newDestination, err := re.svc.AddDestination(newDestination)
	if err != nil {
		re.onError(w, err)
		return
	}
	json.NewEncoder(w).Encode(newDestination)
}

func (re *rest) delete(w http.ResponseWriter, r *http.Request) {
	destinationID, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		re.onError(w, err)
		return
	}
	userID, err := strconv.Atoi(r.PathValue("userID"))
	if err != nil {
		fmt.Println(err)
	}
	re.svc.DeleteDestination(destinationID)
	dests, err := re.svc.ListDestinations(userID)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
		return
	}
	json.NewEncoder(w).Encode(dests)
}

func (re *rest) getByID(w http.ResponseWriter, r *http.Request) {
	destinationID, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		fmt.Println(err)
	}
	selectedDest, err := re.svc.GetDestinationByID(destinationID)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
		return
	}

	json.NewEncoder(w).Encode(selectedDest)
}

func (re *rest) placesAutocompleteAPI(w http.ResponseWriter, r *http.Request) {
	searchParam := r.URL.Query().Get("searchParam")
	response, err := re.svc.GooglePlacesSearchSuggestions(searchParam)

	if err != nil {
		re.onError(w, err)
		return
	}
	json.NewEncoder(w).Encode(response)
}

func (re *rest) placesGetDetails(w http.ResponseWriter, r *http.Request) {
	locationId := r.URL.Query().Get("locationId")
	response, err := re.svc.GooglePlacesGetDetails(locationId)

	if err != nil {
		re.onError(w, err)
		return
	}
	json.NewEncoder(w).Encode(response)
}

func (re *rest) placesGetEstablishments(w http.ResponseWriter, r *http.Request) {
	searchParam := r.URL.Query().Get("searchParam")
	lat := r.URL.Query().Get("lat")
	lng := r.URL.Query().Get("lng")
	response, err := re.svc.GooglePlacesEstablishmentSearch(searchParam, lat, lng)

	if err != nil {
		re.onError(w, err)
		return
	}
	json.NewEncoder(w).Encode(response)
}

func (re *rest) owmGetPlaceWeather(w http.ResponseWriter, r *http.Request) {
	lat, err := strconv.ParseFloat(r.URL.Query().Get("lat"), 64)
	if err != nil {
		re.onError(w, err)
		return
	}
	lng, err := strconv.ParseFloat(r.URL.Query().Get("lng"), 64)
	if err != nil {
		re.onError(w, err)
		return
	}
	response, err := re.svc.OpenWeatherGetCurrentWeather(lat, lng)

	if err != nil {
		re.onError(w, err)
		return
	}
	json.NewEncoder(w).Encode(response)
}
