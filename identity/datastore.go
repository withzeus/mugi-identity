package identity

import (
	"context"
	"fmt"
	"log"

	"github.com/withzeus/mugi-identity/core/db"
	"github.com/withzeus/mugi-identity/core/db/pgsql"
)

type Datastore struct {
	pgx pgsql.IPgx
}

func NewDatastore(db pgsql.IPgx) *Datastore {
	return &Datastore{pgx: db}
}

func (ds Datastore) Create(md Model) (*Model, error) {
	cols := []string{
		"uid",
		"handle",
		"email",
		"phone_number",
		"passkey",
	}
	query := db.NewQueryBuilder(cols).InsertInto(md.TableName()).Select(cols).GetQuery()
	row := ds.pgx.QueryRow(context.Background(), query, md.NewUILD(), md.Handle, md.Email, md.PhoneNumber, md.PassKey)
	log.Printf("[INFO] Datastore - Executed %s", query)

	newMd := new(Model)
	if err := row.Scan(
		&newMd.UID,
		&newMd.Handle,
		&newMd.Email,
		&newMd.PhoneNumber,
		&newMd.PassKey,
	); err != nil {
		log.Printf("[ERROR] Datastore - %+v", err)
		return nil, fmt.Errorf("datastore error")
	}
	return newMd, nil
}
