package routes

import (
	"net/http"
	"strconv"
	"time"

	"github.com/RPL-Cuci-in-Kelompok6/cuciin-backend/db"
	"github.com/RPL-Cuci-in-Kelompok6/cuciin-backend/entity"
	"github.com/gin-gonic/gin"
)

type createMachineBody struct {
	Brand     string `json:"merek" binding:"required"`
	ServiceID int    `json:"id-layanan" binding:"required"`
}

func CreateMachine(c *gin.Context) {
	var body createMachineBody
	if err := bindBodyOrError(c, &body); err != nil {
		return
	}

	machine := entity.WashingMachine{
		Brand:       body.Brand,
		ServiceID:   uint(body.ServiceID),
		AvailableAt: time.Now(),
	}

	result := db.GetConnection().Create(&machine)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to register machine. Please try again later.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Washing machine created successfully",
		"data": gin.H{
			"id": machine.ID,
		},
	})
}

func DeleteMachine(c *gin.Context) {
	machine_id := c.Param("machine_id")

	id, err := strconv.Atoi(machine_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid machine_id. Must be positive number",
		})
		return
	}

	result := db.GetConnection().Delete(&entity.WashingMachine{}, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to delete machine. Please try again later.",
			"data": gin.H{
				"id": id,
			},
		})
		return
	}

	c.JSON(http.StatusInternalServerError, gin.H{
		"success": true,
		"message": "Machine deleted",
		"data": gin.H{
			"id": id,
		},
	})
}
