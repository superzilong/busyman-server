package service

import "gg/models"

// CreateProduct insert new product to DB.
func CreateProduct(product *models.Product) (err error) {
	return db.Create(product).Error
}

// GetAllProducts return all products.
func GetAllProducts(products *[]models.Product) (err error) {
	result := db.Find(products)
	return result.Error
}
