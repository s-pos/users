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
	// from table 'users'
	FindUserById(userId int) (*models.User, error)

	// FindStoreById query for get one data store
	FindStoreById(id int) (*models.Store, error)

	// FindUserStoreById query for get one data user_store (relation)
	// from table 'user_store'
	FindUserStoreById(id int) (*models.UserStore, error)

	// FindUserStoreByUserId will return array of userStore
	// beacause relation 1 to many
	FindUserStoreByUserId(userId int) ([]*models.UserStore, error)
}

func NewRepository(db *sqlx.DB) Repository {
	return &repo{
		db: db,
	}
}
