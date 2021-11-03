package routes

import (
	"spos/users/controllers"

	"github.com/labstack/echo"
	"github.com/s-pos/go-utils/middleware"
)

type route struct {
	middleware middleware.Clients
	controller controllers.Controller
}

type Route interface {
	Router() *echo.Echo
}

func NewRouter(mdl middleware.Clients, ctrl controllers.Controller) Route {
	return &route{
		middleware: mdl,
		controller: ctrl,
	}
}

func (r *route) Router() *echo.Echo {
	e := echo.New()

	user := e.Group("", echo.WrapMiddleware(r.middleware.Logger), echo.WrapMiddleware(r.middleware.Session))
	user.GET("/", r.controller.GetProfileHandler)

	return e
}
