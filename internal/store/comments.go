package store

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Comment struct {
	ID        int64     `json:"id"`
	PostID    int64     `json:"post_id"`
	UserID    int64     `json:"user_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	User      User      `json:"user"`
}

type CommentStore struct {
	pool *pgxpool.Pool
}

func (s *CommentStore) GetByPostID(ctx context.Context, id int64) ([]Comment, error) {
	query := `
		SELECT 
			c.id,
			c.post_id,
			c.user_id,
			c.content,
			c.created_at,
			u.id,
			u.username
		FROM comments c
		JOIN users u 
		ON u.id = c.user_id
		WHERE c.post_id = $1
		ORDER BY c.created_at DESC;
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeout)
	defer cancel()

	rows, err := s.pool.Query(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	comments := []Comment{}

	for rows.Next() {
		c := Comment{}
		if err := rows.Scan(
			&c.ID,
			&c.PostID,
			&c.UserID,
			&c.Content,
			&c.CreatedAt,
			&c.User.ID,
			&c.User.Username,
		); err != nil {
			return nil, err
		}

		comments = append(comments, c)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return comments, nil
}

func (s *CommentStore) Create(ctx context.Context, comment *Comment) error {
	query := `
		INSERT INTO comments(post_id, user_id, content)
		VALUES ($1, $2, $3)
		RETURNING id, created_at;
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeout)
	defer cancel()

	if err := s.pool.QueryRow(ctx, query, comment.PostID, comment.UserID, comment.CreatedAt).Scan(&comment.ID, &comment.CreatedAt); err != nil {
		return err
	}

	return nil

}
