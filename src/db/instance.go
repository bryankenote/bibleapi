package db

import (
	"context"
	sql "database/sql"
	"fmt"
	"github.com/bryankenote/bibleapi/src/codegen/sqlc"
	"log"
	"os"

	_ "embed"

	_ "github.com/mattn/go-sqlite3"
)

type DBInstance struct {
	*sqlc.Queries
	db *sql.DB
}

var Instance *DBInstance

func ConnectToDB() {
	var err error
	Instance, err = newDBInstance()

	if err != nil {
		log.Fatalf("unable to connect to database: %v", err)
	}
}

//go:embed schema.sql
var ddl string

func newDBInstance() (*DBInstance, error) {
	ctx := context.Background()

	db, err := sql.Open("sqlite3", "bibleapi.db")
	if err != nil {
		return nil, err
	}

	if _, err := os.Stat("bibleapi.db"); err != nil {
		// create tables
		if _, err = db.ExecContext(ctx, ddl); err != nil {
			return nil, err
		}
	}

	return &DBInstance{
		db:      db,
		Queries: sqlc.New(db),
	}, nil
}

func (store *DBInstance) ExecTx(ctx context.Context, fn func(*sqlc.Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := sqlc.New(tx)
	err = fn(q)
	if err != nil {
		if rollbackErr := tx.Rollback(); err != nil {
			return fmt.Errorf("tx err: %v, rollback err: %v", err, rollbackErr)
		}
		return err
	}

	return tx.Commit()
}

func NewNullString(s string) sql.NullString {
	if len(s) == 0 {
		return sql.NullString{}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}
