package db

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primaryKey"`
	Handle    string         `gorm:"uniqueIndex"`
	FullName  string         `json:"-"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

func (c *User) Create(o *ORM) Result {
	statement := o.Create(c)
	return o.CreateResult(statement)
}

func (c *User) QueryByHandle(o *ORM) Result {
	statement := o.Where("handle = ?", c.Handle).First(c)
	return o.CreateResult(statement)
}
