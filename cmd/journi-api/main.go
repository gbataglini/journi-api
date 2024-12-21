package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"googlemaps.github.io/maps"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {

	googleClient, err := maps.NewClient(maps.WithAPIKey(""))
	if err != nil {
		return fmt.Errorf("failed to init google client: %w", err)
	}

	   app := pocketbase.New()

    app.OnServe().BindFunc(func(se *core.ServeEvent) error {
        se.Router.GET("/v1/destinations/placesDetails", func(e *core.RequestEvent) error {
			searchParam := e.Request.URL.Query().Get("locationId")
				resp, err := googleClient.PlaceDetails(context.Background(), &maps.PlaceDetailsRequest{
				PlaceID: searchParam,
			})
			if err != nil {
				return  err
			}
			return e.JSON(http.StatusOK, resp)
		})

		se.Router.GET("/v1/destinations/autocomplete", func(e *core.RequestEvent) error {
			searchParam := e.Request.URL.Query().Get("searchParam")
			resp, err := googleClient.PlaceAutocomplete(context.Background(), &maps.PlaceAutocompleteRequest{
				Input: searchParam,
				Types: "(cities)",
			})
			if err != nil {
				return err
			}
			return e.JSON(http.StatusOK, resp)
		})

		se.Router.GET("/v1/destinations/establishmentSearch", func(e *core.RequestEvent) error {

			searchParam := e.Request.URL.Query().Get("searchParam")
			lat := e.Request.URL.Query().Get("lat")
			lng := e.Request.URL.Query().Get("lng")
			
			
			latlng := fmt.Sprintf("%s,%s", lat, lng)
			formatLatLng, err := maps.ParseLatLng(latlng)

			if err != nil {
				return err
			}

			resp, err := googleClient.PlaceAutocomplete(context.Background(), &maps.PlaceAutocompleteRequest{
				Input:        searchParam,
				Types:        "establishment",
				Location:     &formatLatLng,
				Radius:       80500,
				StrictBounds: true,
			})
			if err != nil {
				return err
			}
			return e.JSON(http.StatusOK, resp)
		})

        return se.Next()
    })

    if err := app.Start(); err != nil {
        log.Fatal(err)
    }
	return nil 
}
