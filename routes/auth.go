package routes

import (
	"indochat/handlers"
	// "indochat/pkg/middleware"
	"indochat/pkg/mysql"
	"indochat/repositories"

	"github.com/labstack/echo/v4"
)

func AuthRoutes(e *echo.Group) {
	authRepository := repositories.RepositoryAuth(mysql.DB)
	h := handlers.HandlerAuth(authRepository)

	e.POST("/register", h.Register)
	e.POST("/login", h.Login)         // add this code
	e.GET("/check-auth", h.CheckAuth) // add this code
	e.GET("/test", h.Test)            // add this code
}
