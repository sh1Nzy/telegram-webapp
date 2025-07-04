package models

import "time"

type User struct {
	ID         int64     `db:"id" json:"id"`
	TelegramID int64     `db:"telegram_id" json:"telegram_id"`
	Username   string    `db:"username" json:"username"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
}

type CreateUser struct {
	TelegramID int64  `json:"telegram_id"`
	Username   string `json:"username"`
}
