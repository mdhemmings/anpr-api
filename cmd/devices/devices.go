package devices

import (
	"anpr-api/cmd/db"
	"anpr-api/cmd/structs"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func DeviceHandlerByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Device handler by ID recieved - ", r.RequestURI)
	id := r.URL.Path[len("/devices/"):]
	if r.Method == "GET" {
		fmt.Println("Got a GET")
		var event structs.Device = db.GetDevice(id)
		jsonData, err := json.Marshal(event)
		if err != nil {
			http.Error(w, "Error retrieving Device details", http.StatusBadRequest)
		}
		fmt.Fprint(w, string(jsonData))
		fmt.Println("Sent back - ", string(jsonData))
	}
}

func DeviceHandler(w http.ResponseWriter, r *http.Request) {
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
