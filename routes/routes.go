package routes

import "github.com/labstack/echo/v4"

func RouteInit(e *echo.Group) {
	AuthRoutes(e)
	CategoryRoutes(e)
	OrderRoutes(e)
	ProductRoutes(e)
	UserDataRoutes(e)
}
