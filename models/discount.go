package models

type Discount struct {
	ID             int     `json:"id"`
	Code           string  `json:"code" gorm:"varchar(255)"`
	DiscountAmount float64 `json:"discount_amount"`
}
