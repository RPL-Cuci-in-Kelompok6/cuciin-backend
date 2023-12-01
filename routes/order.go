package routes

import (
	"net/http"

	"github.com/RPL-Cuci-in-Kelompok6/cuciin-backend/db"
	"github.com/RPL-Cuci-in-Kelompok6/cuciin-backend/entity"
	"github.com/gin-gonic/gin"
)

type OrderBody struct {
	CustomerEmail string `json:"customer_email"`
	TotalPrice    uint64 `json:"total_price"`
	MachineID     uint   `json:"machine_id"`
	CustomerID    uint   `json:"customer_id"`
}

type PayBody struct {
	OrderID uint `json:"order_id"`
}

func CreateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		var order OrderBody
		if err := bindBodyOrError(c, &order); err != nil {
			return
		}

		newOrder := &entity.Order{
			CustomerEmail: order.CustomerEmail,
			TotalPrice:    order.TotalPrice,
			CustomerID:    order.CustomerID,
			MachineID:     order.MachineID,
			PaymentStatus: PAYMENT_NOTPAID,
			Status:        ORDER_FILLED,
		}

		result := db.GetConnection().Create(newOrder)

		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": "false",
				"message": "Failed to create order. Please try again later",
			})
		}
	}
}

func PayOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		var pay PayBody
		if err := bindBodyOrError(c, &pay); err != nil {
			return
		}
		result := db.GetConnection().Table("orders").Where("id = ?", pay.OrderID).Update("payment_status", PAYMENT_PAID)

		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Failed to verify your payment. Please try again later",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Your payment has been received.",
		})
	}
}

func CancelOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		var pay PayBody
		if err := bindBodyOrError(c, &pay); err != nil {
			return
		}
		result := db.GetConnection().Table("orders").Where("id = ?", pay.OrderID).Update("status", ORDER_CANCELED)

		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Failed to update your order. Please try again later",
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Your order has been canceled",
		})
	}
}
