package store

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  Password    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (p *Password) Set(text string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)
	if err != nil {
		return nil
	}

	p.Text = &text
	p.Hash = hash

	return nil
}


type Password struct {
	Text *string
	Hash []byte
}

type UserStore struct {
	pool *pgxpool.Pool
}

func (s *UserStore) Create(ctx context.Context, tx pgx.Tx, user *User) error {
	query := `
		INSERT INTO users (username, email, password)
		VALUES ($1, $2, $3)
		RETURNING id, created_at;
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeout)
	defer cancel()

	if err := s.pool.QueryRow(
		ctx,
		query,
		user.Username,
		user.Email,
		user.Password,
	).Scan(
		&user.ID,
		&user.CreatedAt,
	); err != nil {
		return err
	}

	return nil
}

func (s *UserStore) GetByID(ctx context.Context, userID int64) (*User, error) {
	query := `
		SELECT id, username, email, password, created_at
		FROM users
		WHERE id = $1;
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeout)
	defer cancel()

	var user User

	if err := s.pool.QueryRow(ctx, query, userID).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	); err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}

	return &user, nil

}

func (s *UserStore) CreateAndInvite(ctx context.Context, user *User, token string, invitationExp time.Duration) (error) {
	return withTx(s.pool, ctx, func(tx pgx.Tx) error {
		if err := s.Create(ctx, tx, user); err != nil {
			return err
		}

		if err := s.CreateUserInvitation(ctx, tx, token, invitationExp, user.ID); err != nil {
			return err
		}
		
		return nil
	})
}

func (s *UserStore) CreateUserInvitation(ctx context.Context, tx pgx.Tx, token string, exp time.Duration, userID int64) error {
	query := `
		INSERT INTO user_invitations (token, user_id, expiry) 
		VALUES ($1, $2, $3);
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeout)
	defer cancel()

	if _, err := tx.Exec(ctx, query, userID, userID, time.Now().Add(exp)); err != nil {
		return err
	}
	
	return nil
}