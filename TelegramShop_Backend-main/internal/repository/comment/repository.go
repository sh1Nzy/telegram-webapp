package comment

import (
	"context"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"telegramshop_backend/internal/models"
	"time"
)

type Repository interface {
	AddComment(ctx context.Context, comment models.Comment) (models.Comment, error)
	UpdateComment(ctx context.Context, commentID int, newComment string) error
	DeleteComment(ctx context.Context, commentID int) error
	GetCommentsByProduct(ctx context.Context, productID int) ([]models.Comment, error)
	GetCommentsByUser(ctx context.Context, userID int64) ([]models.Comment, error)
}

type repository struct {
	db *sqlx.DB
}

func (r repository) AddComment(ctx context.Context, comment models.Comment) (models.Comment, error) {
	query := `
        INSERT INTO comments (user_id, product_id, comment, created_at)
        VALUES ($1, $2, $3, $4)
        RETURNING id, created_at`

	err := r.db.QueryRowContext(ctx, query,
		comment.UserID,
		comment.ProductID,
		comment.Comment,
		time.Now().Format(time.RFC3339),
	).Scan(&comment.ID, &comment.CreatedAt)
	if err != nil {
		return models.Comment{}, err
	}

	return comment, nil
}

func (r *repository) UpdateComment(ctx context.Context, commentID int, newComment string) error {
	query := `UPDATE comments SET comment = $1 WHERE id = $2`
	_, err := r.db.ExecContext(ctx, query, newComment, commentID)
	return err
}

func (r repository) DeleteComment(ctx context.Context, commentID int) error {
	query := `DELETE FROM comments WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, commentID)
	return err
}

func (r repository) GetCommentsByProduct(ctx context.Context, productID int) ([]models.Comment, error) {
	query := `SELECT id, user_id, product_id, comment, created_at FROM comments WHERE product_id = $1`

	var comments []models.Comment
	err := r.db.SelectContext(ctx, &comments, query, productID)
	if err != nil {
		return nil, err
	}

	return comments, nil
}

func (r repository) GetCommentsByUser(ctx context.Context, userID int64) ([]models.Comment, error) {
	query := `SELECT id, user_id, product_id, comment, created_at FROM comments WHERE user_id = $1`

	var comments []models.Comment
	err := r.db.SelectContext(ctx, &comments, query, userID)
	if err != nil {
		return nil, err
	}

	return comments, nil
}

//DeleteCommentsByUser?

func NewRepository(db *sqlx.DB) Repository {
	return &repository{db: db}
}
