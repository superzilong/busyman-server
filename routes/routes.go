package routes

import (
	"gg/controller"
	"gg/middleware"
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

	// r.GET("/user", middleware.JWTAuthMiddleware(), func(c *gin.Context) {
	// 	u := service.GetUser()
	// 	c.JSON(http.StatusOK, u)
	// })
	r.POST("/refreshToken", middleware.JWTAuthMiddleware(), controller.RefreshToken)
	r.POST("/login", controller.Login)

	r.POST("/products", middleware.JWTAuthMiddleware(), controller.CreateProduct)
	r.GET("/products", middleware.JWTAuthMiddleware(), controller.GetAllProducts)

	r.POST("/customers", middleware.JWTAuthMiddleware())
	r.GET("/customers", middleware.JWTAuthMiddleware())

	r.POST("/vendors", middleware.JWTAuthMiddleware(), controller.CreateVendor)
	r.GET("/vendors", middleware.JWTAuthMiddleware(), controller.GetAllVendors)

	r.POST("/salesOrders", middleware.JWTAuthMiddleware())
	r.GET("/salesOrders", middleware.JWTAuthMiddleware())

	r.POST("/purchaseOrders", middleware.JWTAuthMiddleware(), controller.CreatePurchaseOrder)
	r.GET("/purchaseOrders", middleware.JWTAuthMiddleware(), controller.GetAllPurChaseOrders)

	r.GET("/purchaseOrderItems", middleware.JWTAuthMiddleware(), controller.GetPurchaseOrderItemsByOrderID)
	return r
}
