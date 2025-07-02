package db

import (
	"github.com/wader/gormstore/v2"
	"github.com/withzeus/mugi-identity/helpers"
)

type SessionStore struct {
	*gormstore.Store
}

func NewSessionStore(orm *ORM) *SessionStore {
	secret, _ := helpers.RandomHexEncoded(32)
	store := gormstore.New(orm.DB, []byte(secret))
	return &SessionStore{store}
}
