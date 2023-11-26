package entity

import (
	"time"

	"gorm.io/gorm"
)

type Partner struct {
	gorm.Model
	Name        string
	PhoneNumber string `gorm:"uniqueIndex"`
	Email       string `gorm:"uniqueIndex"`
	Password    string
	MapLink     string

	Services []Service
}

type Customer struct {
	gorm.Model
	Name        string
	PhoneNumber string `gorm:"uniqueIndex"`
	Email       string `gorm:"uniqueIndex"`
	Password    string

	Orders []Order
}

type Service struct {
	gorm.Model
	Name  string
	Price uint64

	WashingMachines []WashingMachine

	PartnerID uint
}

type WashingMachine struct {
	gorm.Model
	AvailableAt time.Time
	Brand       string

	ServiceID uint
}

type Order struct {
	gorm.Model
	CustomerEmail string
	Status        string
	TotalPrice    uint64

	Payment Payment

	MachineID uint
	Machine   WashingMachine

	CustomerID uint
}

type Payment struct {
	gorm.Model
	Status bool

	OrderID uint
}
