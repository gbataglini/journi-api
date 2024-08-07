package destination

import "github.com/gbataglini/journi-backend/domain"

type svc struct {
	store domain.DestinationStore
}

func NewService(store domain.DestinationStore) domain.DestinationService {
	return &svc{
		store: store,
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
