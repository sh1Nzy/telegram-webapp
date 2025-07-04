package products

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"telegramshop_backend/internal/models"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type Repository interface {
	CreateProduct(ctx context.Context, product models.Product) (models.Product, error)
	GetProductByID(ctx context.Context, id int64) (models.Product, error)
	GetAllProducts(ctx context.Context) ([]models.Product, error)
	UpdateProduct(ctx context.Context, id int64, product models.UpdateProductInput) error
	DeleteProduct(ctx context.Context, id int64) error
	AddProductImage(ctx context.Context, id int64, imageURL string) error
	RemoveProductImage(ctx context.Context, id int64, imageURL string) error
	SetProductImages(ctx context.Context, id int64, images []string) error

	IncrementSellCount(ctx context.Context, productID int64, count int) error
	UpdateStock(ctx context.Context, productID int64, stock int) error
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{db: db}
}

func scanProductRow(row *sqlx.Row, product *models.Product) error {
	var attrs []byte
	err := row.Scan(
		&product.ID,
		&product.Name,
		&product.FirmID,
		&product.Description,
		&product.CategoryID,
		&attrs,
		&product.SellCount,
		&product.Stock,
		&product.Image,
	)
	if err != nil {
		return err
	}
	return json.Unmarshal(attrs, &product.Attributes)
}

func (r *repository) CreateProduct(ctx context.Context, product models.Product) (models.Product, error) {
	attrs, err := json.Marshal(product.Attributes)
	if err != nil {
		return models.Product{}, err
	}

	query := `
		INSERT INTO products (name, firm_id, description, category_id, attributes, sell_count, stock, image)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id`

	err = r.db.QueryRowContext(ctx, query,
		product.Name,
		product.FirmID,
		product.Description,
		product.CategoryID,
		attrs,
		product.SellCount,
		product.Stock,
		product.Image,
	).Scan(&product.ID)

	return product, err
}

func (r *repository) GetProductByID(ctx context.Context, id int64) (models.Product, error) {
	query := `
		SELECT id, name, firm_id, description, category_id, attributes, sell_count, stock, image
		FROM products
		WHERE id = $1`

	var product models.Product
	row := r.db.QueryRowxContext(ctx, query, id)

	err := scanProductRow(row, &product)
	if errors.Is(err, sql.ErrNoRows) {
		return models.Product{}, nil
	}
	return product, err
}

func (r *repository) GetAllProducts(ctx context.Context) ([]models.Product, error) {
	query := `
		SELECT id, name, firm_id, description, category_id, attributes, sell_count, stock, image
		FROM products`

	rows, err := r.db.QueryxContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var p models.Product
		var attrs []byte
		err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.FirmID,
			&p.Description,
			&p.CategoryID,
			&attrs,
			&p.SellCount,
			&p.Stock,
			&p.Image,
		)
		if err != nil {
			return nil, err
		}
		if err := json.Unmarshal(attrs, &p.Attributes); err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}

func (r *repository) UpdateProduct(ctx context.Context, id int64, product models.UpdateProductInput) error {
	attrs, err := json.Marshal(product.Attributes)
	if err != nil {
		return err
	}

	query := `
		UPDATE products
		SET name = $1,
			firm_id = $2,
			description = $3,
			category_id = $4,
			attributes = $5,
			stock = $6,
			image = $7
		WHERE id = $8`

	_, err = r.db.ExecContext(ctx, query,
		product.Name,
		product.FirmID,
		product.Description,
		product.CategoryID,
		attrs,
		product.Stock,
		product.Image,
		id,
	)
	return err
}

func (r *repository) DeleteProduct(ctx context.Context, id int64) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM products WHERE id = $1`, id)
	return err
}

func (r *repository) AddProductImage(ctx context.Context, id int64, imageURL string) error {
	_, err := r.db.ExecContext(ctx,
		`UPDATE products SET image = array_append(image, $1) WHERE id = $2`,
		imageURL, id)
	return err
}

func (r *repository) RemoveProductImage(ctx context.Context, id int64, imageURL string) error {
	_, err := r.db.ExecContext(ctx,
		`UPDATE products SET image = array_remove(image, $1) WHERE id = $2`,
		imageURL, id)
	return err
}

func (r *repository) SetProductImages(ctx context.Context, id int64, images []string) error {
	_, err := r.db.ExecContext(ctx,
		`UPDATE products SET image = $1 WHERE id = $2`,
		pq.StringArray(images), id)
	return err
}

func (r *repository) IncrementSellCount(ctx context.Context, productID int64, count int) error {
	_, err := r.db.ExecContext(ctx, `UPDATE products SET sell_count = sell_count + $1 WHERE id = $2`, count, productID)
	return err
}

func (r *repository) UpdateStock(ctx context.Context, productID int64, stock int) error {
	_, err := r.db.ExecContext(ctx, `UPDATE products SET stock = $1 WHERE id = $2`, stock, productID)
	return err
}
