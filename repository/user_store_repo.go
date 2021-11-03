package repository

import (
	"fmt"
	"spos/users/models"
)

var queryUserStore = `select
id, user_id, role_id, store_id, enabled,
created_at, updated_at
from user_store where enabled = true`

func (r *repo) FindUserStoreById(id int) (*models.UserStore, error) {
	var (
		us  models.UserStore
		err error
	)
	query := fmt.Sprintf("%s and id = $1", queryUserStore)

	err = r.db.Get(&us, query, id)
	return &us, err
}

func (r *repo) FindUserStoreByUserId(userId int) ([]*models.UserStore, error) {
	var (
		us  = make([]*models.UserStore, 0)
		err error
	)
	query := fmt.Sprintf("%s and user_id = $1", queryUserStore)

	err = r.db.Select(&us, query, userId)
	return us, err
}
