package main

import (
	"DashboardBackend/handlers/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/api/v1/user/login", auth.AuthMiddleware.LoginHandler)
	router.POST("/api/v1/user/reg", user.RegHandler)
	router.GET("/api/v1/healthcheck", checkAliveHandler)

	console := router.Group("/api/v1")
	console.Use(auth.AuthMiddleware.MiddlewareFunc())
	{
		userGroup := console.Group("/user")
		{
			userGroup.GET("/hello", func(c *gin.Context) {
				c.String(http.StatusOK, "hello")
			})
			userGroup.GET("/refresh_token", auth.AuthMiddleware.RefreshHandler)
		}
		// //在此添加控制台路由 console.GET("/healthcheck", checkAliveHandler)
	}

	router.Run("localhost:8000")
}

func checkAliveHandler(c *gin.Context) {
	c.String(http.StatusOK, "Backend is alive")
}
