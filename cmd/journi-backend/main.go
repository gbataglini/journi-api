package main

import (
	"net/http"
	"time"

	"github.com/gbataglini/journi-backend/internal/destination"
)


func main() {
	mux := http.NewServeMux()

	destinationStore := destination.NewStore()
	destinationService := destination.NewService(destinationStore)
	destinationRouter := destination.NewRest(destinationService)

	destinationRouter.Routes(mux)

	(&http.Server{
		Addr: ":8080",
		Handler: mux,
		ReadTimeout: 10 * time.Second,
	}).ListenAndServe()
}