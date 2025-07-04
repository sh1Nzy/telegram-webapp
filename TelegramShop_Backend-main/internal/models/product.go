package models

import "github.com/lib/pq"

type Product struct {
	ID          int64                  `db:"id" json:"id"`
	Name        string                 `db:"name" json:"name"`
	FirmID      int64                  `db:"firm_id" json:"firm_id"`
	Description string                 `db:"description" json:"description"`
	CategoryID  int64                  `db:"category_id" json:"category_id"`
	Attributes  map[string]interface{} `db:"attributes" json:"attributes" swaggertype:"object"`
	SellCount   int                    `db:"sell_count" json:"sell_count"`
	Stock       int                    `db:"stock" json:"stock"`
	Image       pq.StringArray         `db:"image" json:"image" swaggertype:"array,string" example:"[\"https://example.com/1.jpg\",\"https://example.com/2.jpg\"]"`
}

type UpdateProductInput struct {
	Name        string                 `db:"name" json:"name"`
	FirmID      int64                  `db:"firm_id" json:"firm_id"`
	Description string                 `db:"description" json:"description"`
	CategoryID  int64                  `db:"category_id" json:"category_id"`
	Attributes  map[string]interface{} `db:"attributes" json:"attributes"`
	Stock       int                    `db:"stock" json:"stock"`
	Image       pq.StringArray         `db:"image" json:"image" swaggertype:"array,string" example:"[\"https://example.com/1.jpg\", \"https://example.com/2.jpg\"]"`
}

type ImageInput struct {
	Image string `json:"image" example:"https://example.com/image.jpg"`
}

type ImagesInput struct {
	Images []string `json:"images"`
}

type CountInput struct {
	Count int `json:"count"`
}
type StockInput struct {
	Stock int `json:"stock"`
}
