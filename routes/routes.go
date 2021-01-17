package routes

import (
	"gg/controller"
	"gg/middleware"
	"gg/service"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Setup routes.
func Setup() *gin.Engine {
	r := gin.New()
	config := cors.DefaultConfig()
	config.AllowMethods = append(config.AllowMethods, "OPTIONS")
	config.AllowAllOrigins = true
	config.AddAllowHeaders("Authorization")
	r.Use(middleware.GinLogger(), middleware.GinRecovery(true), cors.New(config))

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("/user", middleware.JWTAuthMiddleware(), func(c *gin.Context) {
		u := service.GetUser()
		c.JSON(http.StatusOK, u)
	})
	r.POST("/refreshToken", middleware.JWTAuthMiddleware(), controller.RefreshToken)

	r.POST("/login", controller.Login)
	return r
}
