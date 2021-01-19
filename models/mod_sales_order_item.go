package models

// SalesOrderItem sales order items.
type SalesOrderItem struct {
	SalesOrderID uint       `json:"salesOrderID"`
	SalesOrder   SalesOrder `json:"salesOrder" gorm:"foreignKey:SalesOrderID"`
	ProductID    uint       `json:"productId"`
	Product      Product    `json:"product" gorm:"foreignKey:ProductID"`
	Quantity     int        `json:"quantity"`
}
