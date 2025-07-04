package users

import (
	"context"
	"database/sql"
	"time"

	"telegramshop_backend/internal/models"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Repository interface {
	CreateUser(ctx context.Context, user models.CreateUser) (models.User, error)
	GetUserByID(ctx context.Context, telegramID int64) (models.User, error)
	GetUserByUsername(ctx context.Context, username string) (models.User, error)
	UpdateUser(ctx context.Context, user models.User) error
	DeleteUser(ctx context.Context, telegramID int64) error
	GetAll(ctx context.Context) ([]models.User, error)
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{db: db}
}

func (r *repository) CreateUser(ctx context.Context, user models.CreateUser) (models.User, error) {
	query := `
		INSERT INTO users (telegram_id, username, created_at)
		VALUES ($1, $2, $3)
		RETURNING id, telegram_id, username, created_at`

	var u models.User
	err := r.db.QueryRowContext(ctx, query, user.TelegramID, user.Username, time.Now()).
		Scan(&u.ID, &u.TelegramID, &u.Username, &u.CreatedAt)
	if err != nil {
		return models.User{}, err
	}

	return u, nil
}

func (r *repository) GetUserByID(ctx context.Context, telegramID int64) (models.User, error) {
	query := `
		SELECT id, telegram_id, username, created_at
		FROM users
		WHERE telegram_id = $1
	`
	var user models.User
	err := r.db.QueryRowContext(ctx, query, telegramID).Scan(&user.ID, &user.TelegramID, &user.Username, &user.CreatedAt)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *repository) GetUserByUsername(ctx context.Context, username string) (models.User, error) {
	query := `SELECT id, telegram_id, username, created_at FROM users WHERE username = $1`

	var user models.User
	err := r.db.GetContext(ctx, &user, query, username)
	if err == sql.ErrNoRows {
		return models.User{}, nil
	}
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *repository) UpdateUser(ctx context.Context, user models.User) error {
	query := `
		UPDATE users 
		SET username = $1
		WHERE telegram_id = $2`

	_, err := r.db.ExecContext(ctx, query, user.Username, user.TelegramID)
	return err
}

func (r *repository) DeleteUser(ctx context.Context, telegramID int64) error {
	query := `
		DELETE FROM users
		WHERE telegram_id = $1
	`
	_, err := r.db.ExecContext(ctx, query, telegramID)
	return err
}

func (r *repository) GetAll(ctx context.Context) ([]models.User, error) {
	query := `
		SELECT id, telegram_id, username, created_at
		FROM users
	`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.TelegramID, &user.Username, &user.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
