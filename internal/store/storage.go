package store

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	ErrNotFound = errors.New("records not found")
)

type Storage struct {
	Posts interface {
		Create(context.Context, *Post) error
		GetByID(context.Context, int64) (*Post, error)
		DeleteByID(context.Context, int64) error
		UpdateByID(context.Context, int64, *Post) error
	}
	Users interface {
		Create(context.Context, *User) error
	}
	Comments interface{
		GetByPostID(context.Context, int64) ([]Comments, error)
	}
}

func NewStorage(pool *pgxpool.Pool) Storage {
	return Storage{
		Posts: &PostStore{pool},
		Users: &UserStore{pool},
		Comments: &CommentStore{pool},
	}
}