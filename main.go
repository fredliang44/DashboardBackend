package main

import (
	"DashboardBackend/handlers/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.POST("/login", auth.GenFunc().LoginHandler)
	router.GET("/healthcheck", checkAliveHandler)

	console := router.Group("/console")
	console.Use(auth.GenFunc().MiddlewareFunc())
	{
		//在此添加控制台路由, 例如console.GET("/healthcheck", checkAliveHandler)
		console.GET("/refresh_token", auth.GenFunc().RefreshHandler)
	}

	router.Run("localhost:8000")
}

func checkAliveHandler(c *gin.Context) {
	c.String(http.StatusOK, "Backend is alive")
}