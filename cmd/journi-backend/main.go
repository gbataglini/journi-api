package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gbataglini/journi-backend/domain"
	"github.com/gbataglini/journi-backend/internal/config"
	"github.com/gbataglini/journi-backend/internal/destination"
	"github.com/gbataglini/journi-backend/internal/user"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	cfg := config.Get()
	fmt.Println(cfg)
	mux := http.NewServeMux()
	postgresqlDbInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		cfg.DbHost, cfg.DbPort, cfg.DbUser, cfg.DbPassword, cfg.DbName)
	db, err := sqlx.Connect("postgres", postgresqlDbInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	destinationStore := destination.NewStore(db)
	userStore := user.NewStore()

	destinationService := destination.NewService(destinationStore)
	userService := user.NewService(userStore)

	destinationRouter := destination.NewRest(destinationService)
	userRouter := user.NewRest(userService)

	for _, router := range []domain.Router{
		destinationRouter,
		userRouter,
	} {
		router.Routes(mux)
	}

	(&http.Server{
		Addr:        ":8080",
		Handler:     mux,
		ReadTimeout: 10 * time.Second,
	}).ListenAndServe()
}
