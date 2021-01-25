package service

import "gg/models"

// CreateVendor create vendor.
func CreateVendor(vendor *models.Vendor) error {
	return db.Create(vendor).Error
}

// GetAllVendors return all vendors.
func GetAllVendors(vendors *[]models.Vendor) error {
	result := db.Find(vendors)
	return result.Error
}
