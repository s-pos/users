package controllers

import (
	"spos/users/usecases"

	"github.com/labstack/echo"
)

type controller struct {
	usecase usecases.Usecase
}

type Controller interface {
	// GetProfileHandler handler for get data profile from authorization
	GetProfileHandler(e echo.Context) error
}

func NewController(uc usecases.Usecase) Controller {
	return &controller{
		usecase: uc,
	}
}
