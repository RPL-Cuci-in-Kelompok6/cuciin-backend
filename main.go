package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/RPL-Cuci-in-Kelompok6/cuciin-backend/db"
	"github.com/RPL-Cuci-in-Kelompok6/cuciin-backend/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	db.Init()

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "%s", "Pong")
	})

	r.POST("/customer/register", routes.RegisterUser(true))
	r.POST("/mitra/register", routes.RegisterUser(false))

	listenAddress := fmt.Sprintf("%s:%s", os.Getenv("ADDRESS"), os.Getenv("PORT"))
	r.Run(listenAddress)
}
