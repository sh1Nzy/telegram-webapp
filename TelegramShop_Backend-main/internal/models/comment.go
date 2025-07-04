package models

type Comment struct {
	ID        int    `db:"id" json:"id"`
	UserID    int64  `db:"user_id" json:"user_id"`
	ProductID int    `db:"product_id" json:"product_id"`
	Comment   string `db:"comment" json:"comment"`
	CreatedAt string `db:"created_at" json:"created_at"`
}
