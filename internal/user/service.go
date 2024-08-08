package user

import "github.com/gbataglini/journi-api/domain"

type svc struct {
	store domain.UserStore
}

func NewService(store domain.UserStore) domain.UserService {
	return &svc{
		store: store,
	}
}

func (s *svc) ListUsers() ([]domain.User, error) {
	return s.store.ListUsers()
}
