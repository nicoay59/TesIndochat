package categorydto

type CreateCategoryRequest struct {
	Name string `json:"name" form:"name" validate:"required"`
}