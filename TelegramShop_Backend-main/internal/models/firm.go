package models

type Firm struct {
	ID   int64  `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

type UpdateFirmInput struct {
	Name string `db:"name" json:"name"`
}
