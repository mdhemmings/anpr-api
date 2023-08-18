package structs

import "time"

type LastInEntry struct {
	Owner   string `json:"Owner"`
	Company string `json:"Company"`
}

type Vehicle struct {
	Id      int    `json:"id" gorm:"unique"`
	Reg     string `json:"Reg"`
	Type    string `json:"Type"`
	Owner   string `json:"Owner"`
	Company string `json:"Company"`
	Allowed bool   `json:"Allowed"`
}

type Event struct {
	Id       int       `json:"id" gorm:"unique"`
	Time     time.Time `json:"Time"`
	Type     string    `json:"Type"`
	Reg      string    `json:"Reg"`
	DeviceID string    `json:"DeviceID"`
	Owner    string    `json:"Owner"`
	Company  string    `json:"Company"`
}

type Integration struct {
	Id          int    `json:"id" gorm:"unique"`
	Type        string `json:"Type"`
	EventType   string `json:"EventType"`
	Destination string `json:"Destination"`
}
