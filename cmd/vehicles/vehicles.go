package vehicles

import (
	"anpr-api/cmd/db"
	"anpr-api/cmd/structs"
	"encoding/json"
	"fmt"
	"net/http"
)

func VehicleHandlerByID(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Vehicle hander by ID recieved - %v \n", r.RequestURI)
	id := r.URL.Path[len("/vehicles/"):]
	if r.Method == "DELETE" {
		fmt.Println("Got a delete")
		db.DeleteVehicle(id)
	}
	if r.Method == "GET" {
		fmt.Println("Got a GET")
		var vehicle structs.Vehicle = db.GetVehicle(id)
		jsonData, err := json.Marshal(vehicle)
		if err != nil {
			http.Error(w, "Error retrieving vehicle details", http.StatusBadRequest)
		}
		fmt.Fprint(w, string(jsonData))
		fmt.Println("Sent back - ", string(jsonData))
	}
	if r.Method == "PATCH" {
		fmt.Println("Got a PATCH")
		var updatedVehicle structs.Vehicle
		err := json.NewDecoder(r.Body).Decode(&updatedVehicle)
		if err != nil {
			http.Error(w, "Error parsing JSON", http.StatusBadRequest)
			return
		}
		db.UpdateVehicle(&updatedVehicle)
	}
}

func VehicleHandler(w http.ResponseWriter, r *http.Request) {
	db.InitDB()
	fmt.Printf("Vehicle hander recieved - %v\n", r.RequestURI)
	if r.Method == "POST" {
		var vehicle structs.Vehicle
		err := json.NewDecoder(r.Body).Decode(&vehicle)
		if err != nil {
			http.Error(w, "Error parsing JSON", http.StatusBadRequest)
			return
		}
		db.AddVehicle(&vehicle)
	}
	if r.Method == "GET" {
		var vehicles []structs.Vehicle = db.ListVehicles()
		jsonData, err := json.Marshal(vehicles)
		if err != nil {
			http.Error(w, "Error retrieving vehicle list", http.StatusBadRequest)
		}
		fmt.Fprint(w, string(jsonData))
	}

}
