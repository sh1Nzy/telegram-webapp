package favorites

import (
	"context"

	"telegramshop_backend/internal/models"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Repository interface {
	GetUserFavorites(ctx context.Context, userID int64) ([]models.Favorite, error)
	IsProductInFavorites(ctx context.Context, userID int64, productID int) (bool, error)
	CreateFavorite(ctx context.Context, input models.CreateFavorite) error
	DeleteFavorite(ctx context.Context, input models.DeleteFavorite) error
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{db: db}
}

func (r *repository) GetUserFavorites(ctx context.Context, userID int64) ([]models.Favorite, error) {
	
	query := `
		SELECT user_id, product_id
		FROM favorites
		WHERE user_id = $1
	`
	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var favorites []models.Favorite
	for rows.Next() {
		var fav models.Favorite
		if err := rows.Scan(&fav.UserID, &fav.ProductID); err != nil {
			return nil, err
		}
		favorites = append(favorites, fav)
	}
	
	return favorites, nil
}

func (r *repository) IsProductInFavorites(ctx context.Context, userID int64, productID int) (bool, error) {
	
	query := `SELECT EXISTS(SELECT 1 FROM favorites WHERE user_id = $1 AND product_id = $2)`

	var exists bool
	err := r.db.GetContext(ctx, &exists, query, userID, productID)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (r *repository) CreateFavorite(ctx context.Context, input models.CreateFavorite) error {
	
	query := `
		INSERT INTO favorites (user_id, product_id)
		VALUES ($1, $2)
	`
	
	_, err := r.db.ExecContext(ctx, query, input.UserID, input.ProductID)
	
	return err
}

func (r *repository) DeleteFavorite(ctx context.Context, input models.DeleteFavorite) error {
	
	query := `
		DELETE FROM favorites
		WHERE user_id = $1 AND product_id = $2
	`
	
	_, err := r.db.ExecContext(ctx, query, input.UserID, input.ProductID)
	
	return err
}
