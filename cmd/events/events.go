package events

import (
	"anpr-api/cmd/db"
	"anpr-api/cmd/structs"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func LastInHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Last in handler recieved - ", r.RequestURI)
	lastIn := db.FindLastEntry()
	jsonData, err := json.Marshal(lastIn)
	if err != nil {
		http.Error(w, "Error retrieving last in details", http.StatusBadRequest)
	}
	fmt.Fprint(w, string(jsonData))
	fmt.Println("Sent back - ", string(jsonData))
}

func EventsHandlerByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*") // Replace "*" with your desired origin(s)
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	fmt.Println("Event handler by ID recieved - ", r.RequestURI)
	id := r.URL.Path[len("/events/"):]
	if r.Method == "GET" {
		fmt.Println("Got a GET")
		var event structs.Event = db.GetEvent(id)
		jsonData, err := json.Marshal(event)
		if err != nil {
			http.Error(w, "Error retrieving Event details", http.StatusBadRequest)
		}
		fmt.Fprint(w, string(jsonData))
		fmt.Println("Sent back - ", string(jsonData))
	}
}

func EventsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*") // Replace "*" with your desired origin(s)
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if r.Method == "POST" {
		var event structs.Event
		err := json.NewDecoder(r.Body).Decode(&event)
		if err != nil {
			http.Error(w, "Error parsing JSON", http.StatusBadRequest)
			return
		}
		owner, company := db.FindOwnerByReg(event)
		db.WriteEvent(&structs.Event{
			Time:     time.Now(),
			Type:     event.Type,
			Reg:      event.Reg,
			DeviceID: event.DeviceID,
			Owner:    owner,
			Company:  company,
		})
		//	integrations.RunIntegrations(&event)
	}
	if r.Method == "GET" {
		events := db.ReadEvents()
		jsonData, err := json.Marshal(events)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Fprint(w, string(jsonData))
	}
}
