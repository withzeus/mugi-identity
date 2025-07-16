package pgsql

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
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

type RootDB struct {
	pool *pgxpool.Pool
}

func NewRootDB(p *pgxpool.Pool) RootDB {
	return RootDB{pool: p}
}

func (dbs RootDB) GetPool() *pgxpool.Pool {
	return dbs.pool
}

func NewPgxPool(db DBConfig) (*pgxpool.Pool, func(), error) {
	f := func() {}
	p, err := pgxpool.New(context.Background(), db.GetCS())

	if err != nil {
		return nil, f, fmt.Errorf("connection failure")
	}

	err = PingPoolInfo(p)
	if err != nil {
		return nil, f, err
	}

	return p, func() { p.Close() }, nil
}

type IPgx interface {
	Begin(context.Context) (pgx.Tx, error)
	Exec(context.Context, string, ...any) (pgconn.CommandTag, error)
	QueryRow(context.Context, string, ...any) pgx.Row
	Query(context.Context, string, ...any) (pgx.Rows, error)
	Ping(context.Context) error
	Close()
}

func PingPoolInfo(pool *pgxpool.Pool) error {
	err := pool.Ping(context.Background())

	if err != nil {
		return fmt.Errorf("connection error")
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
		return fmt.Errorf("no rows were returned")
	case err != nil:
		return fmt.Errorf("connection error")
	default:
		log.Printf("database version: %s\n", dbVersion)
		log.Printf("current database user: %s\n", currentUser)
		log.Printf("current database: %s\n", currentDatabase)
	}
	return nil
}
