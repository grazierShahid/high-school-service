package database

import (
	"context"
	"fmt"

	"github.com/grazierShahid/high-school-management-system/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	Pool *pgxpool.Pool
}

func NewDatabase(cfg *config.Config) (*Database, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DB.USERNAME,
		cfg.DB.PASSWORD,
		cfg.DB.HOST,
		cfg.DB.PORT,
		cfg.DB.NAME,
	)

	pool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		return nil, fmt.Errorf("error creating connection pool: %v", err)
	}

	if err := pool.Ping(context.Background()); err != nil {
		pool.Close()
		return nil, fmt.Errorf("error pinging database: %v", err)
	}
	return &Database{Pool: pool}, nil
}

func (db *Database) Close() {
	if db.Pool != nil {
		db.Pool.Close()
	}
}

func (db *Database) GetPool() *pgxpool.Pool {
	return db.Pool
}
