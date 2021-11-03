package usecases

import (
	"context"
	"spos/users/repository"

	"github.com/s-pos/go-utils/utils/response"
)

type usecase struct {
	repo repository.Repository
}

type Usecase interface {
	// Profile for get user data from header/authorization
	Profile(ctx context.Context, userId int) response.Output
}

func NewUsecases(r repository.Repository) Usecase {
	return &usecase{
		repo: r,
	}
}
