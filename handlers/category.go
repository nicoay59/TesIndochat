package handlers

import (
	categorydto "indochat/dto/category"
	dto "indochat/dto/result"
	"indochat/models"
	"indochat/repositories"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type CategoryHandler struct {
	CategoryRepository repositories.CategoryRepository
}

func HandlerCategory(CategoryRepository repositories.CategoryRepository) *CategoryHandler {
	return &CategoryHandler{CategoryRepository}
}

func (h *CategoryHandler) FindCategory(c echo.Context) error {
	Category, err := h.CategoryRepository.FindCategories()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: Category})
}

func (h *CategoryHandler) GetCategory(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	Category, err := h.CategoryRepository.GetCategory(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: Category})
}

func (h *CategoryHandler) CreateCategory(c echo.Context) error {
	request := new(categorydto.CreateCategoryRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	category := models.Category{
		Name: request.Name,
	}

	data, err := h.CategoryRepository.CreateCategory(category)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: data})
}

func (h *CategoryHandler) DeleteCategory(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	Category, err := h.CategoryRepository.GetCategory(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.CategoryRepository.DeleteCategory(Category, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: data})
}
