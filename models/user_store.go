package models

import (
	"encoding/json"
	"time"
)

type UserStore struct {
	ID        int     `db:"id"`
	UserID    int     `db:"user_id"`
	RoleID    int     `db:"role_id"`
	StoreID   int     `db:"store_id"`
	Enabled   bool    `db:"enabled"`
	CreatedAt string  `db:"created_at"`
	UpdatedAt *string `db:"updated_at"`
	Deleted   bool    `db:"deleted"`

	// relation table
	Store []byte `db:"store"`
	User  []byte `db:"user"`
	Role  []byte `db:"role"`
}

func NewUserStore() *UserStore {
	now := time.Now().UTC()

	return &UserStore{
		RoleID:    1,
		Enabled:   enabled,
		CreatedAt: now.Format(layout),
	}
}

func (us *UserStore) GetID() int {
	return us.ID
}

func (us *UserStore) SetUserID(userId int) {
	us.UserID = userId
}

func (us *UserStore) GetUserID() int {
	return us.UserID
}

func (us *UserStore) SetRoleID(roleId int) {
	us.RoleID = roleId
}

func (us *UserStore) GetRoleID() int {
	return us.RoleID
}

func (us *UserStore) SetStoreID(storeId int) {
	us.StoreID = storeId
}

func (us *UserStore) GetStoreID() int {
	return us.StoreID
}

func (us *UserStore) SetEnabled(enabled bool) {
	us.Enabled = enabled
}

func (us *UserStore) GetEnabled() bool {
	return us.Enabled
}

func (us *UserStore) IsEnabled() bool {
	return us.GetEnabled()
}

func (us *UserStore) SetCreatedAt(createdAt time.Time) {
	created := timeToString(createdAt)
	us.CreatedAt = created
}

func (us *UserStore) GetCreatedAt() time.Time {
	created, _ := stringToTime(us.CreatedAt)
	return created
}

func (us *UserStore) SetUpdatedAt(updatedAt time.Time) {
	updated := timeToString(updatedAt)
	us.UpdatedAt = &updated
}

func (us *UserStore) GetUpdatedAt() time.Time {
	if us.UpdatedAt == nil {
		return time.Time{}
	}

	updatedAt, _ := stringToTime(*us.UpdatedAt)
	return updatedAt
}

func (us *UserStore) SetDeleted(deleted bool) {
	us.Deleted = deleted
}

func (us *UserStore) GetDeleted() bool {
	return us.Deleted
}

func (us *UserStore) SetStore(store *Store) {
	dataByte, err := json.Marshal(store)
	if err != nil {
		return
	}

	us.Store = dataByte
}

func (us *UserStore) GetStore() *Store {
	var store *Store

	err := json.Unmarshal(us.Store, &store)
	if err != nil {
		return nil
	}

	return store
}
