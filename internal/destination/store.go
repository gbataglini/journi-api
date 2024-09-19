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
			Visited: false,
		}},
	}
}

func (s *store) ListDestinations(userID int) ([]domain.Destination, error) {
	var dests []domain.DestinationDAO
	var formattedDests []domain.Destination
	err := s.db.Select(&dests, "SELECT * FROM destinations WHERE user_id=$1", userID)
	if err != nil {
		return nil, fmt.Errorf("failed to query all destinations: %w", err)
	}

	for _, dest := range dests {
		formattedDests = append(formattedDests, dest.ToDestination())
	}

	return formattedDests, nil
}

func (s *store) ListCountries(userID int) ([]domain.Country, error) {
	var dests []domain.Country
	err := s.db.Select(&dests,
		`SELECT country
		 ,COUNT(CASE WHEN visited = TRUE THEN 1 END) AS visited
		 ,ARRAY_AGG(DISTINCT destination_type) AS destination_type
		 FROM destinations
		 GROUP BY country`)
	if err != nil {
		return nil, fmt.Errorf("failed to query all countries: %w", err)
	}
	return dests, nil
}

func (s *store) GetDestinationByID(destinationID int) (domain.Destination, error) {
	var destDetails domain.DestinationDAO
	err := s.db.Get(&destDetails, "SELECT * FROM destinations WHERE id=$1", destinationID)
	if err != nil {
		return domain.Destination{}, fmt.Errorf("failed to get destination: %w", err)
	}
	return destDetails.ToDestination(),
		nil
}

func (s *store) AddDestination(destination domain.Destination) (domain.Destination, error) {
	insertStatement :=
		`INSERT INTO destinations (user_id, city, country, visited, destination_type, googlemaps_id, lat, lng)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	RETURNING id
	`
	id := 0
	err := s.db.QueryRow(insertStatement, destination.UserId, destination.City, destination.Country, destination.Visited, destination.DestinationType, destination.GoogleMapsId, destination.Location.Lat, destination.Location.Lng).Scan(&id)
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
