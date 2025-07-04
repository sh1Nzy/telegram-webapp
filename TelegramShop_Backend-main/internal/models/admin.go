package models

type Admin struct {
	UserID int64 `db:"user_id" json:"user_id"`
}
