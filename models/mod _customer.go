package models

// Customer 客户
type Customer struct {
	Model
	Name        string `json:"name"`
	Address     string `json:"address"`
	PhoneNumber int    `json:"phoneNumber"`
}
