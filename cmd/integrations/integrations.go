package integrations

import (
	"anpr-api/cmd/db"
	"anpr-api/cmd/structs"
	"encoding/json"
	"fmt"
	"net/http"
)

//func RunIntegrations(event *structs.Event) {
//	integrations := db.GetIntegrationsByType(event.Type)
//	for _, integration := range integrations {
//		fmt.Println(integration)
//		switch integration.Type {
//		case "email":
//			fmt.Println("Would have sent an email to ", integration.Destination)
//		case "slack":
//			fmt.Println("Would have sent a slack message to ", integration.Destination)
//		case "relay":
//			fmt.Println("Would have changed state of relay ", integration.Destination)
//		case "webhook":
//			fmt.Println("Would have sent webhook to ", integration.Destination)
//		}
//	}
//}

func IntegrationsHandlerByID(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Integration hander by ID recieved - %v \n", r.RequestURI)
	id := r.URL.Path[len("/integrations/"):]
	if r.Method == "DELETE" {
		fmt.Println("Got a delete")
		db.DeleteIntegration(id)
	}
	if r.Method == "GET" {
		fmt.Println("Got a GET")
		var integration structs.Integration = db.GetIntegration(id)
		jsonData, err := json.Marshal(integration)
		if err != nil {
			http.Error(w, "Error retrieving integration details", http.StatusBadRequest)
		}
		fmt.Fprint(w, string(jsonData))
		fmt.Println("Sent back - ", string(jsonData))
	}
	if r.Method == "PATCH" {
		fmt.Println("Got a PATCH")
		var updatedintegration structs.Integration
		err := json.NewDecoder(r.Body).Decode(&updatedintegration)
		if err != nil {
			http.Error(w, "Error parsing JSON", http.StatusBadRequest)
			return
		}
		db.UpdateIntegration(&updatedintegration)
	}
}

func IntegrationHandler(w http.ResponseWriter, r *http.Request) {
	db.InitDB()
	fmt.Printf("Integration hander recieved - %v\n", r.RequestURI)
	if r.Method == "POST" {
		var integration structs.Integration
		err := json.NewDecoder(r.Body).Decode(&integration)
		if err != nil {
			http.Error(w, "Error parsing JSON", http.StatusBadRequest)
			return
		}
		db.AddIntegration(&integration)
	}
	if r.Method == "GET" {
		var integrations []structs.Integration = db.ListIntegrations()
		jsonData, err := json.Marshal(integrations)
		if err != nil {
			http.Error(w, "Error retrieving integration list", http.StatusBadRequest)
		}
		fmt.Fprint(w, string(jsonData))
	}

}
