package identity

import (
	"fmt"
)

type IIdentity interface {
	Get()
	Find(uid string)
	Create() error
	Update(uid string) error
}

type Identity struct {
	UID         string
	Handle      string
	Email       string
	PhoneNumber string
	PassKey     string
}

func (i *Identity) TableName() string {
	return "uids"
}

func (i *Identity) Validate() error {
	if i.UID == "" ||
		(i.Email == "" && i.PhoneNumber == "") ||
		i.Handle == "" || i.PassKey == "" {
		return fmt.Errorf("identity: data invalid")
	}
	return nil
}

type IdentityResponse struct {
	Handle      string `json:"handle"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	PassKey     string `json:"passkey"`
}

func (i *Identity) ToResponse() *IdentityResponse {
	return &IdentityResponse{
		Handle:      i.Handle,
		Email:       i.Email,
		PhoneNumber: i.PhoneNumber,
		PassKey:     i.PassKey,
	}
}
