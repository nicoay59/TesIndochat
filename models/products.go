package models

type Product struct {
	ID         int        `json:"id" gorm:"primary_key:auto_increment"`
	Name       string     `json:"name" form:"name" gorm:"varchar(255)"`
	Price      float64    `json:"price" form:"price"`
	Desc       string     `json:"desc" form:"desc" gorm:"varchar(255)"`
	Image      string     `json:"image"`
	Category   []Category `json:"category" gorm:"many2many:product_categories;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CategoryID []int      `json:"category_id" form:"category_id" gorm:"-"`
}

type ProductResponse struct {
	ID         int        `json:"id"`
	Name       string     `json:"name"`
	Price      float64    `json:"price"`
	Desc       string     `json:"desc"`
	Image      string     `json:"image"`
	Category   []Category `json:"category" gorm:"many2many:product_categories"`
	CategoryID []int      `json:"-" form:"-" gorm:"-"`
}

func (ProductResponse) TableName() string {
	return "products"
}
