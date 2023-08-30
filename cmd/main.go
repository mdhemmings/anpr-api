package main

import (
	"log"
	"net/http"

	"anpr-api/cmd/auth"
	"anpr-api/cmd/db"
	"anpr-api/cmd/devices"
	"anpr-api/cmd/events"
	"anpr-api/cmd/integrations"
	"anpr-api/cmd/vehicles"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	log.Println("Started")
	db.InitDB()
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
	})
	router := mux.NewRouter().StrictSlash(false)
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Log headers before CORS middleware
			log.Println("Headers before CORS middleware:", w.Header())

			// Call the next middleware
			next.ServeHTTP(w, r)

			// Log headers after CORS middleware
			log.Println("Headers after CORS middleware:", w.Header())
		})
	})
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Log the request method and path
			log.Printf("Request: %s %s", r.Method, r.URL.Path)

			// Log request headers
			log.Println("Request Headers:")
			for name, values := range r.Header {
				for _, value := range values {
					log.Printf("%s: %s", name, value)
				}
			}

			// Call the next middleware
			next.ServeHTTP(w, r)
		})
	})
	router.Methods(http.MethodOptions).Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers for preflight requests
		w.Header().Set("Access-Control-Allow-Origin", "*") // Replace "*" with your desired origin(s)
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Respond with a 200 OK status for preflight requests
		w.WriteHeader(http.StatusOK)
		log.Println("Got pre-flight request")
	}))
	router.Use(auth.AuthMiddleware)
	router.Use(auth.CorsMiddleware)

	router.Handle("/events", http.HandlerFunc(events.EventsHandler))
	router.Handle("/events/{id}", http.HandlerFunc(events.EventsHandlerByID))
	router.Handle("/vehicles", http.HandlerFunc(vehicles.VehicleHandler))
	router.Handle("/vehicles/{id}", http.HandlerFunc(vehicles.VehicleHandlerByID))
	router.Handle("/integrations", http.HandlerFunc(integrations.IntegrationHandler))
	router.Handle("/integrations/{id}", http.HandlerFunc(integrations.IntegrationsHandlerByID))
	router.Handle("/lastin", http.HandlerFunc(events.LastInHandler))
	router.Handle("/devices", http.HandlerFunc(devices.DeviceHandler))
	router.Handle("/devices/{id}", http.HandlerFunc(devices.DeviceHandlerByID))

	handler := c.Handler(router)

	optionsHandler := cors.AllowAll().Handler(nil)
	router.Methods(http.MethodOptions).Handler(optionsHandler)
	log.Fatal(http.ListenAndServe(":8080", handler))
}
