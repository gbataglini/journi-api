package user

import "github.com/gbataglini/journi-backend/domain"

type store struct {
	users []domain.User
}

func NewStore() domain.UserStore {
	return &store{
		users: []domain.User{{
			ID:        "1",
			FirstName: "Ducko",
			LastName:  "Quackins",
			Email:     "ducko@qua.ck",
		}},
	}
}

func (s *store) ListUsers() ([]domain.User, error) {
	return s.users, nil
}
