package marks

import (
	"context"
	"telegramshop_backend/internal/models"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Repository interface {
	AddMark(ctx context.Context, mark models.Marks) (models.Marks, error)
	UpdateMark(ctx context.Context, userID int64, productID int, newMark float64) error
	DeleteMark(ctx context.Context, userID int64, productID int) error
	GetMarksByProduct(ctx context.Context, productID int) ([]models.Marks, error)
	GetMarksByUser(ctx context.Context, userID int64) ([]models.Marks, error)
	GetAvgMarksByProduct(ctx context.Context, productID int) (models.AvgMarks, error)
	GetAllAvgMarks(ctx context.Context) ([]models.AvgMarks, error)
	UpdateAvgMark(ctx context.Context, avgMark models.AvgMarks) error
}

type repository struct {
	db *sqlx.DB
}

func (r *repository) AddMark(ctx context.Context, mark models.Marks) (models.Marks, error) {
	query := `
        INSERT INTO marks (user_id, product_id, mark, created_at)
        VALUES ($1, $2, $3, $4)
        RETURNING created_at`
	err := r.db.QueryRowContext(ctx, query,
		mark.UserID,
		mark.ProductID,
		mark.Mark,
		time.Now().Format(time.RFC3339),
	).Scan(&mark.CreatedAt)
	if err != nil {
		return models.Marks{}, err
	}

	return mark, nil
}

func (r *repository) UpdateMark(ctx context.Context, userID int64, productID int, newMark float64) error {
	query := `UPDATE marks SET mark = $1 WHERE user_id = $2 AND product_id = $3`
	_, err := r.db.ExecContext(ctx, query, newMark, userID, productID)
	return err
}

func (r *repository) DeleteMark(ctx context.Context, userID int64, productID int) error {
	query := `DELETE FROM marks WHERE user_id = $1 AND product_id = $2`
	_, err := r.db.ExecContext(ctx, query, userID, productID)
	return err
}

func (r *repository) GetMarksByProduct(ctx context.Context, productID int) ([]models.Marks, error) {
	query := `SELECT user_id, product_id, mark, created_at FROM marks WHERE product_id = $1`
	var marks []models.Marks
	err := r.db.SelectContext(ctx, &marks, query, productID)
	if err != nil {
		return nil, err
	}

	return marks, nil
}

func (r *repository) GetMarksByUser(ctx context.Context, userID int64) ([]models.Marks, error) {
	query := `SELECT user_id, product_id, mark, created_at FROM marks WHERE user_id = $1`
	var marks []models.Marks
	err := r.db.SelectContext(ctx, &marks, query, userID)
	if err != nil {
		return nil, err
	}

	return marks, nil
}

func (r *repository) GetAvgMarksByProduct(ctx context.Context, productID int) (models.AvgMarks, error) {
	query := `
        SELECT 
            product_id, 
            COALESCE(SUM(mark), 0) AS sum, 
            COUNT(*) AS count 
        FROM marks 
        WHERE product_id = $1 
        GROUP BY product_id`

	var avgMarks models.AvgMarks
	err := r.db.GetContext(ctx, &avgMarks, query, productID)
	if err != nil {
		return models.AvgMarks{}, err
	}

	return avgMarks, nil
}

func (r *repository) GetAllAvgMarks(ctx context.Context) ([]models.AvgMarks, error) {
	query := `
        SELECT 
            product_id, 
            COALESCE(SUM(mark), 0) AS sum, 
            COUNT(*) AS count 
        FROM marks 
        GROUP BY product_id`

	var avgMarksList []models.AvgMarks
	err := r.db.SelectContext(ctx, &avgMarksList, query)
	if err != nil {
		return nil, err
	}

	return avgMarksList, nil
}

func (r *repository) UpdateAvgMark(ctx context.Context, avgMark models.AvgMarks) error {
	query := `
		INSERT INTO avg_marks (product_id, sum, count)
		VALUES ($1, $2, $3)
		ON CONFLICT (product_id) 
		DO UPDATE SET 
			sum = EXCLUDED.sum,
			count = EXCLUDED.count,
			updated_at = NOW()`

	_, err := r.db.ExecContext(ctx, query,
		avgMark.ProductID,
		avgMark.Sum,
		avgMark.Count,
	)
	return err
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{db: db}
}
