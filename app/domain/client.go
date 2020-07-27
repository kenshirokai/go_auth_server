package domain

import "time"

type Client struct {
	ID          uint       `gorm:"primary_key" json:"id"`
	ClientID    string     `gorm:"not null; unique" json:"clientId"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
	DeletedAt   *time.Time `sql:"index" json:"deletedAr"`
	Name        string     `gorm:"not null" json:"name"`
	RedirectURL string     `gorm:"not null" json:"redirectUrl"`
}
