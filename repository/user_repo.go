package repository

import "spos/users/models"

func (r *repo) FindUserById(userId int) (*models.User, error) {
	var (
		user models.User
		err  error
	)

	query := `select
	id, "name", email, phone_number, password, 
	email_verification_at, phone_verification_at
	from users
	where id=$1`

	err = r.db.Get(&user, query, userId)
	return &user, err
}
