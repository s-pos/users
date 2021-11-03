package repository

import "spos/users/models"

func (r *repo) FindStoreById(id int) (*models.Store, error) {
	var (
		store models.Store
		err   error
	)

	query := `select id, owner_id, ref_shop_id, ref_user_id, ref_shop_status,
	name, domain, logo, description, enabled, type, source, created_at
	from stores where id = $1 and enabled = true`

	err = r.db.Get(&store, query, id)
	return &store, err
}
