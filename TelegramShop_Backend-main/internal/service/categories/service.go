package categories

import (
	"context"

	"telegramshop_backend/internal/models"
	"telegramshop_backend/internal/repository/categories"
	"telegramshop_backend/pkg/logger"
)

type Service interface {
	CreateCategory(ctx context.Context, input models.Category) (models.Category, error)
	GetCategoryByID(ctx context.Context, id int64) (models.Category, error)
	GetAllCategories(ctx context.Context) ([]models.Category, error)
	UpdateCategory(ctx context.Context, id int64, input models.UpdateCategoryInput) error
	DeleteCategory(ctx context.Context, id int64) error
	SetImage(ctx context.Context, id int64, imageURL string) error
	RemoveImage(ctx context.Context, id int64) error
}

type service struct {
	repo categories.Repository
}

func NewService(repo categories.Repository) Service {
	return &service{repo: repo}
}

func (s *service) CreateCategory(ctx context.Context, input models.Category) (models.Category, error) {
	logger.Infof("[CreateCategory] Creating category with name=%s", input.Name)

	category, err := s.repo.CreateCategory(ctx, input)
	if err != nil {
		logger.Errorf("[CreateCategory] Error creating category: %v", err)
		return models.Category{}, err
	}

	return category, nil
}

func (s *service) GetCategoryByID(ctx context.Context, id int64) (models.Category, error) {
	logger.Infof("[GetCategoryByID] Getting category with id=%d", id)

	category, err := s.repo.GetCategoryByID(ctx, id)
	if err != nil {
		logger.Errorf("[GetCategoryByID] Error getting category: %v", err)
		return models.Category{}, err
	}

	return category, nil
}

func (s *service) GetAllCategories(ctx context.Context) ([]models.Category, error) {
	logger.Info("[GetAllCategories] Getting all categories")

	categories, err := s.repo.GetAllCategories(ctx)
	if err != nil {
		logger.Errorf("[GetAllCategories] Error getting categories: %v", err)
		return nil, err
	}

	return categories, nil
}

func (s *service) UpdateCategory(ctx context.Context, id int64, input models.UpdateCategoryInput) error {
	logger.Infof("[UpdateCategory] Updating category with id=%d", id)

	err := s.repo.UpdateCategory(ctx, id, input)
	if err != nil {
		logger.Errorf("[UpdateCategory] Error updating category: %v", err)
		return err
	}

	return nil
}

func (s *service) DeleteCategory(ctx context.Context, id int64) error {
	logger.Infof("[DeleteCategory] Deleting category with id=%d", id)

	err := s.repo.DeleteCategory(ctx, id)
	if err != nil {
		logger.Errorf("[DeleteCategory] Error deleting category: %v", err)
		return err
	}

	return nil
}

func (s *service) SetImage(ctx context.Context, id int64, imageURL string) error {
	logger.Infof("[SetImage] Seting image category with id=%d", id)
	err := s.repo.SetImage(ctx, id, imageURL)
	if err != nil {
		logger.Errorf("[SetImage] Error setting image category: %v", err)
		return err
	}
	return nil
}

func (s *service) RemoveImage(ctx context.Context, id int64) error {
	logger.Infof("[RemoveImage] Removing image category with id=%d", id)
	err := s.repo.RemoveImage(ctx, id)
	if err != nil {
		logger.Errorf("[RemoveImage] Error removing image category: %v", err)
		return err
	}
	return nil
}
