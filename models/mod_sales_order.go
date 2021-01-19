package models

// SalesOrder 销售单
type SalesOrder struct {
	Model
	CustomerID   uint     `json:"customerId"`
	Customer     Customer `json:"cumtomer" gorm:"foreignKey:CustomerID"`
	CreateUserID uint     `json:"createUserId"`
	User         User     `json:"createUser" gorm:"foreignKey:CreateUserID"`
}
