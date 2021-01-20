package models

// PurchaseOrderItem 采购单产品项目
type PurchaseOrderItem struct {
	PurchaseOrderID uint          `json:"purchaseOrderId"`
	PurchaseOrder   PurchaseOrder `json:"purchaseOrder" gorm:"foreignKey:PurchaseOrderID"`
	ProductID       uint          `json:"productId"`
	Product         Product       `json:"product" gorm:"foreignKey:ProductID"`
	Quantity        int           `json:"quantity"`
}
