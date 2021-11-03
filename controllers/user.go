package controllers

import (
	"github.com/labstack/echo"
	"github.com/s-pos/go-utils/middleware"
)

func (c *controller) GetProfileHandler(e echo.Context) error {
	var (
		req = e.Request()
		ctx = req.Context()
	)

	userId := ctx.Value(middleware.CTX_KEY_USER).(int)

	return c.usecase.Profile(ctx, userId).Write(e)
}
