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
	s.HandleFunc("GET /api/v1/destinations", re.listDestinations)
	s.HandleFunc("GET /api/v1/destinations/{id}", re.getByID)
	s.HandleFunc("POST /api/v1/destinations", re.create)
	s.HandleFunc("DELETE /api/v1/destinations/{id}", re.delete)
	s.HandleFunc("GET /api/v1/destinations/autocomplete", re.placesAutocompleteAPI)
}

func (re *rest) listDestinations(w http.ResponseWriter, r *http.Request) {
	dests, err := re.svc.ListDestinations()
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
		fmt.Println(err)
	}
	re.svc.DeleteDestination(destinationID)
	dests, err := re.svc.ListDestinations()
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
