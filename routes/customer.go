package routes

import (
	"net/http"
	"time"

	"github.com/RPL-Cuci-in-Kelompok6/cuciin-backend/db"
	"github.com/RPL-Cuci-in-Kelompok6/cuciin-backend/entity"
	"github.com/gin-gonic/gin"
)

type PartnerServiceBody struct {
	PartnerID uint `json:"id-mitra" binding:"required"`
}

type PartnerResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"nama"`
	PhoneNumber string `json:"telepon"`
	Email       string `json:"email"`
	MapLink     string `json:"link-map"`
}

type MachinesResponse struct {
	ID          uint      `json:"id"`
	AvailableAt time.Time `json:"waktu-tersedia"`
	Brand       string    `json:"merk"`
}

type ServiceResponse struct {
	ID              uint               `json:"id"`
	Name            string             `json:"nama"`
	Price           uint64             `json:"harga"`
	WashingMachines []MachinesResponse `json:"mesin"`
}

func GetPartners() gin.HandlerFunc {
	return func(c *gin.Context) {
		var partners []entity.Partner
		if err := db.GetConnection().Select("id", "name", "phone_number", "email", "map_link").Limit(20).Find(&partners).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Failed to get partners, please try again later.",
			})
			return
		}

		var partnerResponse []*PartnerResponse

		for _, x := range partners {
			newPartner := PartnerResponse{
				ID:          x.ID,
				Name:        x.Name,
				Email:       x.Email,
				PhoneNumber: x.PhoneNumber,
				MapLink:     x.MapLink,
			}
			partnerResponse = append(partnerResponse, &newPartner)
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data": gin.H{
				"mitra": partnerResponse,
			},
		})
	}
}

func GetServicesByPartner() gin.HandlerFunc {
	return func(c *gin.Context) {
		var serviceRequest PartnerServiceBody
		if err := bindBodyOrError(c, &serviceRequest); err != nil {
			return
		}

		var services []*entity.Service

		if err := db.GetConnection().Where("partner_id = ?", serviceRequest.PartnerID).Preload("WashingMachines").Find(&services).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Failed to get services, please try again later",
			})
			return
		}

		var response []*ServiceResponse

		for _, x := range services {
			row := &ServiceResponse{
				ID:    x.ID,
				Name:  x.Name,
				Price: x.Price,
			}
			var machines []MachinesResponse
			for _, m := range x.WashingMachines {
				mrow := MachinesResponse{
					ID:          m.ID,
					AvailableAt: m.AvailableAt,
					Brand:       m.Brand,
				}
				machines = append(machines, mrow)
			}
			row.WashingMachines = machines
			response = append(response, row)
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data": gin.H{
				"layanan": response,
			},
		})
	}
}

type OrderQuery struct {
	CustomerID uint `json:"id-customer" binding:"required"`
}

type OrderResponse struct {
	ID         int    `json:"id"`
	Status     string `json:"status"`
	TotalPrice uint64 `json:"total-harga"`
	Paid       bool   `json:"sudah-bayar"`
	Machine    int    `josn:"id-mesin"`
}

func GetOrdersByCustomer() gin.HandlerFunc {
	return func(c *gin.Context) {
		var query OrderQuery
		if err := bindBodyOrError(c, &query); err != nil {
			return
		}

		var orders []*entity.Order

		if result := db.GetConnection().Preload("Payment").Where(&entity.Order{CustomerID: query.CustomerID}).Find(&orders); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Failed to get orders, please try again later",
			})
			return
		}

		var response []OrderResponse
		for _, order := range orders {
			response = append(response, OrderResponse{
				ID:         int(order.ID),
				Status:     order.Status,
				Paid:       order.Payment.Status,
				TotalPrice: order.TotalPrice,
				Machine:    int(order.Machine.ID),
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data": gin.H{
				"pesanan": response,
			},
		})
	}
}
