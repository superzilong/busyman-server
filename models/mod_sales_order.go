package models

// SalesOrder 销售单
type SalesOrder struct {
	Model
	CustomerID   uint     `json:"customerID"`
	Customer     Customer `json:"cumtomer" gorm:"foreignKey:CustomerID"`
	CreateUserID uint     `json:"createUserID"`
	User         User     `json:"createUser" gorm:"foreignKey:CreateUserID"`
}
