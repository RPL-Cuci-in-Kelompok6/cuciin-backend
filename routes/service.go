package routes

import (
	"net/http"
	"strconv"

	"github.com/RPL-Cuci-in-Kelompok6/cuciin-backend/db"
	"github.com/RPL-Cuci-in-Kelompok6/cuciin-backend/entity"
	"github.com/gin-gonic/gin"
)

type createServiceBody struct {
	Name      string `json:"nama" binding:"required"`
	Price     int    `json:"harga" binding:"required"`
	PartnerID int    `json:"id-mitra" binding:"required"`
}

func CreateService(c *gin.Context) {
	var body createServiceBody
	if err := bindBodyOrError(c, &body); err != nil {
		return
	}

	service := entity.Service{
		Name:      body.Name,
		Price:     uint64(body.Price),
		PartnerID: uint(body.PartnerID),
	}

	result := db.GetConnection().Create(&service)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to register service. Please try again later.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Service created successfully",
		"data": gin.H{
			"id": service.ID,
		},
	})
}

func DeleteService(c *gin.Context) {
	service_id := c.Param("service_id")

	id, err := strconv.Atoi(service_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid service_id. Must be positive number",
		})
		return
	}

	result := db.GetConnection().Delete(&entity.WashingMachine{}, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to delete service. Please try again later.",
			"data": gin.H{
				"id": id,
			},
		})
		return
	}

	c.JSON(http.StatusInternalServerError, gin.H{
		"success": true,
		"message": "Service deleted",
		"data": gin.H{
			"id": id,
		},
	})
}
