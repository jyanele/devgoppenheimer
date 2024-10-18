package database

import (
	"context"
	"fmt"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
)

type postgres struct {
	db *pgxpool.Pool
}

var (
	pgInstance *postgres
	pgOnce sync.Once
)

func NewPG(ctx context.Context, connString string) (*postgres, error) {
	var err error
	pgOnce.Do(func() {
			db, e := pgxpool.New(ctx, connString)
			if e != nil {
				err = fmt.Errorf("unable to create connection pool: %w", err)
				return 
			}
			pgInstance = &postgres{db}
		})
	if err != nil {
		return nil, err
	}
	return pgInstance, nil
}

func (pg *postgres) Ping(ctx context.Context) error {
	return pg.db.Ping(ctx)
}

func (pg *postgres) Close() {
	pg.db.Close()
}