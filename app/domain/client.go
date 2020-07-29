package domain

import "time"

type Client struct {
	ID          uint       `gorm:"primary_key" json:"id"`
	ClientID    string     `gorm:"not null; unique" json:"clientId"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
	DeletedAt   *time.Time `sql:"index" json:"deletedAr"`
	Name        string     `gorm:"not null" json:"name"`
	RedirectURI string     `gorm:"not null" json:"redirectUri"`
}
