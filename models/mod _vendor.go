package models

// Vendor 供应商
type Vendor struct {
	Model
	Name        string `json:"name"`
	Address     string `json:"address"`
	PhoneNumber int    `json:"phoneNumber"`
}
