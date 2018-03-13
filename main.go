package main

import (
	"DashboardBackend/handlers/auth"
	"DashboardBackend/handlers/user"
	"DashboardBackend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/v1/user/login", auth.AuthMiddleware.LoginHandler)
	router.POST("/v1/user/reg", user.RegHandler)
	router.GET("/v1/healthcheck", checkAliveHandler)

	console := router.Group("/v1")
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

	router.Run(utils.AppConfig.Host + ":" + utils.AppConfig.Port)
}

func checkAliveHandler(c *gin.Context) {
	c.String(http.StatusOK, "Backend is alive")
}
