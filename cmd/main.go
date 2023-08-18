package main

import (
	"log"
	"net/http"

	"anpr-api/cmd/db"
	"anpr-api/cmd/events"
	"anpr-api/cmd/integrations"
	"anpr-api/cmd/vehicles"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	db.InitDB()
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
	})
	router := mux.NewRouter().StrictSlash(false)
	router.HandleFunc("/events", events.EventsHandler)
	router.HandleFunc("/events/{id}", events.EventsHandlerByID)
	router.HandleFunc("/vehicles", vehicles.VehicleHandler)
	router.HandleFunc("/vehicles/{id}", vehicles.VehicleHandlerByID)
	router.HandleFunc("/integrations", integrations.IntegrationHandler)
	router.HandleFunc("/integrations/{id}", integrations.IntegrationsHandlerByID)
	router.HandleFunc("/lastin", events.LastInHandler)
	handler := c.Handler(router)
	log.Fatal(http.ListenAndServe(":8080", handler))
}
