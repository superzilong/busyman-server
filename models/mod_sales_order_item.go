package models

// SalesOrderItem 销售单产品项目
type SalesOrderItem struct {
	SalesOrderID uint       `json:"salesOrderID"`
	SalesOrder   SalesOrder `json:"salesOrder" gorm:"foreignKey:SalesOrderID"`
	ProductID    uint       `json:"productID"`
	Product      Product    `json:"product" gorm:"foreignKey:ProductID"`
	Quantity     int        `json:"quantity"`
}
