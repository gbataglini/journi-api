package user

import (
	"encoding/json"
	"net/http"

	"github.com/gbataglini/journi-backend/domain"
)


type rest struct {
	svc domain.UserService
}

func NewRest(service domain.UserService) domain.Router {
	return &rest{
		svc: service,
	}
}

func (re *rest) Routes(s * http.ServeMux) {
	s.HandleFunc("GET /api/v1/users", re.listUsers)
}

func (re *rest) listUsers(w http.ResponseWriter, r *http.Request) {
	users, err := re.svc.ListUsers()
	if err != nil {
		w.WriteHeader(500)
		return
	}
	
	json.NewEncoder(w).Encode(users)
}