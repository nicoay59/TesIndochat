package routes

import (
	"indochat/handlers"
	"indochat/pkg/middleware"
	"indochat/pkg/mysql"
	"indochat/repositories"

	"github.com/labstack/echo/v4"
)

func ProductRoutes(e *echo.Group) {
	productRepository := repositories.RepositoryProduct(mysql.DB)
	h := handlers.HandlerProduct(productRepository)

	e.GET("/products", middleware.Auth(h.FindProducts))
	e.GET("/product/:id", middleware.Auth(h.GetProduct))
	e.POST("/product", middleware.Auth(middleware.UploadFile(h.CreateProduct)))
	// e.DELETE("/product/:id", middleware.Auth(h.DeleteProduct))
	// e.PATCH("/product/:id", middleware.Auth(middleware.UploadFile(h.UpdateProduct)))
}
