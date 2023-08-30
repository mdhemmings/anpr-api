package db

import (
	"anpr-api/cmd/structs"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getDB() (db *gorm.DB, err error) {
	dbtype := os.Getenv("DBTYPE")
	log.Println(dbtype)
	switch dbtype {
	case "postgres":
		environmentVariables := []string{"DBHOST", "DBUSER", "DBPASSWORD", "DBPORT", "DBNAME"}
		envValues := make(map[string]string)
		for _, envVar := range environmentVariables {
			envValue := os.Getenv(envVar)
			if envVar == "" {
				log.Println("Environment variable not set - ", envVar)
				return
			}
			envValues[envVar] = envValue
		}
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s port=%s dbname=%s sslmode=disable TimeZone=Europe/London",
			envValues["DBHOST"],
			envValues["DBUSER"],
			envValues["DBPASSWORD"],
			envValues["DBPORT"],
			envValues["DBNAME"],
		)
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Println("Error opening database", err)
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unsupported database type: %s", dbtype)
	}

	return db, nil
}

func ReadEvents() (events []structs.Event) {
	db, err := getDB()
	if err != nil {
		log.Println(err)
	}
	db.Find(&events)
	return
}

func InitDB() {
	db, err := getDB()
	if err != nil {
		log.Println(err)
	}
	db.AutoMigrate(&structs.Event{})
	db.AutoMigrate(&structs.Vehicle{})
	db.AutoMigrate(&structs.Integration{})
	db.AutoMigrate(&structs.Device{})
}

func WriteEvent(event *structs.Event) {
	db, err := getDB()
	if err != nil {
		log.Println(err)
	}
	db.Create(event)
}

func AddVehicle(vehicle *structs.Vehicle) {
	db, err := getDB()
	if err != nil {
		log.Println(err)
	}
	db.Create(vehicle)
}

func ListVehicles() (vehicles []structs.Vehicle) {
	db, err := getDB()
	if err != nil {
		log.Println(err)
	}
	db.Find(&vehicles)
	return
}

func GetVehicle(id string) (vehicle structs.Vehicle) {
	db, err := getDB()
	if err != nil {
		log.Println(err)
	}
	db.Where("Id = ?", id).First(&vehicle)
	return
}

func DeleteVehicle(id string) {
	var vehicle structs.Vehicle
	db, err := getDB()
	if err != nil {
		log.Println(err)
	}
	db.Where("Id = ?", id).Delete(&vehicle)
}

func UpdateVehicle(vehicle *structs.Vehicle) {
	db, err := getDB()
	if err != nil {
		log.Println(err)
	}
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

func ReadDevices() (devices []structs.Device) {
	db, err := getDB()
	if err != nil {
		log.Println(err)
	}
	db.Find(&devices)
	return
}

func GetDevice(id string) (device structs.Device) {
	db, err := getDB()
	if err != nil {
		log.Println(err)
	}
	db.Where("Id = ?", id).First(&device)
	return
}

func UpdateDevice(device *structs.Device) {
	db, err := getDB()
	if err != nil {
		log.Println(err)
	}
	var currentDevice *structs.Device
	if err := db.First(&currentDevice, device.Id).Error; err != nil {
		fmt.Println("Error fetching vehicle:", err)
		return
	}
	currentDevice = device
	if err := db.Save(&currentDevice).Error; err != nil {
		fmt.Println("Error updating vehicle:", err)
		return
	}
	fmt.Println("Vehicle updated - ", device.Id)
}

func GetEvent(id string) (event structs.Event) {
	db, err := getDB()
	if err != nil {
		log.Println(err)
	}
	db.Where("Id = ?", id).First(&event)
	return
}

func AddIntegration(integration *structs.Integration) {
	db, err := getDB()
	if err != nil {
		log.Println(err)
	}
	db.Create(integration)
}

func ListIntegrations() (integrations []structs.Integration) {
	db, err := getDB()
	if err != nil {
		log.Println(err)
	}
	db.Find(&integrations)
	return
}

func GetIntegrationsByType(EventType string) (integrations []structs.Integration) {
	db, err := getDB()
	if err != nil {
		log.Println(err)
	}
	db.Where("Type contains ?", EventType).Find(&integrations)
	return
}

func GetIntegration(id string) (integration structs.Integration) {
	db, err := getDB()
	if err != nil {
		log.Println(err)
	}
	db.Where("Id = ?", id).First(&integration)
	return
}

func DeleteIntegration(id string) {
	var integration structs.Integration
	db, err := getDB()
	if err != nil {
		log.Println(err)
	}
	db.Where("Id = ?", id).Delete(&integration)
}

func UpdateIntegration(integration *structs.Integration) {
	db, err := getDB()
	if err != nil {
		log.Println(err)
	}
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
	db, err := getDB()
	if err != nil {
		log.Println(err)
	}
	db.Where("Type = ?", "Entry").Where("Owner != ?", "").Order("time DESC").Find(&event)
	fmt.Println(event)
	LastInEntry = structs.LastInEntry{
		Owner:   event.Owner,
		Company: event.Company,
	}
	return
}

func FindOwnerByReg(event structs.Event) (owner, company string) {
	db, err := getDB()
	if err != nil {
		log.Println(err)
	}
	var vehicle structs.Vehicle
	db.Where("Reg = ?", event.Reg).First(&vehicle)
	fmt.Println(vehicle)
	owner = vehicle.Owner
	company = vehicle.Company
	return
}
