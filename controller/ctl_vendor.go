package controller

import (
	"gg/models"
	"gg/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateVendor is the callback of /vendors post route.
func CreateVendor(c *gin.Context) {
	var vendor models.Vendor
	err := c.ShouldBind(&vendor)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 2001,
			"msg":  "无效的参数",
		})
		return
	}

	err = service.CreateVendor(&vendor)
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

// GetAllVendors is the callback of /vendors Get route.
func GetAllVendors(c *gin.Context) {
	var vendors []models.Vendor
	err := service.GetAllVendors(&vendors)
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
		"data": vendors,
	})
}
