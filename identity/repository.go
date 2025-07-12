package identity

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type IPgx interface {
	Begin(context.Context) (pgx.Tx, error)
	Exec(context.Context, string, ...any) (pgconn.CommandTag, error)
	QueryRow(context.Context, string, ...any) pgx.Row
	Query(context.Context, string, ...any) (pgx.Rows, error)
	Ping(context.Context) error
	Close()
}

type DB struct {
	pgx IPgx
}

func NewDB(db IPgx) DB {
	return DB{pgx: db}
}

func (db DB) Create(u Identity) (*Identity, error) {
	q := `INSERT INTO uids (handle,email,phone_number,passkey)
          VALUES ($1,$2,$3,$4) RETURNING uid,handle,email,phone_number,passkey`

	r := db.pgx.QueryRow(context.Background(), q, u.Handle, u.Email, u.PhoneNumber, u.PassKey)
	i := new(Identity)

	if err := r.Scan(
		&u.UID,
		&u.Handle,
		&u.Email,
		&u.PhoneNumber,
		&u.PassKey,
	); err != nil {
		return nil, fmt.Errorf("identity: %s", err.Error())
	}
	return i, nil
}
