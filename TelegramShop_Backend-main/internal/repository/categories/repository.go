package categories

import (
	"context"
	"database/sql"

	"telegramshop_backend/internal/models"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	CreateCategory(ctx context.Context, category models.Category) (models.Category, error)
	GetCategoryByID(ctx context.Context, id int64) (models.Category, error)
	GetAllCategories(ctx context.Context) ([]models.Category, error)
	UpdateCategory(ctx context.Context, id int64, category models.UpdateCategoryInput) error
	DeleteCategory(ctx context.Context, id int64) error
	SetImage(ctx context.Context, id int64, imageURL string) error
	RemoveImage(ctx context.Context, id int64) error
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{db: db}
}

func (r *repository) CreateCategory(ctx context.Context, category models.Category) (models.Category, error) {
	query := `INSERT INTO categories (name) VALUES ($1) RETURNING id`
	err := r.db.QueryRowContext(ctx, query, category.Name).Scan(&category.ID)
	return category, err
}

func (r *repository) GetCategoryByID(ctx context.Context, id int64) (models.Category, error) {
	query := `SELECT id, name, image FROM categories WHERE id = $1`

	var category models.Category
	err := r.db.GetContext(ctx, &category, query, id)
	if err == sql.ErrNoRows {
		return models.Category{}, nil
	}
	return category, err
}

func (r *repository) GetAllCategories(ctx context.Context) ([]models.Category, error) {
	query := `SELECT id, name, image FROM categories`

	var categories []models.Category
	err := r.db.SelectContext(ctx, &categories, query)
	return categories, err
}

func (r *repository) UpdateCategory(ctx context.Context, id int64, category models.UpdateCategoryInput) error {
	query := `UPDATE categories SET name = $1, image = $2 WHERE id = $3`
	_, err := r.db.ExecContext(ctx, query, category.Name, category.Image, id)
	return err
}

func (r *repository) DeleteCategory(ctx context.Context, id int64) error {
	query := `DELETE FROM categories WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *repository) SetImage(ctx context.Context, id int64, imageURL string) error {
	query := `UPDATE categories SET image = $1 WHERE id = $2`
	_, err := r.db.ExecContext(ctx, query, imageURL, id)
	return err
}

func (r *repository) RemoveImage(ctx context.Context, id int64) error {
	query := `UPDATE categories SET image = NULL WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
