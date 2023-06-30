package routes

import (
	"indochat/handlers"
	"indochat/pkg/middleware"
	"indochat/pkg/mysql"
	"indochat/repositories"

	"github.com/labstack/echo/v4"
)

func OrderRoutes(e *echo.Group) {
	OrdersRepository := repositories.RepositoryOrders(mysql.DB)
	ProductRepository := repositories.RepositoryProduct(mysql.DB) // Create an instance of the ProductsRepository

	h := handlers.HandlerOrders(OrdersRepository, ProductRepository)

	e.GET("/orders", middleware.Auth(h.FindOrders))
	e.POST("/order", middleware.Auth(h.CreateOrders))
	e.GET("/orderuser", middleware.Auth(h.GetOrderByUSer))
	// e.POST("/notification", h.Notification)
}
