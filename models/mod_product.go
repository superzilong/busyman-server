package models

// Product 产品
type Product struct {
	Model
	Name string `json:"name"`
	Unit string `json:"unit"`
}
