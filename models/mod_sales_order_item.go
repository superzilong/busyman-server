package models

// SalesOrderItem 销售单产品项目
type SalesOrderItem struct {
	SalesOrderID uint       `json:"salesOrderId"`
	SalesOrder   SalesOrder `json:"salesOrder" gorm:"foreignKey:SalesOrderID"`
	ProductID    uint       `json:"productId"`
	Product      Product    `json:"product" gorm:"foreignKey:ProductID"`
	Quantity     int        `json:"quantity"`
}
