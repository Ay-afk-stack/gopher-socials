package store

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Comments struct {
	ID int64 `json:"id"`
	PostID int64 `json:"post_id"`
	UserID int64 `json:"user_id"`
	Content string	`json:"content"`
	CreatedAt time.Time `json:"created_at"`
	User User `json:"user"`
}

type CommentStore struct {
	pool *pgxpool.Pool
}

func (s *CommentStore) GetByPostID(ctx context.Context, id int64) ([]Comments, error) {
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

	rows, err := s.pool.Query(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	comments := []Comments{}

	for rows.Next(){
		c := Comments{}
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