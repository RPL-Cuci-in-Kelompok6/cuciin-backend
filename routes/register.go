package routes

import (
	"errors"
	"net/http"

	"github.com/RPL-Cuci-in-Kelompok6/cuciin-backend/db"
	"github.com/RPL-Cuci-in-Kelompok6/cuciin-backend/entity"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RegisterBody struct {
	Name        string `json:"nama" binding:"required"`
	Email       string `json:"email" binding:"required"`
	Password    string `json:"password" binding:"required"`
	PhoneNumber string `json:"telpon" binding:"required"`
}

func RegisterUser(createCustomer bool) func(*gin.Context) {
	registerUser := registerCustomer
	userRole := ROLE_CUSTOMER
	if !createCustomer {
		registerUser = registerPartner
		userRole = ROLE_PARTNER
	}

	return func(c *gin.Context) {
		var body RegisterBody
		if err := bindBodyOrError(c, &body); err != nil {
			return
		}

		id, result := registerUser(body)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
				c.JSON(http.StatusBadRequest, gin.H{
					"success":   false,
					"duplicate": true,
					"message":   "Email or phone already exist.",
				})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Failed to register user. Please try again later.",
			})
			return
		}

		token, err := generateJWT(id, userRole)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success":    false,
				"registered": true,
				"message":    "User registered but failed to generate token. Try to login.",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success":    true,
			"registered": true,
			"message":    "User registered successfully",
			"token":      token,
			"id":         id,
		})
	}
}

func registerCustomer(user RegisterBody) (uint, *gorm.DB) {
	customer := entity.Customer{
		Name:        user.Name,
		Email:       user.Email,
		Password:    user.Password,
		PhoneNumber: user.PhoneNumber,
	}

	result := db.GetConnection().Create(&customer)
	return customer.ID, result
}

func registerPartner(user RegisterBody) (uint, *gorm.DB) {
	partner := entity.Partner{
		Name:        user.Name,
		Email:       user.Email,
		Password:    user.Password,
		PhoneNumber: user.PhoneNumber,
	}

	result := db.GetConnection().Create(&partner)
	return partner.ID, result
}
