package routes

import (
	"chat_api/controllers"

	"github.com/labstack/echo/v4"
)

type Route struct {
	Controller controllers.Controller
}

func NewRoute(Controller controllers.Controller) Route {
	return Route{Controller: Controller}
}

func (u Route) UserRoutes(e *echo.Echo) {
	e.GET("/users", u.Controller.GetAll)
	e.GET("/users/:id", u.Controller.GetById)
	e.POST("/users", u.Controller.Create)
	e.PUT("/users/:id", u.Controller.Update)
	e.DELETE("/users/:id", u.Controller.Delete)
}
