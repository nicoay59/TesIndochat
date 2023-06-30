package handlers

import (
	customerdto "indochat/dto/customer"
	dto "indochat/dto/result"
	"indochat/models"
	"indochat/repositories"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type CustomerHandler struct {
	CustomerRepository repositories.UserRepository
}

func HandlerUserData(UserRepository repositories.UserRepository) *CustomerHandler {
	return &CustomerHandler{UserRepository}
}

func (h *CustomerHandler) FindUsers(c echo.Context) error {
	userData, err := h.CustomerRepository.FindUsers()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: userData})
}

// func (h *CustomerHandler) GetTrans(c echo.Context) error {
// 	userData, err := h.CustomerRepository.GetTrans()
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
// 	}
// 	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: userData})
// }

func (h *CustomerHandler) GetUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	userData, err := h.CustomerRepository.GetUser(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: userData})
}

func (h *CustomerHandler) CreateUser(c echo.Context) error {
	request := new(customerdto.CreateCustomerRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	user := models.Customer{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}

	data, err := h.CustomerRepository.CreateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponse(data)})
}

// func (h *CustomerHandler) UpdateUser(c echo.Context) error {
// 	request := new(customerdto.upda)
// 	if err := c.Bind(&request); err != nil {
// 		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
// 	}

// 	id, _ := strconv.Atoi(c.Param("id"))

// 	user, err := h.UserRepository.GetUser(id)

// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
// 	}

// 	if request.FullName != "" {
// 		user.FullName = request.FullName
// 	}

// 	if request.Email != "" {
// 		user.Email = request.Email
// 	}

// 	if request.Password != "" {
// 		user.Password = request.Password
// 	}
// 	if request.Address != "" {
// 		user.Address = request.Address
// 	}
// 	if request.Phone != "" {
// 		user.Phone = request.Phone
// 	}
// 	if request.Role != "" {
// 		user.Role = request.Role
// 	}

// 	data, err := h.UserRepository.UpdateUser(user)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
// 	}

// 	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponse(data)})
// }

func (h *CustomerHandler) DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := h.CustomerRepository.GetUser(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.CustomerRepository.DeleteUser(user, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponse(data)})
}

func convertResponse(u models.Customer) customerdto.CreateCustomerResponse {
	return customerdto.CreateCustomerResponse{
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
	}
}
