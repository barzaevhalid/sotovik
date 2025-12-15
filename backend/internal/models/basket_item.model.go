package models

import "time"

type BasketItem struct {
	ID       int64 `db:"id" json:"id"`
	BasketId int64 `db:"basket_id" json:"basket_id"`
	ItemId   int64 `db:"item_id" json:"item_id"`
	Quantity int64 `db:"quantity" json:"quantity"`

	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
