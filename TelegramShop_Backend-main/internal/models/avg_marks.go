package models

type AvgMarks struct {
	ProductID int     `db:"product_id" json:"product_id"`
	Sum       float64 `db:"sum" json:"sum"`
	Count     int     `db:"count" json:"count"`
}
