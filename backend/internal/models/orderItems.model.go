package models

import "time"

type OrderItems struct {
	ID        int64     `db:"id" json:"id"`
	OrderId   int64     `db:"order_id" json:"order_id"`
	ProductID int64     `db:"product_id" json:"product_id"`
	Quantity  int64     `db:"quantity" json:"quantity"`
	Price     int64     `db:"price" json:"price"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}
