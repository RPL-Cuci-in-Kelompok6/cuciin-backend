package routes

import (
	"net/http"

	"github.com/RPL-Cuci-in-Kelompok6/cuciin-backend/db"
	"github.com/RPL-Cuci-in-Kelompok6/cuciin-backend/entity"
	"github.com/gin-gonic/gin"
)

type OrderBody struct {
	CustomerEmail string `json:"email" binding:"required"`
	TotalPrice    uint64 `json:"harga" binding:"required"`
	MachineID     uint   `json:"id-mesin" binding:"required"`
	CustomerID    uint   `json:"id-customer" binding:"required"`
}

type PayBody struct {
	OrderID uint `json:"id-pesanan" binding:"required"`
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
				"success": false,
				"message": "Failed to create order. Please try again later",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Order created successfully",
			"data": gin.H{
				"id": newOrder.ID,
			},
		})
	}
}

func PayOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		var pay PayBody
		if err := bindBodyOrError(c, &pay); err != nil {
			return
		}

		var order entity.Order
		if err := db.GetConnection().Preload("Payment").Where("id = ?", pay.OrderID).First(&order).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Failed to fetch the order. Please try again later",
			})
			return
		}

		order.Status = ORDER_COMPLETED
		order.PaymentStatus = PAYMENT_PAID
		order.Payment.Status = true

		if err := db.GetConnection().Save(&order).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Failed to update order status. Please try again later",
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
			return
		}
		
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Your order has been canceled",
		})
	}
}
