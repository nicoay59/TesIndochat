package orderdto

type CreateOrderRequest struct {
	Status       string `json:"status"`
	DiscountCode string `json:"discount_code"`
	CustomerID   int    `json:"customer_id"`
	ProductID    int    `json:"product_id"`
}
