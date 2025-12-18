package user

import "time"

type User struct {
	ID           int64   `db:"id" json:"id"`
	Email        string  `db:"email" json:"email"`
	Username     string  `db:"username" json:"username"`
	PasswordHash string  `db:"password_hash" json:"-"`
	Role         string  `db:"role" json:"role"`
	IsBlocked    bool    `db:"is_blocked" json:"is_blocked"`
	Store        *string `db:"store" json:"store"`
	Phone        string  `db:"phone" json:"phone"`

	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
