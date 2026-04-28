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
	Comments []Comment `json:"comments"`
	Version int `json:"version"`
	CreatedAt time.Time	`json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User User `json:"user"`
}

type PostWithMetadata struct {
	Post
	CommentCount int `json:"comment_count"`
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

	ctx, cancel := context.WithTimeout(ctx, QueryTimeout)
	defer cancel()

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
		SELECT id, title, content, tags, user_id, created_at, updated_at, version
		FROM posts
		WHERE id = $1;
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeout)
	defer cancel()

	var post Post

	if err := s.pool.QueryRow(ctx, query, id).Scan(
		&post.ID,
		&post.Title,
		&post.Content,
		&post.Tags,
		&post.UserID,
		&post.CreatedAt,
		&post.UpdatedAt,
		&post.Version,
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

	ctx, cancel := context.WithTimeout(ctx, QueryTimeout)
	defer cancel()

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
			content = $2,
			version = version + 1
		WHERE id = $3 AND version = $4
		RETURNING version;
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeout)
	defer cancel()

	if err := s.pool.QueryRow(
		ctx,
		query,
		post.Title,
		post.Content,
		post.ID,
		post.Version,
	).Scan(&post.Version); err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			return ErrNotFound
		default:
			return err
		}
	}

	return nil

}

func (s *PostStore) GetFeed(ctx context.Context, userID int64) ([]PostWithMetadata, error) {
	query := `
	SELECT
		p.id,
		p.title,
		p."content",
		p.tags,
		p."version",
		p.user_id,
		u.username,
		p.created_at,
		COUNT(c.id) AS comment_counts
	FROM posts p
	LEFT JOIN comments c
	ON c.post_id = p.id
	LEFT JOIN users u
	ON u.id = p.user_id
	JOIN followers f
	ON f.follower_id = p.user_id OR p.user_id = $1
	WHERE f.user_id = $1 OR p.user_id = $1
	GROUP BY p.id, u.username, u.email
	ORDER BY p.created_at DESC;
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeout)
	defer cancel()

	rows, err := s.pool.Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var feeds []PostWithMetadata

	for rows.Next() {
		var p PostWithMetadata
		if err := rows.Scan(
			&p.ID,
			&p.Title,
			&p.Content,
			&p.Tags,
			&p.Version,
			&p.UserID,
			&p.User.Username,
			&p.CreatedAt,
			&p.CommentCount,
		); err != nil {
			return nil, err
		}

		feeds = append(feeds, p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return feeds, err
}