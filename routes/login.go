package routes

import (
	"errors"
	"net/http"

	"github.com/RPL-Cuci-in-Kelompok6/cuciin-backend/db"
	"github.com/RPL-Cuci-in-Kelompok6/cuciin-backend/entity"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LoginBody struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func LoginUser(createCustomer bool) func(*gin.Context) {
	authenticateUser := authenticateCustomer
	userRole := ROLE_CUSTOMER
	if !createCustomer {
		authenticateUser = authenticatePartner
		userRole = ROLE_PARTNER
	}

	return func(c *gin.Context) {
		var body LoginBody
		if err := bindBodyOrError(c, &body); err != nil {
			return
		}

		id, ok, err := authenticateUser(body.Email, body.Password)
		if errors.Is(err, gorm.ErrRecordNotFound) || !ok {
			c.JSON(http.StatusConflict, gin.H{
				"success": false,
				"message": "Email not registered or password missmatch",
			})
			return
		} else if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Failed to login. Please try again later",
			})
			return
		}

		token, err := generateJWT(id, userRole)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Failed to generate token. Please try again",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Login success.",
			"token":   token,
			"id":      id,
		})
	}
}

func authenticateCustomer(email, password string) (uint, bool, error) {
	user := entity.Customer{}
	if err := db.GetConnection().Where("email = ?", email).First(&user).Error; err != nil {
		return 0, false, err
	}
	return user.ID, validatePassword(password, user.Password), nil
}

func authenticatePartner(email, password string) (uint, bool, error) {
	user := entity.Partner{}
	if err := db.GetConnection().Where("email = ?", email).First(&user).Error; err != nil {
		return 0, false, err
	}
	return user.ID, validatePassword(password, user.Password), nil
}
