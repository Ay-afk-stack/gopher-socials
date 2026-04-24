package store

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Post struct{
	ID int64 `json:"id"`
	Content string `json:"content"`
	Title string `json:"title"`
	UserID int64 `json:"user_id"`
	Tags []string `json:"tags"`
	Comments []Comments `json:"comments"`
	CreatedAt time.Time	`json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PostStore struct {
	pool *pgxpool.Pool
}

func(s *PostStore) Create(ctx context.Context, post *Post) error {
	query := `
		INSERT INTO posts (content, title, user_id, tags)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at, updated_at;
	`

	if err := s.pool.QueryRow(
		ctx,
		query,
		post.Content,
		post.Title,
		post.UserID,
		post.Tags,
		).Scan(
			&post.ID,
			&post.CreatedAt,
			&post.UpdatedAt,
		); err != nil {
			return err
		}

		return nil
}

func (s *PostStore) GetByID(ctx context.Context, id int64) (*Post, error) {
	query := `
		SELECT id, title, content, tags, user_id, created_at, updated_at
		FROM posts
		WHERE id = $1;
	`

	var post Post

	if err := s.pool.QueryRow(ctx, query, id).Scan(
		&post.ID,
		&post.Title,
		&post.Content,
		&post.Tags,
		&post.UserID,
		&post.CreatedAt,
		&post.UpdatedAt,
	); err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}

	return &post, nil
}

func (s *PostStore) DeleteByID(ctx context.Context, id int64) error {
	query := `
		DELETE FROM posts
		WHERE id = $1;
	`

	cmd, err := s.pool.Exec(ctx, query, id)
	if err != nil {
		return nil
	}

	if cmd.RowsAffected() == 0 {
		return ErrNotFound
	}

	return nil
}

func (s *PostStore) UpdateByID(ctx context.Context, id int64, post *Post) error {
	query := `
		UPDATE posts
		SET
			title = $1,
			content = $2
		WHERE id = $3
	`

	cmd, err := s.pool.Exec(ctx, query, post.Title, post.Content, post.ID)
	if err != nil {
		return err
	}

	if cmd.RowsAffected() == 0 {
		return ErrNotFound
	}

	return nil

}