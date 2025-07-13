package identity

import (
	"context"
	"fmt"

	"github.com/withzeus/mugi-identity/core/db"
)

type Datastore struct {
	pgx db.IPgx
}

func NewDatastore(db db.IPgx) *Datastore {
	return &Datastore{pgx: db}
}

func (db Datastore) Create(md Model) (*Model, error) {
	q := `INSERT INTO users (uid,handle,email,phone_number,passkey)
          VALUES ($1,$2,$3,$4,$5) RETURNING uid,handle,email,phone_number,passkey`

	r := db.pgx.QueryRow(context.Background(), q, md.ULID(), md.Handle, md.Email, md.PhoneNumber, md.PassKey)
	i := new(Model)

	if err := r.Scan(
		&i.UID,
		&i.Handle,
		&i.Email,
		&i.PhoneNumber,
		&i.PassKey,
	); err != nil {
		return nil, fmt.Errorf("datastore error")
	}
	return i, nil
}
