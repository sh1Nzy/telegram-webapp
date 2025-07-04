package marks

import (
	"context"
	"telegramshop_backend/internal/models"
	"telegramshop_backend/internal/repository/marks"
	"telegramshop_backend/pkg/logger"
)

type MarksService interface {
	GetUserMarks(ctx context.Context, id int64) ([]models.Marks, error)
	GetProductUserMark(ctx context.Context, id int64, productID int) (models.Marks, error)
	AddMark(ctx context.Context, id int64, productID int, markValue float64) (models.Marks, error)
	DeleteMark(ctx context.Context, id int64, productID int) error
}

type service struct {
	repo marks.Repository
}

func (s *service) GetUserMarks(ctx context.Context, id int64) ([]models.Marks, error) {
	return s.repo.GetMarksByUser(ctx, id)
}

func (s *service) GetProductUserMark(ctx context.Context, id int64, productID int) (models.Marks, error) {

	logger.Infof("[GetProductUserMark] Get mark with id=%d, productID=%d", id, productID)

	marks, err := s.repo.GetMarksByUser(ctx, id)
	if err != nil {
		logger.Errorf("[GetProductUserMark] Error getting mark: %s", err)
		return models.Marks{}, err
	}

	for _, mark := range marks {
		if mark.ProductID == productID {
			return mark, nil
		}
	}

	return models.Marks{}, nil
}

func (s *service) AddMark(ctx context.Context, id int64, productID int, markValue float64) (models.Marks, error) {

	mark := models.Marks{
		UserID:    id,
		ProductID: productID,
		Mark:      markValue,
	}
	return s.repo.AddMark(ctx, mark)
}

func (s *service) DeleteMark(ctx context.Context, id int64, productID int) error {
	return s.repo.DeleteMark(ctx, id, productID)
}

func NewService(repo marks.Repository) MarksService {
	return &service{repo: repo}
}
