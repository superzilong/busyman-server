package models

// PurchaseOrder 采购单
type PurchaseOrder struct {
	Model
	VendorID     uint   `json:"VendorID"`
	Vendor       Vendor `json:"vendor" gorm:"foreignKey:VendorID"`
	CreateUserID uint   `json:"createUserID"`
	User         User   `json:"createUser" gorm:"foreignKey:CreateUserID"`
}
