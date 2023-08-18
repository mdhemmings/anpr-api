package db

import (
	"anpr-api/cmd/structs"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func getDB() (db *gorm.DB) {
	db, err := gorm.Open(sqlite.Open("anpr.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return
}

func ReadEvents() (events []structs.Event) {
	db := getDB()
	db.Find(&events)
	return
}

func InitDB() {
	db := getDB()
	db.AutoMigrate(&structs.Event{})
	db.AutoMigrate(&structs.Vehicle{})
	db.AutoMigrate(&structs.Integration{})
}

func WriteEvent(event *structs.Event) {
	db := getDB()
	db.Create(event)
}

func AddVehicle(vehicle *structs.Vehicle) {
	db := getDB()
	db.Create(vehicle)
}

func ListVehicles() (vehicles []structs.Vehicle) {
	db := getDB()
	db.Find(&vehicles)
	return
}

func GetVehicle(id string) (vehicle structs.Vehicle) {
	db := getDB()
	db.Where("Id = ?", id).First(&vehicle)
	return
}

func DeleteVehicle(id string) {
	var vehicle structs.Vehicle
	db := getDB()
	db.Where("Id = ?", id).Delete(&vehicle)
}

func UpdateVehicle(vehicle *structs.Vehicle) {
	db := getDB()
	var currentVehicle *structs.Vehicle
	if err := db.First(&currentVehicle, vehicle.Id).Error; err != nil {
		fmt.Println("Error fetching vehicle:", err)
		return
	}
	currentVehicle = vehicle
	if err := db.Save(&currentVehicle).Error; err != nil {
		fmt.Println("Error updating vehicle:", err)
		return
	}
	fmt.Println("Vehicle updated - ", vehicle.Id)
}

func GetEvent(id string) (event structs.Event) {
	db := getDB()
	db.Where("Id = ?", id).First(&event)
	return
}

func AddIntegration(integration *structs.Integration) {
	db := getDB()
	db.Create(integration)
}

func ListIntegrations() (integrations []structs.Integration) {
	db := getDB()
	db.Find(&integrations)
	return
}

func GetIntegrationsByType(EventType string) (integrations []structs.Integration) {
	db := getDB()
	db.Where("Type contains ?", EventType).Find(&integrations)
	return
}

func GetIntegration(id string) (integration structs.Integration) {
	db := getDB()
	db.Where("Id = ?", id).First(&integration)
	return
}

func DeleteIntegration(id string) {
	var integration structs.Integration
	db := getDB()
	db.Where("Id = ?", id).Delete(&integration)
}

func UpdateIntegration(integration *structs.Integration) {
	db := getDB()
	var currentIntegration *structs.Integration
	if err := db.First(&currentIntegration, integration.Id).Error; err != nil {
		fmt.Println("Error fetching integration:", err)
		return
	}
	currentIntegration = integration
	if err := db.Save(&currentIntegration).Error; err != nil {
		fmt.Println("Error updating integration:", err)
		return
	}
	fmt.Println("Integration updated - ", integration.Id)
}

func FindLastEntry() (LastInEntry structs.LastInEntry) {
	var event structs.Event
	db := getDB()
	db.Where("Type = ?", "Entry").Where("Owner != ?", "").Order("time DESC").Find(&event)
	fmt.Println(event)
	LastInEntry = structs.LastInEntry{
		Owner:   event.Owner,
		Company: event.Company,
	}
	return
}

func FindOwnerByReg(event structs.Event) (owner, company string) {
	db := getDB()
	var vehicle structs.Vehicle
	db.Where("Reg = ?", event.Reg).First(&vehicle)
	fmt.Println(vehicle)
	owner = vehicle.Owner
	company = vehicle.Company
	return
}
