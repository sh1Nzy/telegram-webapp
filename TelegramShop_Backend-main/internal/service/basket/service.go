package basket

import (
	"context"

	"telegramshop_backend/internal/models"
	"telegramshop_backend/internal/repository/basket"
	"telegramshop_backend/pkg/logger"
)

type Service interface {
	GetUserBasket(ctx context.Context, userID int64) ([]models.BasketItem, error)
	AddToBasket(ctx context.Context, input models.BasketItem) (models.BasketItem, error)
	UpdateBasketItem(ctx context.Context, input models.BasketItem) (models.BasketItem, error)
	RemoveFromBasket(ctx context.Context, userID int64, productID int) error
}

type service struct {
	repo basket.Repository
}

func NewService(repo basket.Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetUserBasket(ctx context.Context, userID int64) ([]models.BasketItem, error) {
	logger.Infof("[GetUserBasket] Getting basket items for user with id=%d", userID)

	items, err := s.repo.GetUserBasket(ctx, userID)
	if err != nil {
		logger.Errorf("[GetUserBasket] Error getting basket items: %v", err)
		return nil, err
	}

	return items, nil
}

func (s *service) AddToBasket(ctx context.Context, input models.BasketItem) (models.BasketItem, error) {
	logger.Infof("[AddToBasket] Adding product %d to basket for user %d", input.ProductID, input.UserID)

	err := s.repo.CreateBasketItem(ctx, models.CreateBasketItem{
		UserID:    input.UserID,
		ProductID: input.ProductID,
		Quantity:  input.Quantity,
	})
	if err != nil {
		logger.Errorf("[AddToBasket] Error adding to basket: %v", err)
		return models.BasketItem{}, err
	}

	return input, nil
}

func (s *service) UpdateBasketItem(ctx context.Context, input models.BasketItem) (models.BasketItem, error) {
	logger.Infof("[UpdateBasketItem] Updating product %d in basket for user %d", input.ProductID, input.UserID)

	err := s.repo.UpdateBasketItem(ctx, models.CreateBasketItem{
		UserID:    input.UserID,
		ProductID: input.ProductID,
		Quantity:  input.Quantity,
	})
	if err != nil {
		logger.Errorf("[UpdateBasketItem] Error updating basket item: %v", err)
		return models.BasketItem{}, err
	}

	return input, nil
}

func (s *service) RemoveFromBasket(ctx context.Context, userID int64, productID int) error {
	logger.Infof("[RemoveFromBasket] Removing product %d from basket for user %d", productID, userID)

	err := s.repo.RemoveFromBasket(ctx, productID)
	if err != nil {
		logger.Errorf("[RemoveFromBasket] Error removing from basket: %v", err)
		return err
	}

	return nil
}
