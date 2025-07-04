package products

import (
	"context"

	"telegramshop_backend/internal/models"
	"telegramshop_backend/internal/repository/products"
	"telegramshop_backend/pkg/logger"
)

type Service interface {
	CreateProduct(ctx context.Context, input models.Product) (models.Product, error)
	GetProductByID(ctx context.Context, id int64) (models.Product, error)
	GetAllProducts(ctx context.Context) ([]models.Product, error)
	UpdateProduct(ctx context.Context, id int64, input models.UpdateProductInput) error
	DeleteProduct(ctx context.Context, id int64) error
	AddProductImage(ctx context.Context, id int64, imageURL string) error
	RemoveProductImage(ctx context.Context, id int64, imageURL string) error
	SetProductImages(ctx context.Context, id int64, images []string) error
	IncrementSellCount(ctx context.Context, productID int64, count int) error
	UpdateStock(ctx context.Context, productID int64, stock int) error
}

type service struct {
	repo products.Repository
}

func NewService(repo products.Repository) Service {
	return &service{repo: repo}
}

func (s *service) CreateProduct(ctx context.Context, input models.Product) (models.Product, error) {
	logger.Infof("[CreateProduct] Creating product with name=%s", input.Name)

	product, err := s.repo.CreateProduct(ctx, input)
	if err != nil {
		logger.Errorf("[CreateProduct] Error creating product: %v", err)
		return models.Product{}, err
	}

	return product, nil
}

func (s *service) GetProductByID(ctx context.Context, id int64) (models.Product, error) {
	logger.Infof("[GetProductByID] Getting product with id=%d", id)

	product, err := s.repo.GetProductByID(ctx, id)
	if err != nil {
		logger.Errorf("[GetProductByID] Error getting product: %v", err)
		return models.Product{}, err
	}

	return product, nil
}

func (s *service) GetAllProducts(ctx context.Context) ([]models.Product, error) {
	logger.Info("[GetAllProducts] Getting all products")

	products, err := s.repo.GetAllProducts(ctx)
	if err != nil {
		logger.Errorf("[GetAllProducts] Error getting products: %v", err)
		return nil, err
	}

	return products, nil
}

func (s *service) UpdateProduct(ctx context.Context, id int64, input models.UpdateProductInput) error {
	logger.Infof("[UpdateProduct] Updating product with id=%d", id)

	err := s.repo.UpdateProduct(ctx, id, input)
	if err != nil {
		logger.Errorf("[UpdateProduct] Error updating product: %v", err)
		return err
	}

	return nil
}

func (s *service) DeleteProduct(ctx context.Context, id int64) error {
	logger.Infof("[DeleteProduct] Deleting product with id=%d", id)

	err := s.repo.DeleteProduct(ctx, id)
	if err != nil {
		logger.Errorf("[DeleteProduct] Error deleting product: %v", err)
		return err
	}

	return nil
}

func (s *service) AddProductImage(ctx context.Context, id int64, imageURL string) error {
	logger.Infof("[AddProductImage] Adding image %s to product with id=%d", imageURL, id)

	err := s.repo.AddProductImage(ctx, id, imageURL)
	if err != nil {
		logger.Errorf("[AddProductImage] Error adding image: %v", err)
		return err
	}

	return nil
}

func (s *service) RemoveProductImage(ctx context.Context, id int64, imageURL string) error {
	logger.Infof("[RemoveProductImage] Removing image %s from product with id=%d", imageURL, id)

	err := s.repo.RemoveProductImage(ctx, id, imageURL)
	if err != nil {
		logger.Errorf("[RemoveProductImage] Error removing image: %v", err)
		return err
	}

	return nil
}

func (s *service) SetProductImages(ctx context.Context, id int64, images []string) error {
	logger.Infof("[SetProductImages] Setting images for product with id=%d", id)

	err := s.repo.SetProductImages(ctx, id, images)
	if err != nil {
		logger.Errorf("[SetProductImages] Error setting images: %v", err)
		return err
	}

	return nil
}

func (s *service) IncrementSellCount(ctx context.Context, productID int64, count int) error {
	logger.Infof("[IncrementSellCount] Incrementing sell count by %d for product with id=%d", count, productID)

	err := s.repo.IncrementSellCount(ctx, productID, count)
	if err != nil {
		logger.Errorf("[IncrementSellCount] Error incrementing sell count: %v", err)
		return err
	}

	return nil
}

func (s *service) UpdateStock(ctx context.Context, productID int64, stock int) error {
	logger.Infof("[UpdateStock] Updating stock to %d for product with id=%d", stock, productID)

	err := s.repo.UpdateStock(ctx, productID, stock)
	if err != nil {
		logger.Errorf("[UpdateStock] Error updating stock: %v", err)
		return err
	}

	return nil
}
