package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/RPL-Cuci-in-Kelompok6/cuciin-backend/db"
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

	listenAddress := fmt.Sprintf("%s:%s", os.Getenv("ADDRESS"), os.Getenv("PORT"))
	r.Run(listenAddress)
}
