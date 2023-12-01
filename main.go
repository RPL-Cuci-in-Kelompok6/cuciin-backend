package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/RPL-Cuci-in-Kelompok6/cuciin-backend/db"
	"github.com/RPL-Cuci-in-Kelompok6/cuciin-backend/middleware"
	"github.com/RPL-Cuci-in-Kelompok6/cuciin-backend/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	db.Init()
	// db.MigrateDummy()

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "%s", "Pong")
	})
	r.GET("/secret", middleware.Authenticate, func(c *gin.Context) {
		c.String(http.StatusOK, "%s", "You are authorized")
	})
	r.GET("/secret/customer", middleware.Authenticate, middleware.Authorize(routes.ROLE_CUSTOMER), func(c *gin.Context) {
		c.String(http.StatusOK, "%s", "You are authorized as customer")
	})
	r.GET("/secret/mitra", middleware.Authenticate, middleware.Authorize(routes.ROLE_PARTNER), func(c *gin.Context) {
		c.String(http.StatusOK, "%s", "You are authorized as mitra")
	})

	// User Registration & Login
	r.POST("/customer/register", routes.RegisterUser(true))
	r.POST("/mitra/register", routes.RegisterUser(false))
	r.POST("/customer/login", routes.LoginUser(true))
	r.POST("/mitra/login", routes.LoginUser(false))

	// Customers
	r.GET("/customer/mitra", routes.GetPartners())
	r.POST("/customer/mitra/services", routes.GetServicesByPartner())
	r.POST("/customer/order", routes.GetOrdersByCustomer())
	r.POST("/customer/order/create", routes.CreateOrder())
	r.POST("/customer/order/cancel", routes.CancelOrder())
	r.POST("/customer/order/pay", routes.PayOrder())

	listenAddress := fmt.Sprintf("%s:%s", os.Getenv("ADDRESS"), os.Getenv("PORT"))
	r.Run(listenAddress)
}
