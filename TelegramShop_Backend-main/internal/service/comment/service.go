package comment

import (
	"context"
	"errors"
	"telegramshop_backend/internal/models"
	"telegramshop_backend/internal/repository/comment"
	"telegramshop_backend/pkg/logger"
)

type CommentService interface {
	AddComment(ctx context.Context, id int64, productID int, commentText string) (models.Comment, error)
	EditComment(ctx context.Context, id int64, productID int, newCommentText string) error
	DeleteComment(ctx context.Context, id int64, productID int) error
	GetCommentsByProduct(ctx context.Context, productID int) ([]models.Comment, error)
}

type service struct {
	repo comment.Repository
}

func NewService(repo comment.Repository) CommentService {
	return &service{repo: repo}
}

func (s *service) AddComment(ctx context.Context, id int64, productID int, commentText string) (models.Comment, error) {

	logger.Infof("[AddComment] Create comment with id=%d, productID=%d, text=%s", id, productID, commentText)
	comment := models.Comment{
		UserID:    id,
		ProductID: productID,
		Comment:   commentText,
	}
	return s.repo.AddComment(ctx, comment)
}

func (s *service) EditComment(ctx context.Context, id int64, productID int, newCommentText string) error {

	logger.Infof("[EditComment]: id: %d, productID: %d", id, productID)

	comments, err := s.repo.GetCommentsByUser(ctx, id)
	if err != nil {
		logger.Errorf("[EditComment] Error editing comment: %v", err)
		return err
	}

	commentID := -1
	for _, comment := range comments {
		if comment.ProductID == productID {
			commentID = comment.ID
			break
		}
	}

	if commentID == -1 {
		return errors.New("comment not found for this user and product")
	}

	return s.repo.UpdateComment(ctx, commentID, newCommentText)
}

func (s *service) DeleteComment(ctx context.Context, id int64, productID int) error {

	logger.Infof("[DeleteComment]: id: %d, productID: %d", id, productID)

	comments, err := s.repo.GetCommentsByUser(ctx, id)
	if err != nil {
		logger.Errorf("[DeleteComment] Error deleting comment: %v", err)
		return err
	}

	commentID := -1
	for _, comment := range comments {
		if comment.ProductID == productID {
			commentID = comment.ID
			break
		}
	}

	if commentID == -1 {
		return errors.New("comment not found for this user and product")
	}

	return s.repo.DeleteComment(ctx, commentID)
}

func (s *service) GetCommentsByProduct(ctx context.Context, productID int) ([]models.Comment, error) {
	return s.repo.GetCommentsByProduct(ctx, productID)
}
