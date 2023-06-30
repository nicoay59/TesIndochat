package discountdto

type CreateDiscountRequest struct {
	Code           string  `json:"code" form:"code" validate:"required"`
	DiscountAmount float64 `json:"discount_amount"`
}
