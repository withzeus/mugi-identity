package identity

import (
	"fmt"

	"github.com/withzeus/mugi-identity/core/db"
)

type Model struct {
	UID         string
	Handle      string `json:"handle"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	PassKey     string `json:"passkey"`
	CreatedAt   string
	UpdatedAt   string
	DeletedAt   string
	db.Model
}

func (i *Model) TableName() string {
	return "users"
}

func (i *Model) Validate() error {
	if (i.Email == "" && i.PhoneNumber == "") ||
		i.Handle == "" || i.PassKey == "" {
		return fmt.Errorf("bad request")
	}
	return nil
}

type Response struct {
	Handle      string `json:"handle"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	PassKey     string `json:"passkey"`
}

func (i *Model) ToResponse() *Response {
	return &Response{
		Handle:      i.Handle,
		Email:       i.Email,
		PhoneNumber: i.PhoneNumber,
		PassKey:     i.PassKey,
	}
}
