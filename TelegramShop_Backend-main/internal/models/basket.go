package models

import "time"

type BasketItem struct {
	UserID    int64     `db:"user_id" json:"user_id"`
	ProductID int       `db:"product_id" json:"product_id"`
	Quantity  int       `db:"quantity" json:"quantity"`
	AddedAt   time.Time `db:"added_at" json:"added_at"`
}

type CreateBasketItem struct {
	UserID    int64 `json:"user_id"`
	ProductID int   `json:"product_id"`
	Quantity  int   `json:"quantity"`
}

type DeleteBasketItem struct {
	UserID    int64 `json:"user_id"`
	ProductID int   `json:"product_id"`
}
