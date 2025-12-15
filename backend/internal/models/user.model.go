package models

import "time"

type User struct {
	ID        int64   `db:"id" json:"id"`
	Username  string  `db:"username" json:"username"`
	Role      string  `db:"role" json:"role"`
	IsBlocked bool    `db:"is_blocked" json:"is_blocked"`
	Store     *string `db:"store" json:"store"`
	Phone     string  `db:"phone" json:"phone"`

	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
