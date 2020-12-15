package routes

import (
	"gg/controller"
	"gg/dao"
	"gg/logger"
	"gg/middleware"
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
	r.POST("/refreshToken", middleware.JWTAuthMiddleware(), controller.RefreshToken)

	r.POST("/login", controller.Login)
	return r
}
