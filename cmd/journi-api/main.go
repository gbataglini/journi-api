package main

import (
	"cmp"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"googlemaps.github.io/maps"
)

type PlaceSuggestion struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {

	googleClient, err := maps.NewClient(maps.WithAPIKey(os.Getenv("JOURNIAPI_GOOGLEAPIKEY")))
	if err != nil {
		return fmt.Errorf("failed to init google client: %w", err)
	}

	openaiClient := openai.NewClient(
		option.WithAPIKey(os.Getenv("JOURNIAPI_GPTKEY")),
	)
	app := pocketbase.New()

	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		se.Router.GET("/v1/destinations/placesDetails", func(e *core.RequestEvent) error {
			searchParam := e.Request.URL.Query().Get("locationId")
			resp, err := googleClient.PlaceDetails(context.Background(), &maps.PlaceDetailsRequest{
				PlaceID: searchParam,
			})
			if err != nil {
				return err
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

			resp, err := googleClient.PlaceAutocomplete(e.Request.Context(), &maps.PlaceAutocompleteRequest{
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

		se.Router.GET("/v1/suggestions", func(e *core.RequestEvent) error {

			city := e.Request.URL.Query().Get("city")
			country := e.Request.URL.Query().Get("country")
			limit := e.Request.URL.Query().Get("limit")

			role := `You are a travel agent, working with yuppies look to get the lay of the land surrounded their dream vacay abode.
			You will mainly be working with Watchmojo style Top 5 queries.`

			prompt := fmt.Sprintf(`
			Provide a list of the top %s things to do in %s, %s.
			
			Focus on cultural attractions and places with local history. 
			
			Suggest 3 cultural attractions if available, and at least one café and one restaurant. If there aren't enough cultural attractions, suggest more popular restaurants or cafés.
			
			 You must answer in a JSON format, structured like so:
			
			 [
				{ "name":":name", "type": ":type", "description":":description" }
			 ]
			
			 Make sure the JSON is pretty printed. Only print the JSON. Don't print the markdown code formatting, ONLY PRINT JSON.

			`, cmp.Or(limit, "5"), city, country)

			fmt.Println("making query")
			completion, err := openaiClient.Chat.Completions.New(e.Request.Context(), openai.ChatCompletionNewParams{
				Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
					openai.UserMessage(prompt),
					openai.SystemMessage(role),
				}),
				Model: openai.F(openai.ChatModelGPT4oMini),
			})

			if err != nil {
				return err
			}

			suggestionSchema := []PlaceSuggestion{}
			err = json.Unmarshal([]byte(completion.Choices[0].Message.Content), &suggestionSchema)
			if err != nil {
				return err
			}

			return e.JSON(http.StatusOK, suggestionSchema)
		})

		return se.Next()
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
	return nil
}
