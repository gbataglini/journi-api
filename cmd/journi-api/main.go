package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gbataglini/journi-api/domain"
	"github.com/gbataglini/journi-api/internal/config"
	"github.com/gbataglini/journi-api/internal/destination"
	"github.com/gbataglini/journi-api/internal/user"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"googlemaps.github.io/maps"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	cfg := config.Get()
	fmt.Println(cfg)
	mux := http.NewServeMux()

	googleClient, err := maps.NewClient(maps.WithAPIKey(cfg.GoogleApiKey))
	if err != nil {
		return fmt.Errorf("failed to init google client: %w", err)
	}

	postgresqlDbInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		cfg.DbHost, cfg.DbPort, cfg.DbUser, cfg.DbPassword, cfg.DbName)
	db, err := sqlx.Connect("postgres", postgresqlDbInfo)
	if err != nil {
		return err
	}
	defer db.Close()

	destinationStore := destination.NewStore(db)
	userStore := user.NewStore()

	destinationService := destination.NewService(destinationStore, googleClient)
	userService := user.NewService(userStore)

	destinationRouter := destination.NewRest(destinationService)
	userRouter := user.NewRest(userService)

	for _, router := range []domain.Router{
		destinationRouter,
		userRouter,
	} {
		router.Routes(mux)
	}

	return (&http.Server{
		Addr:        ":8080",
		Handler:     mux,
		ReadTimeout: 10 * time.Second,
	}).ListenAndServe()
}
