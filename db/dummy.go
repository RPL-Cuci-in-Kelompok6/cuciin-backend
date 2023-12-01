package db

import (
	"time"

	"github.com/RPL-Cuci-in-Kelompok6/cuciin-backend/entity"
)

func insertDataDummy() {
	// Data Dummy: Customer
	GetConnection().Create([]*entity.Customer{
		{
			Name:        "John Doe",
			PhoneNumber: "+6218378261",
			Email:       "johndoe@gmail.com",
			Password:    "test123",
		},
		{
			Name:        "Alice Doe",
			PhoneNumber: "+62183732891",
			Email:       "alicedoe@gmail.com",
			Password:    "test123",
		},
		{
			Name:        "Claier Ecler",
			PhoneNumber: "+621837826111",
			Email:       "claierecler@gmail.com",
			Password:    "test123",
		},
	})

	// Data Dummy: Partner
	partners := []*entity.Partner{
		{
			Name:        "Dodo Laundry",
			PhoneNumber: "+621837826131232",
			Email:       "dodolaundry@gmail.com",
			Password:    "test123",
			MapLink:     "https://maps.app.goo.gl/c2VnSd9FauWmL9cR6",
		},
		{
			Name:        "Tica Laundry",
			PhoneNumber: "+62183782621321",
			Email:       "ticalaundry@gmail.com",
			Password:    "test123",
			MapLink:     "https://maps.app.goo.gl/UCWoNJgBzVyQ8X8k9",
		},
	}
	GetConnection().Create(partners)

	// Data Dummy: Service
	var dodoLaundry entity.Partner
	var ticaLaundry entity.Partner
	GetConnection().Table("partners").Where(&entity.Partner{Name: "Dodo Laundry"}).Select("id").First(&dodoLaundry)
	GetConnection().Table("partners").Where(&entity.Partner{Name: "Tica Laundry"}).Select("id").First(&ticaLaundry)
	var services = []*entity.Service{
		{
			Name:      "Cuci 7kg - W1",
			Price:     10000,
			PartnerID: dodoLaundry.ID,
		},
		{
			Name:      "Cuci 7kg - W2",
			Price:     10000,
			PartnerID: dodoLaundry.ID,
		},
		{
			Name:      "Cuci 7kg - W1",
			Price:     10000,
			PartnerID: ticaLaundry.ID,
		},
		{
			Name:      "Cuci 7kg - W2",
			Price:     10000,
			PartnerID: ticaLaundry.ID,
		},
		{
			Name:      "Kering 25 Menit - D1",
			Price:     15000,
			PartnerID: dodoLaundry.ID,
		},
		{
			Name:      "Kering 25 Menit - D2",
			Price:     15000,
			PartnerID: dodoLaundry.ID,
		},
		{
			Name:      "Kering 25 Menit - D1",
			Price:     15000,
			PartnerID: ticaLaundry.ID,
		},
	}
	GetConnection().Create(services)

	// Data Dummy: Washing Machine
	washingMachine := []*entity.WashingMachine{
		{
			AvailableAt: time.Now(),
			Brand:       "LG",
		},
		{
			AvailableAt: time.Now(),
			Brand:       "Samsung",
		},
		{
			AvailableAt: time.Now(),
			Brand:       "LG",
		},
		{
			AvailableAt: time.Now(),
			Brand:       "Samsung",
		},
	}
	for i, w := range washingMachine {
		var x entity.Service
		result := GetConnection().Where(&entity.Service{Name: services[i].Name}).First(&x)
		w.ServiceID = x.ID
		if err := result.Error; err != nil {
			panic(err.Error())
		}
	}
	GetConnection().Create(washingMachine)
}
