package favorites

import (
	"context"

	"telegramshop_backend/internal/models"
	"telegramshop_backend/internal/repository/favorites"
	"telegramshop_backend/pkg/logger"
)

type Service interface {
	GetUserFavorites(ctx context.Context, userID int64) ([]int64, error)
	AddToFavorites(ctx context.Context, input models.Favorite) (models.Favorite, error)
	RemoveFromFavorites(ctx context.Context, userID int64, productID int) error
}

type service struct {
	repo favorites.Repository
}

func NewService(repo favorites.Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetUserFavorites(ctx context.Context, userID int64) ([]int64, error) {
	logger.Infof("[GetUserFavorites] Getting favorite products for user with id=%d", userID)

	favs, err := s.repo.GetUserFavorites(ctx, userID)
	if err != nil {
		logger.Errorf("[GetUserFavorites] Error getting favorite products: %v", err)
		return nil, err
	}

	productIDs := make([]int64, len(favs))
	for i, fav := range favs {
		productIDs[i] = int64(fav.ProductID)
	}
	return productIDs, nil
}

func (s *service) AddToFavorites(ctx context.Context, input models.Favorite) (models.Favorite, error) {
	logger.Infof("[AddToFavorites] Adding product %d to favorites for user %d", input.ProductID, input.UserID)

	err := s.repo.CreateFavorite(ctx, models.CreateFavorite{
		UserID:    input.UserID,
		ProductID: input.ProductID,
	})
	if err != nil {
		logger.Errorf("[AddToFavorites] Error adding to favorites: %v", err)
		return models.Favorite{}, err
	}

	return input, nil
}

func (s *service) RemoveFromFavorites(ctx context.Context, userID int64, productID int) error {
	logger.Infof("[RemoveFromFavorites] Removing product %d from favorites for user %d", productID, userID)

	err := s.repo.DeleteFavorite(ctx, models.DeleteFavorite{
		UserID:    userID,
		ProductID: productID,
	})
	if err != nil {
		logger.Errorf("[RemoveFromFavorites] Error removing from favorites: %v", err)
		return err
	}

	return nil
}
