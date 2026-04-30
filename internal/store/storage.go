package store

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	ErrNotFound  = errors.New("records not found")
	ErrConflict  = errors.New("records already exist")
	QueryTimeout = time.Second * 5
)

type Storage struct {
	Posts interface {
		Create(context.Context, *Post) error
		GetByID(context.Context, int64) (*Post, error)
		DeleteByID(context.Context, int64) error
		UpdateByID(context.Context, int64, *Post) error
		GetFeed(context.Context, int64, *PaginationFeedQuery) ([]PostWithMetadata, error)
	}
	Users interface {
		Create(context.Context, pgx.Tx, *User) error
		GetByID(context.Context, int64) (*User, error)
		CreateAndInvite(context.Context, *User, string, time.Duration) (error)
		CreateUserInvitation(context.Context, pgx.Tx, string, time.Duration, int64) error
	}
	Comments interface {
		Create(context.Context, *Comment) error
		GetByPostID(context.Context, int64) ([]Comment, error)
	}
	Followers interface {
		Follow(context.Context, int64, int64) error
		UnFollow(context.Context, int64, int64) error
	}
}

func NewStorage(pool *pgxpool.Pool) Storage {
	return Storage{
		Posts:     &PostStore{pool},
		Users:     &UserStore{pool},
		Comments:  &CommentStore{pool},
		Followers: &FollowerStore{pool},
	}
}

func withTx(pool *pgxpool.Pool, ctx context.Context, fn func(pgx.Tx) error) error {
	tx, err := pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}

	if err := fn(tx); err != nil {
		_ = tx.Rollback(ctx)
		return err
	}

	return tx.Commit(ctx)
}