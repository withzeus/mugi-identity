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
	q := `INSERT INTO uids (handle,email,phone_number,passkey)
          VALUES ($1,$2,$3,$4) RETURNING uid,handle,email,phone_number,passkey`

	r := db.pgx.QueryRow(context.Background(), q, md.Handle, md.Email, md.PhoneNumber, md.PassKey)
	i := new(Model)

	if err := r.Scan(
		&md.UID,
		&md.Handle,
		&md.Email,
		&md.PhoneNumber,
		&md.PassKey,
	); err != nil {
		return nil, fmt.Errorf("pkg identity: %s", err.Error())
	}
	return i, nil
}
