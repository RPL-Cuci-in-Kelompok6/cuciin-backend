package routes

import (
	"net/http"

	"github.com/RPL-Cuci-in-Kelompok6/cuciin-backend/db"
	"github.com/RPL-Cuci-in-Kelompok6/cuciin-backend/entity"
	"github.com/gin-gonic/gin"
)

type PartnerServiceBody struct {
	PartnerID uint `json:"partner_id"`
}

func GetPartners() gin.HandlerFunc {
	return func(c *gin.Context) {
		var partners []entity.Partner
		if err := db.GetConnection().Select("name", "phone_number", "email", "map_link").Limit(20).Find(&partners).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Failed to get partners, please try again later.",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success":  true,
			"partners": partners,
		})
	}
}

func GetServicesByPartner() gin.HandlerFunc {
	return func(c *gin.Context) {
		var serviceRequest PartnerServiceBody
		if err := bindBodyOrError(c, &serviceRequest); err != nil {
			return
		}

		var partner entity.Partner

		if err := db.GetConnection().Where(&entity.Service{PartnerID: serviceRequest.PartnerID}).Select("partner_id", "services").Find(&partner); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Failed to get services, please try again later",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success":  true,
			"services": partner.Services,
		})
	}
}
