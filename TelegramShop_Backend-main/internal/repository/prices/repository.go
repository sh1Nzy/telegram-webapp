package prices

import (
	"context"
	"database/sql"

	"telegramshop_backend/internal/models"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	CreatePrice(ctx context.Context, price models.Price) (models.Price, error)
	GetPriceByID(ctx context.Context, id int64) (models.Price, error)
	GetPricesByProductID(ctx context.Context, productID int64) ([]models.Price, error)
	UpdatePrice(ctx context.Context, id int64, price models.UpdatePriceInput) error
	DeletePrice(ctx context.Context, id int64) error
	DeletePricesByProductID(ctx context.Context, productID int64) error
	UpdatePriceCount(ctx context.Context, id int64, newCount int) error
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{db: db}
}

func (r *repository) CreatePrice(ctx context.Context, price models.Price) (models.Price, error) {
	query := `
		INSERT INTO prices (product_id, count, price)
		VALUES ($1, $2, $3)
		RETURNING id`

	err := r.db.QueryRowContext(ctx, query, price.ProductID, price.Count, price.Price).Scan(&price.ID)
	return price, err
}

func (r *repository) GetPriceByID(ctx context.Context, id int64) (models.Price, error) {
	query := `
		SELECT id, product_id, count, price
		FROM prices
		WHERE id = $1`

	var price models.Price
	err := r.db.GetContext(ctx, &price, query, id)
	if err == sql.ErrNoRows {
		return models.Price{}, sql.ErrNoRows
	}
	return price, err
}

func (r *repository) GetPricesByProductID(ctx context.Context, productID int64) ([]models.Price, error) {
	query := `
		SELECT id, product_id, count, price
		FROM prices
		WHERE product_id = $1`

	var prices []models.Price
	err := r.db.SelectContext(ctx, &prices, query, productID)
	return prices, err
}

func (r *repository) UpdatePrice(ctx context.Context, id int64, price models.UpdatePriceInput) error {
	query := `
		UPDATE prices
		SET price = $1, count = $2
		WHERE id = $3`

	_, err := r.db.ExecContext(ctx, query, price.Price, price.Count, id)
	return err
}

func (r *repository) DeletePrice(ctx context.Context, id int64) error {
	query := `DELETE FROM prices WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *repository) DeletePricesByProductID(ctx context.Context, productID int64) error {
	query := `DELETE FROM prices WHERE product_id = $1`
	_, err := r.db.ExecContext(ctx, query, productID)
	return err
}
func (r *repository) UpdatePriceCount(ctx context.Context, id int64, newCount int) error {
	query := `
		UPDATE prices
		SET count = $1
		WHERE id = $2`
	_, err := r.db.ExecContext(ctx, query, newCount, id)
	return err
}
