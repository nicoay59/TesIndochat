package routes

import (
	"indochat/handlers"
	"indochat/pkg/mysql"
	"indochat/repositories"

	"github.com/labstack/echo/v4"
)

func UserDataRoutes(e *echo.Group) {
	userRepository := repositories.RepositoryUser(mysql.DB)
	h := handlers.HandlerUserData(userRepository)

	e.GET("/users", h.FindUsers)
	e.GET("/user/:id", h.GetUser)
	e.POST("/user", h.CreateUser)
	// e.PATCH("/user/:id", h.UpdateUser)
	e.DELETE("/user/:id", h.DeleteUser)
}
