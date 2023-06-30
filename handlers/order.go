package handlers

import (
	"fmt"
	orderdto "indochat/dto/order"
	dto "indochat/dto/result"
	"indochat/models"
	"indochat/repositories"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type handlerOrders struct {
	OrdersRepository  repositories.OrdersRepository
	ProductRepository repositories.ProductsRepository
}

func HandlerOrders(OrdersRepository repositories.OrdersRepository, ProductRepository repositories.ProductsRepository) *handlerOrders {
	return &handlerOrders{OrdersRepository, ProductRepository}
}

// for view all data
func (h *handlerOrders) FindOrders(c echo.Context) error {
	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(float64)

	transactions, err := h.OrdersRepository.FindOrders(int(userId))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: transactions})
}

func (h *handlerOrders) GetOrderByUSer(c echo.Context) error {
	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(float64)

	transaction, err := h.OrdersRepository.GetOrderByUSer(int(userId))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	// transaction.Attachment = path_file + transaction.Attachment
	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: transaction})
}

func (h *handlerOrders) CreateOrders(c echo.Context) error {

	request := new(orderdto.CreateOrderRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(float64)
	productid, _ := strconv.Atoi(c.FormValue("product_id"))
	fmt.Println("Product ID:", productid) // Add this line for debugging

	localTime := time.Now()
	formattedTime := localTime.Format("Monday, 02-Jan-06 15:04:05 MST")

	// request := orderdto.CreateOrderRequest{
	// 	Status:       c.FormValue("status"),
	// 	ProductID:    productid,
	// 	CustomerID:   int(userId),
	// 	DiscountCode: c.FormValue("discount_code"), // Add this line to get the discount code from the request

	// }

	// request.Status = "pending"
	// request.CustomerID = int(userId)

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	product, err := h.ProductRepository.GetProduct(request.ProductID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}
	discountCode := c.FormValue("discount_code")
	if discountCode != "" {
		applyDiscount(discountCode, &product)
	}

	orders := models.Order{
		Date:         formattedTime,
		Status:       request.Status,
		CustomerID:   int(userId),
		ProductID:    request.ProductID,
		DiscountCode: request.DiscountCode, // Assign the discount code to the Order model

	}

	data, err := h.OrdersRepository.CreateOrders(orders)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: data})
}

func applyDiscount(discountCode string, product *models.Product) {
	switch discountCode {
	case "IC003":
		applyPercentageDiscount(product, 10)
	case "IC042":
		applyCategoryDiscount(product, "electronic", 5)
	case "IC015":
		if isWeekend() {
			applyPercentageDiscount(product, 10)
		}
	}
}

func applyPercentageDiscount(product *models.Product, discountPercent float64) {
	discount := product.Price * discountPercent / 100
	product.Price -= discount
}

func applyCategoryDiscount(product *models.Product, category string, discountPercent float64) {
	for _, productCategory := range product.Category {
		if productCategory.Name == category {
			applyPercentageDiscount(product, discountPercent)
			break
		}
	}
}

func isWeekend() bool {
	today := time.Now().Weekday()
	return today == time.Saturday || today == time.Sunday
}
