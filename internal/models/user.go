package models

import uuid "github.com/jackc/pgx/pgtype/ext/satori-uuid"

type User struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Nickname string    `json:"nickname"`
	Email    string    `json:"email"`
	// TODO add as needed
}
