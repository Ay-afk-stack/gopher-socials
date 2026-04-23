package store

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage struct {
	Posts interface {
		Create(ctx context.Context, post *Post) error
	}
	Users interface {
		Create(ctx context.Context, user *User) error
	}
}

func NewStorage(pool *pgxpool.Pool) Storage {
	return Storage{
		Posts: &PostStore{pool},
		Users: &UserStore{pool},
	}
}