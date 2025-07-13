package identity

import (
	"fmt"

	"github.com/oklog/ulid/v2"
)

type Model struct {
	UID         string
	Handle      string
	Email       string
	PhoneNumber string
	PassKey     string
}

func (i *Model) TableName() string {
	return "users"
}

func (i *Model) ULID() string {
	return ulid.Make().String()
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
