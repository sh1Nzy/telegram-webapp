package models

type Marks struct {
	UserID    int64   `db:"user_id" json:"user_id"`
	ProductID int     `db:"product_id" json:"product_id"`
	Mark      float64 `db:"mark" json:"mark"`
	CreatedAt string  `db:"created_at" json:"created_at"`
}
