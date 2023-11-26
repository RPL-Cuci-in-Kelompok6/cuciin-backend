package entity

import (
	"time"

	"gorm.io/gorm"
)

type Partner struct {
	gorm.Model
	Name        string
	PhoneNumber string
	Email       string
	Password    string
	MapLink     string

	Services []Service
}

type Customer struct {
	gorm.Model
	Name     string
	Email    string
	Password string

	Orders []Order
}

type Service struct {
	gorm.Model
	Name  string
	Price uint64

	WashingMachines []WashingMachine
}

type WashingMachine struct {
	gorm.Model
	AvailableAt time.Time
	Brand       string
}

type Order struct {
	gorm.Model
	CustomerEmail string
	Status        string
	TotalPrice    uint64

	Payment Payment
	Machine WashingMachine
}

type Payment struct {
	gorm.Model
	Status bool
}
