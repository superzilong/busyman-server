package service

import "gg/models"

// CreateProduct insert new product to DB.
func CreateProduct(product *models.Product) error {
	return db.Create(product).Error
}

// GetAllProducts return all products.
func GetAllProducts(products *[]models.Product) error {
	return db.Find(products).Error
}
