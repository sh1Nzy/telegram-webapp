package models

import "time"

type Favorite struct {
	UserID    int64     `db:"user_id" json:"user_id"`
	ProductID int       `db:"product_id" json:"product_id"`
	AddedAt   time.Time `db:"added_at" json:"added_at"`
}

type CreateFavorite struct {
	UserID    int64 `json:"user_id"`
	ProductID int   `json:"product_id"`
}

type DeleteFavorite struct {
	UserID    int64 `json:"user_id"`
	ProductID int   `json:"product_id"`
}
