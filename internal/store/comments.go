package store

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Comment struct {
	ID int64 `json:"id"`
	PostID int64 `json:"post_id"`
	UserID int64 `json:"user_id"`
	Content string	`json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

type CommentStore struct {
	pool *pgxpool.Pool
}

func (s *CommentStore) GetByPostID(ctx context.Context, id int64) (*[]Comment, error) {
	query := `
		SELECT 
			* 
		FROM comments c
		JOIN users u 
		ON u.id = c.user_id
		WHERE c.post_id = 2
		ORDER BY c.created_at DESC;
	`
}