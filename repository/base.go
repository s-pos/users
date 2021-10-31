package repository

import (
	"spos/users/models"

	"github.com/jmoiron/sqlx"
)

type repo struct {
	db *sqlx.DB
}

type Repository interface {
	// FindUserById query for get one data user
	// from views 'users_view'
	FindUserById(userId int) (*models.User, error)
}

func NewRepository(db *sqlx.DB) Repository {
	return &repo{
		db: db,
	}
}
