package main

import (
	"github.com/gin-gonic/gin"
	"DashboardBackend/handlers/auth"
	"net/http"
)



func main() {
	router := gin.Default()

	router.POST("/login", auth.GenFunc().LoginHandler)

	console := router.Group("/console")
	console.Use(auth.GenFunc().MiddlewareFunc())
	{
		console.GET("/hello", checkAliveHandler)
		console.GET("/refresh_token", auth.GenFunc().RefreshHandler)
	}

	router.Run(":8000")
}


func checkAliveHandler(c *gin.Context)  {
	c.String(http.StatusOK, "Backend is alive")
}