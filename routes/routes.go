package routes

import (
	"gg/dao"
	"gg/logger"
	"gg/middleware"
	"gg/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Setup routes.
func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("/user", middleware.JWTAuthMiddleware(), func(c *gin.Context) {
		u := dao.GetUser()
		c.JSON(http.StatusOK, u)
	})
	r.POST("/login", func(c *gin.Context) {
		var user models.UserInfo
		err := c.ShouldBind(&user)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 2001,
				"msg":  "无效的参数",
			})
			return
		}

		if dao.GetUserInfoByName(user.Name).Password != user.Password {
			c.JSON(http.StatusOK, gin.H{
				"code": 2002,
				"msg":  "用户名或密码不对",
			})
			return
		}

		tokenString, err := middleware.GenToken(user.Name)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 2006,
				"msg":  "生成token出错",
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"code": 2000,
			"msg":  "success",
			"data": gin.H{"token": tokenString},
		})

	})
	return r
}
