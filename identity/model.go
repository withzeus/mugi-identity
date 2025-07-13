package identity

import (
	"fmt"
)

type Model struct {
	UID         string
	Handle      string
	Email       string
	PhoneNumber string
	PassKey     string
}

func (i *Model) TableName() string {
	return "uids"
}

func (i *Model) Validate() error {
	if i.UID == "" ||
		(i.Email == "" && i.PhoneNumber == "") ||
		i.Handle == "" || i.PassKey == "" {
		return fmt.Errorf("identity: data invalid")
	}
	return nil
}

type JsonResponse struct {
	Handle      string `json:"handle"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	PassKey     string `json:"passkey"`
}

func (i *Model) ToJsonResponse() *JsonResponse {
	return &JsonResponse{
		Handle:      i.Handle,
		Email:       i.Email,
		PhoneNumber: i.PhoneNumber,
		PassKey:     i.PassKey,
	}
}
