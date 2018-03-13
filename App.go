package main

import (
	"DashboardBackend/handlers/user"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.POST("/user", user.Handler)
	}

	router.Run(":8080")
}
