package middleware

import (
	"gg/pkg/jwt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// JWTAuthMiddleware is the middleware of parse jwt token.
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 2101,
				"msg":  "请求头中Authorization为空",
			})
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 2102,
				"msg":  "请求头中auth格式有误",
			})
			c.Abort()
			return
		}

		mc, err := jwt.ParseAccessToken(parts[1])
		if err != nil {
			if jwt.IsTokenExpiredErr(err) {
				if c.FullPath() != "/refreshToken" {
					c.JSON(452, gin.H{
						"code": 2103,
						"msg":  "accessToken expires.",
					})
					c.Abort()
					return
				}
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{
					"code": 2104,
					"msg":  "Invalid accessToken.",
				})
				c.Abort()
				return
			}

		}

		c.Set("username", mc.Username)
		c.Next()
	}
}
