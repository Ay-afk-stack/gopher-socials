package store

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type RoleStore struct {
	pool *pgxpool.Pool 
}

type Role struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Level int `json:"level"`
	Description string `json:"description"`
}

func (s *RoleStore) GetByName(ctx context.Context, roleName string) (*Role, error) {
	query := `
		SELECT id, name, level, description
		FROM roles
		WHERE name = $1;
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeout)
	defer cancel()

	var role Role

	if err := s.pool.QueryRow(ctx, query, roleName).Scan(
		&role.Id,
		&role.Name,
		&role.Level,
		&role.Description,
	); err != nil {
		switch err {
		case pgx.ErrNoRows:
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}

	return &role, nil
}