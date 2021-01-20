package controller

import (
	"gg/models"
	"gg/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateProduct is the callback of /products post route.
func CreateProduct(c *gin.Context) {
	var product models.Product
	err := c.ShouldBind(&product)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 2001,
			"msg":  "无效的参数",
		})
		return
	}

	err = service.CreateProduct(&product)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 3000,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 2000,
		"msg":  "success",
	})
}

// GetAllProducts is the callback of /products Get route.
func GetAllProducts(c *gin.Context) {
	var products []models.Product
	err := service.GetAllProducts(&products)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 3000,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 2000,
		"msg":  "success",
		"data": products,
	})
}
