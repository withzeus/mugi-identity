package tenant

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
		"id",
		"name",
		"secret",
		"website",
		"logo",
		"redirect_uri",
	}
	sel := []string{
		"id",
		"name",
		"website",
		"logo",
		"redirect_uri",
	}
	query := db.NewQueryBuilder(cols).InsertInto(md.TableName()).Select(sel).GetQuery()
	r := ds.pgx.QueryRow(context.Background(), query, md.NewUILD(), md.Name, md.Secret, md.Website, md.Logo, md.RedirectUri)
	log.Printf("[INFO] Datastore - Executed %s", query)

	newMd := new(Model)

	if err := r.Scan(
		&newMd.ID,
		&newMd.Name,
		&newMd.Website,
		&newMd.Logo,
		&newMd.RedirectUri,
	); err != nil {
		log.Printf("[ERROR] Datastore - %+v", err)
		return nil, fmt.Errorf("datastore error")
	}
	return newMd, nil
}
