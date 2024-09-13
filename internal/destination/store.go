package destination

import (
	"fmt"

	"github.com/gbataglini/journi-api/domain"
	"github.com/jmoiron/sqlx"
)

type store struct {
	dests []domain.Destination
	db    *sqlx.DB
}

func NewStore(db *sqlx.DB) domain.DestinationStore {
	return &store{
		db: db,
		dests: []domain.Destination{{
			ID:      1,
			City:    "Pompeii",
			Country: "Italy",
			Visited: "false",
		}},
	}
}

func (s *store) ListDestinations(userID int) ([]domain.Destination, error) {
	var dests []domain.Destination
	err := s.db.Select(&dests, "SELECT * FROM destinations WHERE user_id=$1", userID)
	if err != nil {
		return nil, fmt.Errorf("failed to query all destinations: %w", err)
	}
	return dests, nil
}

func (s *store) GetDestinationByID(destinationID int) (domain.Destination, error) {
	var destinationDetails domain.Destination
	err := s.db.Get(&destinationDetails, "SELECT * FROM destinations WHERE id=$1", destinationID)
	if err != nil {
		return domain.Destination{}, fmt.Errorf("failed to get destination: %w", err)
	}
	return destinationDetails, nil
}

func (s *store) AddDestination(destination domain.Destination) (domain.Destination, error) {
	insertStatement :=
		`INSERT INTO destinations (user_id, city, country, visited)
	VALUES ($1, $2, $3, $4)
	RETURNING id
	`
	id := 0
	err := s.db.QueryRow(insertStatement, destination.UserId, destination.City, destination.Country, destination.Visited).Scan(&id)
	if err != nil {
		return domain.Destination{}, fmt.Errorf("failed to add destination: %w", err)
	}

	fmt.Println("New record ID is:", id)
	return s.GetDestinationByID(id)
}

func (s *store) DeleteDestination(destinationID int) ([]domain.Destination, error) {
	for index, destination := range s.dests {
		if destination.ID == destinationID {
			s.dests = append(s.dests[:index], s.dests[index+1:]...)
			break
		}
	}
	return s.dests, nil
}
