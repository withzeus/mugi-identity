package ext

import "github.com/oklog/ulid/v2"

type ULID struct{}

func (u *ULID) NewUILD() string {
	return ulid.Make().String()
}
