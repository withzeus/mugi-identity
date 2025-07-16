package tenant

import (
	"fmt"

	"github.com/withzeus/mugi-identity/core/db"
)

type Model struct {
	ID          string
	Name        string `json:"name"`
	Secret      string
	Website     string `json:"website"`
	Logo        string `json:"logo"`
	RedirectUri string `json:"redirect_uri"`
	CreatedAt   string
	UpdatedAt   string
	DeletedAt   string
	db.Model
}

func (m *Model) TableName() string {
	return "tenants"
}

func (m *Model) Validate() error {
	if m.Name == "" || m.RedirectUri == "" || m.Website == "" {
		return fmt.Errorf("bad request")
	}
	return nil
}

type Response struct {
	Name        string `json:"name"`
	Website     string `json:"website"`
	Logo        string `json:"logo"`
	RedirectUri string `json:"redirect_uri"`
}

func (m *Model) ToResponse() *Response {
	return &Response{
		Name:        m.Name,
		Website:     m.Website,
		Logo:        m.Logo,
		RedirectUri: m.RedirectUri,
	}
}
