package routes

import (
	"net/http"

	"github.com/RPL-Cuci-in-Kelompok6/cuciin-backend/db"
	"github.com/RPL-Cuci-in-Kelompok6/cuciin-backend/entity"
	"github.com/gin-gonic/gin"
)

type BaseProfileResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"nama"`
	PhoneNumber string `json:"telepon"`
	Email       string `json:"email"`
}

type RequestBody struct {
	Email string `json:"email" binding:"required"`
}

func GetProfile(role int) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request RequestBody
		if err := bindBodyOrError(c, &request); err != nil {
			return
		}

		var response any

		switch role {
		case ROLE_CUSTOMER:
			{
				var query entity.Customer
				if err := db.GetConnection().Where(&entity.Customer{Email: request.Email}).Select("id", "name", "phone_number", "email").First(&query).Error; err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{
						"success": false,
						"message": "Failed to get profile. Please try again later",
					})
					return
				}
				response = BaseProfileResponse{
					ID:          query.ID,
					Name:        query.Name,
					PhoneNumber: query.PhoneNumber,
					Email:       query.Email,
				}
			}
		case ROLE_PARTNER:
			{
				var query entity.Partner
				if err := db.GetConnection().Where(&entity.Partner{Email: request.Email}).Select("id", "name", "phone_number", "email", "map_link").First(&query).Error; err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{
						"success": false,
						"message": "Failed to get profile. Please try again later",
					})
					return
				}
				var resp struct {
					BaseProfileResponse
					MapLink string
				}
				resp.ID = query.ID
				resp.Email = query.Email
				resp.Name = query.Name
				resp.PhoneNumber = query.PhoneNumber

				response = resp
			}
		}
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"profile": response,
		})
	}
}
