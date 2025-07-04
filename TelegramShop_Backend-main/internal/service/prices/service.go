package prices

import (
	"context"
	"database/sql"

	"telegramshop_backend/internal/models"
	"telegramshop_backend/internal/repository/prices"
	"telegramshop_backend/pkg/logger"
)

type Service interface {
	CreatePrice(ctx context.Context, input models.Price) (models.Price, error)
	GetPriceByID(ctx context.Context, id int64) (models.Price, error)
	GetPricesByProductID(ctx context.Context, productID int64) ([]models.Price, error)
	UpdatePrice(ctx context.Context, id int64, input models.UpdatePriceInput) error
	DeletePrice(ctx context.Context, id int64) error
	DeletePricesByProductID(ctx context.Context, productID int64) error
	UpdatePriceCount(ctx context.Context, id int64, newCount int) error
}

type service struct {
	repo prices.Repository
}

func NewService(repo prices.Repository) Service {
	return &service{repo: repo}
}

func (s *service) CreatePrice(ctx context.Context, input models.Price) (models.Price, error) {
	logger.Infof("[CreatePrice] Creating price for product_id=%d with count=%d and price=%f",
		input.ProductID, input.Count, input.Price)

	price, err := s.repo.CreatePrice(ctx, input)
	if err != nil {
		logger.Errorf("[CreatePrice] Error creating price: %v", err)
		return models.Price{}, err
	}

	return price, nil
}

func (s *service) GetPriceByID(ctx context.Context, id int64) (models.Price, error) {
	logger.Infof("[GetPriceByID] Getting price with id=%d", id)

	price, err := s.repo.GetPriceByID(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Infof("[GetPriceByID] No prices found with id=%d", id)
		}
		logger.Errorf("[GetPriceByID] Error getting price: %v", err)
		return models.Price{}, err
	}

	return price, nil
}

func (s *service) GetPricesByProductID(ctx context.Context, productID int64) ([]models.Price, error) {
	logger.Infof("[GetPricesByProductID] Getting prices for product_id=%d", productID)

	prices, err := s.repo.GetPricesByProductID(ctx, productID)
	if err != nil {
		logger.Errorf("[GetPricesByProductID] Error getting prices: %v", err)
		return nil, err
	}

	return prices, nil
}

func (s *service) UpdatePrice(ctx context.Context, id int64, input models.UpdatePriceInput) error {
	logger.Infof("[UpdatePrice] Updating price with id=%d", id)

	err := s.repo.UpdatePrice(ctx, id, input)
	if err != nil {
		logger.Errorf("[UpdatePrice] Error updating price: %v", err)
		return err
	}

	return nil
}

func (s *service) DeletePrice(ctx context.Context, id int64) error {
	logger.Infof("[DeletePrice] Deleting price with id=%d", id)

	err := s.repo.DeletePrice(ctx, id)
	if err != nil {
		logger.Errorf("[DeletePrice] Error deleting price: %v", err)
		return err
	}

	return nil
}

func (s *service) DeletePricesByProductID(ctx context.Context, productID int64) error {
	logger.Infof("[DeletePricesByProductID] Deleting all prices for product_id=%d", productID)

	err := s.repo.DeletePricesByProductID(ctx, productID)
	if err != nil {
		logger.Errorf("[DeletePricesByProductID] Error deleting prices: %v", err)
		return err
	}

	return nil
}

func (s *service) UpdatePriceCount(ctx context.Context, id int64, newCount int) error {
	logger.Infof("[UpdatePriceCount] Updating count to %d for price with id=%d", newCount, id)

	err := s.repo.UpdatePriceCount(ctx, id, newCount)
	if err != nil {
		logger.Errorf("[UpdatePriceCount] Error updating price count: %v", err)
		return err
	}

	return nil
}
