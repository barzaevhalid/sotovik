package models

import "time"

type Item struct {
	ID             int64   `db:"id" json:"id"`
	Title          string  `db:"title" json:"title"`
	Description    *string `db:"description" json:"description"`
	ImageUrl       *string `db:"image_url" json:"image_url"`
	Price          int64   `db:"price" json:"price"`
	WholesalePrice *int64  `db:"wholesale_price" json:"wholesale_price"`
	IsActive       bool    `db:"is_active" json:"is_active"`
	InStock        bool    `db:"in_stock" json:"in_stock"`
	CategoryId     int64   `db:"category_id" json:"category_id"`

	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
