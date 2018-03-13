package auth

import (
	"time"

	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
)

func checkPassword(userId string, password string, c *gin.Context) (string, bool) {
	//TODO 对接数据库用户表
	if (userId == "admin" && password == "admin") || (userId == "test" && password == "test") {
		return userId, true
	}

	return userId, false
}

func checkUserID(userId string, c *gin.Context) bool {
	//TODO 确定是否有管理员权限
	if userId == "admin" {
		return true
	}

	return false
}

func GenFunc() *jwt.GinJWTMiddleware {
	// the jwt middleware
	authMiddleware := &jwt.GinJWTMiddleware{
		Realm:         "test zone",
		Key:           []byte("secret key"),
		Timeout:       time.Hour,
		MaxRefresh:    time.Hour,
		Authenticator: checkPassword,
		Authorizator:  checkUserID,
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		TokenLookup: "header:Authorization",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	}

	return authMiddleware
}
