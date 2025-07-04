package firms

import (
	"context"
	"database/sql"

	"telegramshop_backend/internal/models"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	CreateFirm(ctx context.Context, firm models.Firm) (models.Firm, error)
	GetFirmByID(ctx context.Context, id int64) (models.Firm, error)
	GetAllFirms(ctx context.Context) ([]models.Firm, error)
	UpdateFirm(ctx context.Context, id int64, input models.UpdateFirmInput) error
	DeleteFirm(ctx context.Context, id int64) error
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{db: db}
}

func (r *repository) CreateFirm(ctx context.Context, firm models.Firm) (models.Firm, error) {
	query := `
		INSERT INTO firms (name)
		VALUES ($1)
		RETURNING id`

	err := r.db.QueryRowContext(ctx, query, firm.Name).Scan(&firm.ID)
	if err != nil {
		return models.Firm{}, err
	}

	return firm, nil
}

func (r *repository) GetFirmByID(ctx context.Context, id int64) (models.Firm, error) {
	query := `
		SELECT id, name
		FROM firms
		WHERE id = $1`

	var firm models.Firm
	err := r.db.GetContext(ctx, &firm, query, id)
	if err == sql.ErrNoRows {
		return models.Firm{}, sql.ErrNoRows
	}
	if err != nil {
		return models.Firm{}, err
	}

	return firm, nil
}

func (r *repository) GetAllFirms(ctx context.Context) ([]models.Firm, error) {
	query := `
		SELECT id, name
		FROM firms`

	var firms []models.Firm
	err := r.db.SelectContext(ctx, &firms, query)
	if err != nil {
		return nil, err
	}

	return firms, nil
}

func (r *repository) UpdateFirm(ctx context.Context, id int64, input models.UpdateFirmInput) error {
	query := `
		UPDATE firms
		SET name = $1
		WHERE id = $2`

	_, err := r.db.ExecContext(ctx, query, input.Name, id)
	return err
}

func (r *repository) DeleteFirm(ctx context.Context, id int64) error {
	query := `DELETE FROM firms WHERE id = $1`

	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
