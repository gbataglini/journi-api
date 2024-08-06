package destination

import (
	"encoding/json"
	"net/http"

	"github.com/gbataglini/journi-backend/domain"
)

type rest struct{
	svc domain.DestinationService
}

func NewRest(service domain.DestinationService) domain.Router {
	return &rest{
		svc: service,
	}
}

func (re *rest) Routes(s * http.ServeMux) {
	s.HandleFunc("GET /api/v1/destinations", re.listDestinations)
	s.HandleFunc("GET /api/v1/destinations/{id}", re.getByID)
	s.HandleFunc("POST /api/v1/destinations", re.create)
	s.HandleFunc("DELETE /api/v1/destinations/{id}", re.delete)
}

func (re * rest) listDestinations(w http.ResponseWriter, r *http.Request) {
	dests, err := re.svc.ListDestinations()
	if err != nil {
		w.WriteHeader(500)
		return
	}

	json.NewEncoder(w).Encode(dests)
} 

func (re * rest) create(w http.ResponseWriter, r *http.Request) { 
	var newDestination domain.Destination
	json.NewDecoder(r.Body).Decode(&newDestination)
	re.svc.AddDestination(newDestination)
	json.NewEncoder(w).Encode(newDestination)
}

func (re * rest) delete(w http.ResponseWriter, r *http.Request) {
	destinationID := r.PathValue("id") 
	re.svc.DeleteDestination(destinationID)
	dests, err := re.svc.ListDestinations()
		if err != nil {
		w.WriteHeader(500)
		return
	}
	json.NewEncoder(w).Encode(dests)
}

func (re * rest) getByID(w http.ResponseWriter, r *http.Request) {
	destinationID := r.PathValue("id") 
	selectedDest, err := re.svc.GetDestinationByID(destinationID)
		if err != nil {
			w.WriteHeader(500)
			return
		}

	json.NewEncoder(w).Encode(selectedDest)
}