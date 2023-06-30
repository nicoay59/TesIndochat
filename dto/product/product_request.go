package productdto

type CreateProductRequest struct {
	Name       string `json:"name" form:"name" validate:"required"`
	Price      int    `json:"price" form:"price" validate:"required"`
	Desc       string `json:"desc" form:"desc" validate:"required"`
	Image      string `json:"image" form:"image"`
	CategoryID []int  `json:"category_id" form:"category_id" validate:"required"`
}
