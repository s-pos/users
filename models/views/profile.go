package views

import "spos/users/models"

type ResponseProfile struct {
	ID              int          `json:"id"`
	Name            string       `json:"name"`
	Email           string       `json:"email"`
	Phone           string       `json:"phone"`
	IsEmailVerified bool         `json:"isEmailVerified"`
	Stores          []*UserStore `json:"stores"`
}

type UserStore struct {
	ID      int           `json:"id"`
	Enabled bool          `json:"enabled"`
	Store   *models.Store `json:"store"`
}
