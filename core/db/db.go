package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DBConfig struct {
	Username string
	Password string
	Hostname string
	Port     string
	DBName   string
}

func (db DBConfig) GetCS() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		db.Username, db.Password, db.Hostname, db.Port, db.DBName)
}

type DBS struct {
	pool *pgxpool.Pool
}

func NewDBS(p *pgxpool.Pool) DBS {
	return DBS{pool: p}
}

func (dbs DBS) GetPool() *pgxpool.Pool {
	return dbs.pool
}

func NewDBP(db DBConfig) (*pgxpool.Pool, func(), error) {
	f := func() {}
	p, err := pgxpool.New(context.Background(), db.GetCS())

	if err != nil {
		return nil, f, errors.New("db: connection failure")
	}

	err = validateDBP(p)
	if err != nil {
		return nil, f, err
	}

	return p, func() { p.Close() }, nil
}

func validateDBP(pool *pgxpool.Pool) error {
	err := pool.Ping(context.Background())

	if err != nil {
		return errors.New("db: connection error")
	}

	var (
		currentDatabase string
		currentUser     string
		dbVersion       string
	)

	sqlStatement := `select current_database(), current_user, version();`
	row := pool.QueryRow(context.Background(), sqlStatement)
	err = row.Scan(&currentDatabase, &currentUser, &dbVersion)

	switch {
	case err == sql.ErrNoRows:
		return errors.New("no rows were returned")
	case err != nil:
		return errors.New("database connection error")
	default:
		log.Printf("database version: %s\n", dbVersion)
		log.Printf("current database user: %s\n", currentUser)
		log.Printf("current database: %s\n", currentDatabase)
	}
	return nil
}
