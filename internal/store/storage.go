package store

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	ErrNotFound = errors.New("records not found")
	ErrConflict = errors.New("records already exist")
	QueryTimeout = time.Second * 5
)

type Storage struct {
	Posts interface {
		Create(context.Context, *Post) error
		GetByID(context.Context, int64) (*Post, error)
		DeleteByID(context.Context, int64) error
		UpdateByID(context.Context, int64, *Post) error
		GetFeed(context.Context, int64) ([]PostWithMetadata, error)
	}
	Users interface {
		Create(context.Context, *User) error
		GetByID(context.Context, int64) (*User, error)
	}
	Comments interface{
		Create(context.Context, *Comment) error
		GetByPostID(context.Context, int64) ([]Comment, error)
	}
	Followers interface{
		Follow(context.Context, int64, int64) error
		UnFollow(context.Context, int64, int64) error
	}
}

func NewStorage(pool *pgxpool.Pool) Storage {
	return Storage{
		Posts: &PostStore{pool},
		Users: &UserStore{pool},
		Comments: &CommentStore{pool},
		Followers: &FollowerStore{pool},
	}
}