package services

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserService struct {
	db *pgxpool.Pool
}

func NewUserService(db *pgxpool.Pool) *UserService {
	return &UserService{
		db: db,
	}
}

func (s *UserService) Ping(ctx context.Context) {
	err := s.db.Ping(ctx)
	if err != nil {
		fmt.Printf("not ping: %v", err)
	}
}
