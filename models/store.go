package models

import (
	"spos/users/constant"
	"time"
)

type Store struct {
	ID            int     `db:"id"`
	OwnerID       int     `db:"owner_id"`
	RefShopID     *int    `db:"ref_shop_id"`
	RefUserID     *int    `db:"ref_user_id"`
	RefShopStatus *string `db:"ref_shop_status"`
	Name          string  `db:"name"`
	Domain        *string `db:"domain"`
	Logo          string  `db:"logo"`
	Description   *string `db:"description"`
	Enabled       bool    `db:"enabled"`
	Type          string  `db:"type"`
	Source        string  `db:"source"`
	CreatedAt     string  `db:"created_at"`
	UpdatedAt     *string `db:"updated_at"`
	Deleted       bool    `db:"deleted"`

	// relation table
	Owner []byte `db:"owner"`
}

func NewStore() *Store {
	// let time now is UTC +0
	now := time.Now().UTC()

	return &Store{
		Enabled:   !enabled,
		Type:      constant.Offline,
		Source:    constant.Offline,
		CreatedAt: now.Format(layout),
	}
}

func (s *Store) GetId() int {
	return s.ID
}

func (s *Store) SetOwnerID(ownerId int) {
	s.OwnerID = ownerId
}

func (s *Store) GetOwnerID() int {
	return s.OwnerID
}

func (s *Store) SetRefShopID(refShopId int) {
	s.RefShopID = &refShopId
}

func (s *Store) GetRefShopID() int {
	if s.RefShopID == nil {
		return defaultInt
	}
	return *s.RefShopID
}

func (s *Store) SetRefUserID(refUserId int) {
	s.RefUserID = &refUserId
}

func (s *Store) GetRefUserID() int {
	if s.RefUserID == nil {
		return defaultInt
	}
	return *s.RefUserID
}

func (s *Store) SetRefShopStatus(refShopStatus string) {
	s.RefShopStatus = &refShopStatus
}

func (s *Store) GetRefShopStatus() string {
	if s.RefShopStatus == nil {
		return defaultString
	}
	return *s.RefShopStatus
}

func (s *Store) SetName(name string) {
	s.Name = name
}

func (s *Store) GetName() string {
	return s.Name
}

func (s *Store) SetDomain(domain string) {
	s.Domain = &domain
}

func (s *Store) GetDomain() string {
	if s.Domain == nil {
		return defaultString
	}
	return *s.Domain
}

func (s *Store) SetLogo(logo string) {
	s.Logo = logo
}

func (s *Store) GetLogo() string {
	return s.Logo
}

func (s *Store) SetDescription(description string) {
	s.Description = &description
}

func (s *Store) GetDescription() string {
	if s.Description == nil {
		return defaultString
	}
	return *s.Description
}

func (s *Store) SetEnabled(enabled bool) {
	s.Enabled = enabled
}

func (s *Store) GetEnabled() bool {
	return s.Enabled
}

func (s *Store) IsEnabled() bool {
	return s.GetEnabled()
}

func (s *Store) SetType(types string) {
	s.Type = types
}

func (s *Store) GetType() string {
	return s.Type
}

func (s *Store) IsTypeOffline() bool {
	return s.GetType() == constant.Offline
}

func (s *Store) IsTypeOnline() bool {
	return s.GetType() == constant.Online
}

func (s *Store) SetSource(sourcee string) {
	s.Source = sourcee
}

func (s *Store) GetSource() string {
	return s.Source
}

func (s *Store) IsSourceOffline() bool {
	return s.GetSource() == constant.Offline
}

func (s *Store) IsSourceTokopedia() bool {
	return s.GetSource() == constant.Tokopedia
}

func (s *Store) IsSourceShopee() bool {
	return s.GetSource() == constant.Shopee
}

func (s *Store) SetCreatedAt(CreatedAt time.Time) {
	createdAt := timeToString(CreatedAt)
	s.CreatedAt = createdAt
}

func (s *Store) GetCreatedAt() time.Time {
	created, _ := stringToTime(s.CreatedAt)
	return created
}

func (s *Store) SetUpdatedAt(UpdatedAt time.Time) {
	updatedAt := timeToString(UpdatedAt)
	s.UpdatedAt = &updatedAt
}

func (s *Store) GetUpdatedAt() time.Time {
	if s.UpdatedAt == nil {
		return time.Time{}
	}

	updatedAt, _ := stringToTime(*s.UpdatedAt)
	return updatedAt
}

func (s *Store) SetDeleted(deleted bool) {
	s.Deleted = deleted
}

func (s *Store) GetDeleted() bool {
	return s.Deleted
}
