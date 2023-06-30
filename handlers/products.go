package handlers

import (
	"encoding/json"
	productdto "indochat/dto/product"
	dto "indochat/dto/result"
	"indochat/models"
	"indochat/repositories"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type handlerProduct struct {
	ProductRepository repositories.ProductsRepository
}

func HandlerProduct(ProductRepository repositories.ProductsRepository) *handlerProduct {
	return &handlerProduct{ProductRepository}
}

// for view all data
func (h *handlerProduct) FindProducts(c echo.Context) error {
	product, err := h.ProductRepository.FindProducts()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: product})
}

func (h *handlerProduct) GetProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var product models.Product
	product, err := h.ProductRepository.GetProduct(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: product})
}

func (h *handlerProduct) CreateProduct(c echo.Context) error {
	var err error

	filename := c.Get("dataFile").(string)

	price, _ := strconv.Atoi(c.FormValue("price"))
	// orders_id, _ := strconv.Atoi(r.FormValue("orders_id"))

	categoryIdString := c.FormValue("category_id")
	if categoryIdString == "" {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: "Error: category_id form value is missing."})
	}

	var categoriesId []int
	err = json.Unmarshal([]byte(categoryIdString), &categoriesId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	if len(categoriesId) == 0 {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: "Error: category_id form value is missing."})
	}

	request := productdto.CreateProductRequest{
		Name:       c.FormValue("name"),
		Image:      c.FormValue("image"),
		Price:      price,
		CategoryID: categoriesId,
		Desc:       c.FormValue("desc"),
	}

	// validation
	validation := validator.New()
	err = validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	bill := int(request.Price)
	discount := int(10)
	afterDiscount := float64(bill - (bill * discount / 100))
	categories, _ := h.ProductRepository.FindCategoriesById(request.CategoryID)

	product := models.Product{
		Name:     request.Name,
		Desc:     request.Desc,
		Price:    afterDiscount,
		Category: categories,
		Image:    filename,
	}

	product, err = h.ProductRepository.CreateProduct(product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	product, _ = h.ProductRepository.GetProduct(product.ID)

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: product})
}
