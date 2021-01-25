package controller

import (
	"gg/models"
	"gg/models/request"
	"gg/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreatePurchaseOrder is the callback of /purchaseOrders post route.
func CreatePurchaseOrder(c *gin.Context) {
	var inputPurchaseOrder request.PurchaseOrder
	err := c.ShouldBind(&inputPurchaseOrder)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 2001,
			"msg":  "无效的参数",
		})
		return
	}
	modPurchaseOrder := models.PurchaseOrder{
		VendorID:     inputPurchaseOrder.VendorID,
		CreateUserID: 1,
	}
	err = service.CreatePurchaseOrder(&modPurchaseOrder)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 3000,
			"msg":  err.Error(),
		})
		return
	}

	orderItems := make([]models.PurchaseOrderItem, 0, len(inputPurchaseOrder.OrderItems))
	for _, value := range inputPurchaseOrder.OrderItems {
		orderItems = append(orderItems, models.PurchaseOrderItem{
			PurchaseOrderID: modPurchaseOrder.ID,
			ProductID:       value.ProductID,
			Quantity:        value.Quantity,
		})
	}

	err = service.CreatePurchaseOrderItems(&orderItems)
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

// GetAllPurChaseOrders is the callback of /vendors Get route.
func GetAllPurChaseOrders(c *gin.Context) {
	var purchaseOrders []models.PurchaseOrder
	err := service.GetPurchaseOrderList(&purchaseOrders)
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
		"data": purchaseOrders,
	})
}

// GetPurchaseOrderItemsByOrderID return order items by specific purchase order ID.
func GetPurchaseOrderItemsByOrderID(c *gin.Context) {
	params := c.Request.URL.Query()
	strPurchaseOrderID := params.Get("purchaseOrderID")
	purchaseOrderID, err := strconv.ParseUint(strPurchaseOrderID, 10, 32)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 3000,
			"msg":  err.Error(),
		})
		return
	}
	purchaseOrderItems := make([]models.PurchaseOrderItem, 0)
	err = service.GetPurchaseOrderItemsByOrderID(uint(purchaseOrderID), &purchaseOrderItems)
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
		"data": purchaseOrderItems,
	})
}
