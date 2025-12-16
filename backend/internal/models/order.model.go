package models

import "time"

type Order struct {
	ID             int64  `db:"id" json:"id"`
	UserId         int64  `db:"user_id" json:"user_id"`
	Status         string `db:"status" json:"status"`
	Total          int64  `db:"total" json:"total"`
	DeliveryName   string `db:"delivery_name" json:"delivery_name"`
	DeliveryPhone  string `db:"delivery_phone" json:"delivery_phone"`
	DeliveryAdress string `db:"delivery_adress" json:"delivery_adress"`

	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
