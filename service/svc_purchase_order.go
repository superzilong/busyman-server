package service

import "gg/models"

// CreatePurchaseOrder create purchase order.
func CreatePurchaseOrder(purchaseOrder *models.PurchaseOrder) error {
	return db.Create(purchaseOrder).Error
}

// GetPurchaseOrderList get all purchase order list.
func GetPurchaseOrderList(purchaseOrderList *[]models.PurchaseOrder) error {
	return db.Find(purchaseOrderList).Error
}

// CreatePurchaseOrderItems create purchase order items in bulk.
func CreatePurchaseOrderItems(purchaseOrderItems *[]models.PurchaseOrderItem) error {
	return db.Create(purchaseOrderItems).Error
}

// GetPurchaseOrderItemsByOrderID get purchase order items by purchase order ID.
func GetPurchaseOrderItemsByOrderID(purchaseOrderID uint, items *[]models.PurchaseOrderItem) error {
	return db.Where("purchase_order_id = ?", purchaseOrderID).Preload("Product").Find(items).Error
}
