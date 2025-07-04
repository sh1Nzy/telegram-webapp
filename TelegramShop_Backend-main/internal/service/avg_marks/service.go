package avg_marks

import (
	"context"
	"telegramshop_backend/internal/models"
	"telegramshop_backend/internal/repository/marks"
	"telegramshop_backend/pkg/logger"
)

type AvgMarksService interface {
	GetAvgMark(ctx context.Context, productID int) (models.AvgMarks, error)
	GetAllAvgMarks(ctx context.Context) ([]models.AvgMarks, error)
	RecalculateAvgMark(ctx context.Context, productID int) error
}

type service struct {
	repo marks.Repository
}

func (s *service) GetAvgMark(ctx context.Context, productID int) (models.AvgMarks, error) {
	logger.Infof("[GetAvgMark] Getting average mark for productID=%d", productID)
	avgMark, err := s.repo.GetAvgMarksByProduct(ctx, productID)
	if err != nil {
		logger.Errorf("[GetAvgMark] Error getting average mark: %v", err)
		return models.AvgMarks{}, err
	}
	return avgMark, nil
}

func (s *service) GetAllAvgMarks(ctx context.Context) ([]models.AvgMarks, error) {
	logger.Info("[GetAllAvgMarks] Getting all average marks")
	avgMarks, err := s.repo.GetAllAvgMarks(ctx)
	if err != nil {
		logger.Errorf("[GetAllAvgMarks] Error getting all average marks: %v", err)
		return nil, err
	}
	return avgMarks, nil
}

func (s *service) RecalculateAvgMark(ctx context.Context, productID int) error {
	logger.Infof("[RecalculateAvgMark] Recalculating average mark for productID=%d", productID)

	marks, err := s.repo.GetMarksByProduct(ctx, productID)
	if err != nil {
		logger.Errorf("[RecalculateAvgMark] Error getting marks: %v", err)
		return err
	}

	var sum float64
	count := len(marks)

	if count == 0 {
		logger.Infof("[RecalculateAvgMark] No marks found for productID=%d", productID)
		return nil
	}

	for _, mark := range marks {
		sum += mark.Mark
	}

	avgMark := models.AvgMarks{
		ProductID: productID,
		Sum:       sum,
		Count:     count,
	}

	err = s.repo.UpdateAvgMark(ctx, avgMark)
	if err != nil {
		logger.Errorf("[RecalculateAvgMark] Error updating average mark: %v", err)
		return err
	}

	logger.Infof("[RecalculateAvgMark] Successfully recalculated average mark for productID=%d: sum=%.2f, count=%d",
		productID, sum, count)
	return nil
}

func NewService(repo marks.Repository) AvgMarksService {
	return &service{repo: repo}
}
