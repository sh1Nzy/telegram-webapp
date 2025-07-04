package models

type Price struct {
	ID        int64   `db:"id" json:"id"`
	ProductID int64   `db:"product_id" json:"product_id"`
	Count     int     `db:"count" json:"count"`
	Price     float64 `db:"price" json:"price"`
}

type UpdatePriceInput struct {
	Price float64 `db:"price" json:"price"`
	Count int     `db:"count" json:"count"`
}

type UpdatePriceCount struct {
	NewCount int `json:"new_count" example:"15"`
}
