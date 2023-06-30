package routes

import (
	"indochat/handlers"
	"indochat/pkg/mysql"
	"indochat/repositories"

	"github.com/labstack/echo/v4"
)

func CategoryRoutes(e *echo.Group) {
	categoryRepository := repositories.RepositoryCategory(mysql.DB)
	h := handlers.HandlerCategory(categoryRepository)

	e.GET("/categories", h.FindCategory)
	e.GET("/category/:id", h.GetCategory)
	e.POST("/category", h.CreateCategory)
	e.DELETE("/category/:id", h.DeleteCategory)
	// e.PATCH("/category/:id", h.upda)
}
