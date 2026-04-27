package store

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Follower struct {
	UserID int64 `json:"user_id"`
	FollowerID int64 `json:"follower_id"`
	CreatedAt time.Time `json:"created_at"`
}

type FollowerStore struct {
	pool *pgxpool.Pool
}

func (s *FollowerStore) Follow(ctx context.Context, followerID, userID int64) error {
	query := `
		INSERT INTO followers (user_id, follower_id)
		VALUES ($1, $2);
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeout)
	defer cancel()

	if _, err := s.pool.Exec(ctx, query, userID, followerID); err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return ErrConflict
			}
		}

		return err
	}

	return nil
}
		

func (s *FollowerStore) UnFollow(ctx context.Context, followerID, userID int64) error {
	query := `
		DELETE FROM followers
		WHERE user_id = $1 AND follower_id = $2;
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeout)
	defer cancel()

	if _, err := s.pool.Exec(ctx, query, userID, followerID); err != nil {
		return err
	}

	return nil
}

