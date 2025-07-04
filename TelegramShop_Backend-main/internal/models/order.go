package models

import "time"

type (
	OrderWithProducts struct {
		ID        int64          `db:"id" json:"id"`
		UserID    int64          `db:"user_id" json:"user_id"`
		Status    string         `db:"status" json:"status"`
		CreatedAt time.Time      `db:"created_at" json:"created_at"`
		Products  []OrderProduct `json:"products"`
	}

	Order struct {
		ID          int     `db:"id" json:"id"`
		UserID      int64   `db:"user_id" json:"user_id"`
		CreatedAt   string  `db:"created_at" json:"created_at"`
		TotalAmount float64 `db:"total_amount" json:"total_amount"`
	}

	OrderProduct struct {
		ID        int     `db:"id" json:"id"`
		OrderID   int     `db:"order_id" json:"order_id"`
		ProductID int     `db:"product_id" json:"product_id"`
		Quantity  int     `db:"quantity" json:"quantity"`
		Price     float64 `db:"price" json:"price"`
	}

	CreateOrder struct {
		UserID int64 `json:"user_id"`
		Items  []struct {
			ProductID int `json:"product_id"`
			Quantity  int `json:"quantity"`
		} `json:"items"`
	}
)
