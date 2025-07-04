package firms

import (
	"context"
	"database/sql"

	"telegramshop_backend/internal/models"
	"telegramshop_backend/internal/repository/firms"
	"telegramshop_backend/pkg/logger"
)

type Service interface {
	CreateFirm(ctx context.Context, input models.Firm) (models.Firm, error)
	GetFirmByID(ctx context.Context, id int64) (models.Firm, error)
	GetAllFirms(ctx context.Context) ([]models.Firm, error)
	UpdateFirm(ctx context.Context, id int64, input models.UpdateFirmInput) error
	DeleteFirm(ctx context.Context, id int64) error
}

type service struct {
	repo firms.Repository
}

func NewService(repo firms.Repository) Service {
	return &service{repo: repo}
}

func (s *service) CreateFirm(ctx context.Context, input models.Firm) (models.Firm, error) {
	logger.Infof("[CreateFirm] Creating firm with name=%s", input.Name)

	firm, err := s.repo.CreateFirm(ctx, input)
	if err != nil {
		logger.Errorf("[CreateFirm] Error creating firm: %v", err)
		return models.Firm{}, err
	}

	return firm, nil
}

func (s *service) GetFirmByID(ctx context.Context, id int64) (models.Firm, error) {
	logger.Infof("[GetFirmByID] Getting firm with id=%d", id)

	firm, err := s.repo.GetFirmByID(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Errorf("[GetFirmByID] Firm not found: id=%d", id)
			return models.Firm{}, err
		}
		logger.Errorf("[GetFirmByID] Error getting firm: %v", err)
		return models.Firm{}, err
	}

	return firm, nil
}

func (s *service) GetAllFirms(ctx context.Context) ([]models.Firm, error) {
	logger.Info("[GetAllFirms] Getting all firms")

	firms, err := s.repo.GetAllFirms(ctx)
	if err != nil {
		logger.Errorf("[GetAllFirms] Error getting firms: %v", err)
		return nil, err
	}

	return firms, nil
}

func (s *service) UpdateFirm(ctx context.Context, id int64, input models.UpdateFirmInput) error {
	logger.Infof("[UpdateFirm] Updating firm with id=%d", id)

	err := s.repo.UpdateFirm(ctx, id, input)
	if err != nil {
		logger.Errorf("[UpdateFirm] Error updating firm: %v", err)
		return err
	}

	return nil
}

func (s *service) DeleteFirm(ctx context.Context, id int64) error {
	logger.Infof("[DeleteFirm] Deleting firm with id=%d", id)

	err := s.repo.DeleteFirm(ctx, id)
	if err != nil {
		logger.Errorf("[DeleteFirm] Error deleting firm: %v", err)
		return err
	}

	return nil
}
