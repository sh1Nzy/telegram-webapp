package basket

import (
	"context"
	"telegramshop_backend/internal/models"
	"telegramshop_backend/pkg/logger"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Repository interface {
	UpdateBasketItemQuantity(ctx context.Context, itemID int, quantity int) error
	RemoveFromBasket(ctx context.Context, itemID int) error
	GetUserBasket(ctx context.Context, userID int64) ([]models.BasketItem, error)
	ClearUserBasket(ctx context.Context, userID int64) error
	CreateBasketItem(ctx context.Context, input models.CreateBasketItem) error
	DeleteBasketItem(ctx context.Context, input models.DeleteBasketItem) error
	UpdateBasketItem(ctx context.Context, input models.CreateBasketItem) error
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{db: db}
}

func (r *repository) UpdateBasketItemQuantity(ctx context.Context, itemID int, quantity int) error {

	query := `UPDATE basket SET quantity = $1 WHERE id = $2`

	_, err := r.db.ExecContext(ctx, query, quantity, itemID)

	return err
}

func (r *repository) RemoveFromBasket(ctx context.Context, itemID int) error {

	query := `DELETE FROM basket WHERE id = $1`

	_, err := r.db.ExecContext(ctx, query, itemID)

	return err
}

func (r *repository) GetUserBasket(ctx context.Context, userID int64) ([]models.BasketItem, error) {
	query := `
		SELECT user_id, product_id, quantity
		FROM basket
		WHERE user_id = $1
	`

	var items []models.BasketItem
	err := r.db.SelectContext(ctx, &items, query, userID)
	if err != nil {
		logger.Errorf("[GetUserBasket] Error getting basket items: %v", err)
		return nil, err
	}

	return items, nil
}

func (r *repository) ClearUserBasket(ctx context.Context, userID int64) error {
	query := `DELETE FROM basket WHERE user_id = $1`

	_, err := r.db.ExecContext(ctx, query, userID)

	return err
}

func (r *repository) CreateBasketItem(ctx context.Context, input models.CreateBasketItem) error {
	query := `
		INSERT INTO basket (user_id, product_id, quantity)
		VALUES ($1, $2, $3)
	`

	_, err := r.db.ExecContext(ctx, query, input.UserID, input.ProductID, input.Quantity)
	if err != nil {
		logger.Errorf("[CreateBasketItem] Error creating basket item: %v", err)
		return err
	}

	return nil
}

func (r *repository) DeleteBasketItem(ctx context.Context, input models.DeleteBasketItem) error {

	query := `
		DELETE FROM basket
		WHERE user_id = $1 AND product_id = $2
	`

	_, err := r.db.ExecContext(ctx, query, input.UserID, input.ProductID)

	return err
}

func (r *repository) UpdateBasketItem(ctx context.Context, input models.CreateBasketItem) error {
	query := `
		UPDATE basket
		SET quantity = $3
		WHERE user_id = $1 AND product_id = $2
	`

	_, err := r.db.ExecContext(ctx, query, input.UserID, input.ProductID, input.Quantity)
	if err != nil {
		logger.Errorf("[UpdateBasketItem] Error updating basket item: %v", err)
		return err
	}

	return nil
}
