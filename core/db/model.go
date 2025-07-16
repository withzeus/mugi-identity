package db

import (
	"github.com/oklog/ulid/v2"
)

type Model struct{}

func (i *Model) ULID() string {
	return ulid.Make().String()
}
