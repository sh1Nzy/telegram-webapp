package orders

import (
	"context"

	"telegramshop_backend/internal/models"
	"telegramshop_backend/internal/repository/orders"
	"telegramshop_backend/pkg/logger"
)

type Service interface {
	GetAll(ctx context.Context) ([]models.OrderWithProducts, error)
	CreateOrder(ctx context.Context, input models.CreateOrder) (models.OrderWithProducts, error)
	GetOrderByID(ctx context.Context, id int) (models.OrderWithProducts, error)
	GetUserOrders(ctx context.Context, userID int64) ([]models.OrderWithProducts, error)
}

type service struct {
	repo orders.Repository
}

func NewService(repo orders.Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetAll(ctx context.Context) ([]models.OrderWithProducts, error) {

	logger.Info("[GetAll] Getting all orders")

	orders, err := s.repo.GetAll(ctx)
	if err != nil {
		logger.Errorf("[GetAll] Error getting orders: %v", err)
		return nil, err
	}

	return orders, nil
}

func (s *service) CreateOrder(ctx context.Context, input models.CreateOrder) (models.OrderWithProducts, error) {

	logger.Infof("[CreateOrder] Creating order for user %d", input.UserID)

	createdOrder, err := s.repo.CreateOrder(ctx, input)
	if err != nil {
		logger.Errorf("[CreateOrder] Error creating order: %v", err)
		return models.OrderWithProducts{}, err
	}

	return createdOrder, nil
}

func (s *service) GetOrderByID(ctx context.Context, id int) (models.OrderWithProducts, error) {
	logger.Infof("[GetOrderByID] Getting order with id=%d", id)

	order, err := s.repo.GetOrderByID(ctx, id)
	if err != nil {
		logger.Errorf("[GetOrderByID] Error getting order: %v", err)
		return models.OrderWithProducts{}, err
	}

	return order, nil
}

func (s *service) GetUserOrders(ctx context.Context, userID int64) ([]models.OrderWithProducts, error) {
	logger.Infof("[GetUserOrders] Getting orders for user %d", userID)

	orders, err := s.repo.GetUserOrders(ctx, userID)
	if err != nil {
		logger.Errorf("[GetUserOrders] Error getting user orders: %v", err)
		return nil, err
	}

	return orders, nil
}
