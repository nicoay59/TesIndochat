package handlers

import (
	discountdto "indochat/dto/discount"
	dto "indochat/dto/result"
	"indochat/models"
	"indochat/repositories"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type DiscountHandler struct {
	DiscountRepository repositories.DiscountRepository
}

func HandlerDiscount(DiscountRepository repositories.DiscountRepository) *DiscountHandler {
	return &DiscountHandler{DiscountRepository}
}

func (h *DiscountHandler) FindDiscount(c echo.Context) error {
	Discount, err := h.DiscountRepository.FindDiscounts()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: Discount})
}

func (h *DiscountHandler) GetDiscount(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	Discount, err := h.DiscountRepository.GetDiscount(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: Discount})
}

func (h *DiscountHandler) CreateDiscount(c echo.Context) error {
	request := new(discountdto.CreateDiscountRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	Discount := models.Discount{
		Code:           request.Code,
		DiscountAmount: request.DiscountAmount,
	}

	data, err := h.DiscountRepository.CreateDiscount(Discount)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: data})
}

func (h *DiscountHandler) DeleteDiscount(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	Discount, err := h.DiscountRepository.GetDiscount(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.DiscountRepository.DeleteDiscount(Discount, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: data})
}
