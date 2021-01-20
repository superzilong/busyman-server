package models

// PurchaseOrder 采购单
type PurchaseOrder struct {
	Model
	VendorID     uint   `json:"VendorId"`
	Vendor       Vendor `json:"vendor" gorm:"foreignKey:VendorID"`
	CreateUserID uint   `json:"createUserId"`
	User         User   `json:"createUser" gorm:"foreignKey:CreateUserID"`
}
