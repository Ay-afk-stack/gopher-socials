package store

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type User struct {
	ID int64 `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserStore struct {
	pool *pgxpool.Pool
}

func(s *UserStore) Create(ctx context.Context, user *User) error {
	query := `
		INSERT INTO users (username, email, password)
		VALUES ($1, $2, $3)
		RETURNING id, created_at;
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeout)
	defer cancel()

	if err := s.pool.QueryRow(
		ctx,
		query,
		user.Username,
		user.Email,
		user.Password,
	).Scan(
		&user.ID,
		&user.CreatedAt,
	); err != nil {
		return err
	}

	return nil
}