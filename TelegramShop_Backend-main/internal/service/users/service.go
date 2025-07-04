package users

import (
	"context"

	"telegramshop_backend/internal/models"
	"telegramshop_backend/internal/repository/users"
	"telegramshop_backend/pkg/logger"
)

type Service interface {
	GetUserByID(ctx context.Context, id int64) (models.User, error)
	CreateUser(ctx context.Context, input models.CreateUser) (models.User, error)
	GetAll(ctx context.Context) ([]models.User, error)
	DeleteUser(ctx context.Context, id int64) error
}

type service struct {
	repo users.Repository
}

func NewService(repo users.Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetUserByID(ctx context.Context, id int64) (models.User, error) {

	logger.Infof("[GetUserByID] Getting user with id=%d", id)

	user, err := s.repo.GetUserByID(ctx, id)
	if err != nil {
		logger.Errorf("[GetUserByID] Error getting user: %v", err)
		return models.User{}, err
	}

	return user, nil
}

func (s *service) CreateUser(ctx context.Context, input models.CreateUser) (models.User, error) {

	logger.Infof("[CreateUser] Creating user with id=%d, username=%s", input.TelegramID, input.Username)

	user := models.CreateUser{
		TelegramID: input.TelegramID,
		Username:   input.Username,
	}

	createdUser, err := s.repo.CreateUser(ctx, user)
	if err != nil {
		logger.Errorf("[CreateUser] Error creating user: %v", err)
		return models.User{}, err
	}

	return createdUser, nil
}

func (s *service) GetAll(ctx context.Context) ([]models.User, error) {

	logger.Info("[GetAll] Getting all users")

	users, err := s.repo.GetAll(ctx)
	if err != nil {
		logger.Errorf("[GetAll] Error getting users: %v", err)
		return nil, err
	}

	return users, nil
}

func (s *service) DeleteUser(ctx context.Context, id int64) error {

	logger.Infof("[DeleteUser] Deleting user with id=%d", id)

	err := s.repo.DeleteUser(ctx, id)
	if err != nil {
		logger.Errorf("[DeleteUser] Error deleting user: %v", err)
		return err
	}

	return nil
}