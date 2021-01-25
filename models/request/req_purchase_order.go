package request

// OrderItem one item in the order.
type OrderItem struct {
	ProductID uint `json:"productID"`
	Quantity  int  `json:"quantity"`
}

// PurchaseOrder purchase order in the request body.
type PurchaseOrder struct {
	VendorID   uint        `json:"vendorID"`
	OrderItems []OrderItem `json:"orderItems"`
}
