package main

import (
	"receipt-processor-module/api"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	api.ImportRoutes(router)
	router.Run("localhost:4000")
}