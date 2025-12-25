package user

import (
	"context"
	"errors"
	"fmt"

	"github.com/barzaevhalid/sotovik/internal/domain"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, user *User) error {

	query := `
	INSERT INTO users (
	username,
	email,
	password_hash,
	store,
	phone
	)
	VALUES($1, $2, $3, $4, $5)
	RETURNING
	id, role, is_blocked
	`

	err := r.db.QueryRow(ctx, query, user.Username, user.Email, user.PasswordHash, user.Store, user.Phone).Scan(&user.ID, &user.Role, &user.IsBlocked)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return domain.ErrUserAlreadyExists
			}
		}
		return fmt.Errorf("create user: %w", err)
	}
	return nil
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*User, error) {
	user := &User{}
	query := `SELECT  id, username, email, password_hash ,role, is_blocked, store, phone FROM users  WHERE email=$1`
	err := r.db.QueryRow(ctx, query, email).Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.Role, &user.IsBlocked, &user.Store, &user.Phone)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrUserNotFound
		}
		return nil, fmt.Errorf("get user by email: %w", err)
	}
	return user, nil
}

func (r *UserRepository) GetById(ctx context.Context, id int64) (*User, error) {
	user := &User{}
	query := `SELECT  id, username, email, role, is_blocked, store, phone FROM users  WHERE id=$1`
	err := r.db.QueryRow(ctx, query, id).Scan(&user.ID, &user.Username, &user.Email, &user.Role, &user.IsBlocked, &user.Store, &user.Phone)

	if err != nil {
		return nil, err
	}
	return user, nil
}
