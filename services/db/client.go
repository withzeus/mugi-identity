package db

import (
	"time"

	"gorm.io/gorm"
)

type Client struct {
	ID           uint   `gorm:"primaryKey"`
	Name         string `gorm:"uniqueIndex"`
	ClientSecret string `json:"-"`
	Website      string
	Logo         string
	RedirectURI  string         `json:"redirect_uri"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"-" gorm:"index"`
}

func (c *Client) Create(o *ORM) error {
	statement := o.Create(c)
	return statement.Error
}
