package destination

import "github.com/gbataglini/journi-backend/domain"

type store struct {
	dests []domain.Destination
}

func NewStore() domain.DestinationStore {
	return &store{
		dests: []domain.Destination{{
			ID: "1",
			City: "Pompeii",
			Country: "Italy",
			Visited: "false",
		}},
	}
}

func (s *store) ListDestinations() ([]domain.Destination, error) {
	return s.dests, nil
}

func (s *store) GetDestinationByID(destinationID string) (domain.Destination, error) {
	var destinationDetails domain.Destination
	for _, destination := range s.dests {
		if destination.ID == destinationID {
			destinationDetails = destination
			break
		}
	}
	return destinationDetails, nil
}

func (s *store) AddDestination(destination domain.Destination) (error) {
	s.dests = append(s.dests, destination)
	return nil
}

func (s *store) DeleteDestination(destinationID string) ([]domain.Destination, error) {
	for index, destination :=  range s.dests {
		if destination.ID == destinationID {
			s.dests = append(s.dests[:index], s.dests[index + 1:]...)
			break
		}
	}
	return s.dests, nil
}