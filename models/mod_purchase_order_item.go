package models

// PurchaseOrderItem 采购单产品项目
type PurchaseOrderItem struct {
	PurchaseOrderID uint          `json:"purchaseOrderID"`
	PurchaseOrder   PurchaseOrder `json:"-" gorm:"foreignKey:PurchaseOrderID"`
	ProductID       uint          `json:"productID"`
	Product         Product       `json:"product" gorm:"foreignKey:ProductID"`
	Quantity        int           `json:"quantity"`
}
