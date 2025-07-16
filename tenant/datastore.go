package tenant

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
	q := `INSERT INTO clients (id,name,secret,website,logo,redirect_uri)
          VALUES ($1,$2,$3,$4,$5,$6) RETURNING id,name,website,logo,redirect_uri`

	r := db.pgx.QueryRow(context.Background(), q, md.ULID(), md.Name, md.Secret, md.Website, md.Logo, md.RedirectUri)
	c := new(Model)

	if err := r.Scan(
		&c.ID,
		&c.Name,
		&c.Website,
		&c.Logo,
		&c.RedirectUri,
	); err != nil {
		return nil, fmt.Errorf("datastore error")
	}
	return c, nil
}
