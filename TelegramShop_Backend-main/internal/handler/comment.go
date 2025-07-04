package handler

import (
	"strconv"
	"telegramshop_backend/pkg/logger"
	"telegramshop_backend/pkg/web"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) AddComment(c *fiber.Ctx) error {
	userID, err := strconv.ParseInt(c.Params("user_id"), 10, 64)
	if err != nil {
		logger.Errorf("[AddComment] Error parsing user_id: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(web.ErrorResp("error", "Invalid user_id"))
	}

	productID, err := strconv.Atoi(c.Params("product_id"))
	if err != nil {
		logger.Errorf("[AddComment] Error parsing product_id: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(web.ErrorResp("error", "Invalid product_id"))
	}

	var request struct {
		Comment string `json:"comment"`
	}

	if err := c.BodyParser(&request); err != nil {
		logger.Errorf("[AddComment] Error parsing request body: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(web.ErrorResp("error", "Invalid request body"))
	}

	comment, err := h.commentService.AddComment(c.Context(), userID, productID, request.Comment)
	if err != nil {
		logger.Errorf("[AddComment] Error adding comment: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(web.ErrorResp("error", "Failed to add comment"))
	}

	return c.JSON(web.OkResp("success_add_comment", comment))
}

func (h *Handler) EditComment(c *fiber.Ctx) error {
	userID, err := strconv.ParseInt(c.Params("user_id"), 10, 64)
	if err != nil {
		logger.Errorf("[EditComment] Error parsing user_id: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(web.ErrorResp("error", "Invalid user_id"))
	}

	productID, err := strconv.Atoi(c.Params("product_id"))
	if err != nil {
		logger.Errorf("[EditComment] Error parsing product_id: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(web.ErrorResp("error", "Invalid product_id"))
	}

	var request struct {
		Comment string `json:"comment"`
	}

	if err := c.BodyParser(&request); err != nil {
		logger.Errorf("[EditComment] Error parsing request body: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(web.ErrorResp("error", "Invalid request body"))
	}

	err = h.commentService.EditComment(c.Context(), userID, productID, request.Comment)
	if err != nil {
		logger.Errorf("[EditComment] Error editing comment: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(web.ErrorResp("error", "Failed to edit comment"))
	}

	return c.JSON(web.OkResp("success_edit_comment", nil))
}

func (h *Handler) DeleteComment(c *fiber.Ctx) error {
	userID, err := strconv.ParseInt(c.Params("user_id"), 10, 64)
	if err != nil {
		logger.Errorf("[DeleteComment] Error parsing user_id: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(web.ErrorResp("error", "Invalid user_id"))
	}

	productID, err := strconv.Atoi(c.Params("product_id"))
	if err != nil {
		logger.Errorf("[DeleteComment] Error parsing product_id: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(web.ErrorResp("error", "Invalid product_id"))
	}

	err = h.commentService.DeleteComment(c.Context(), userID, productID)
	if err != nil {
		logger.Errorf("[DeleteComment] Error deleting comment: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(web.ErrorResp("error", "Failed to delete comment"))
	}

	return c.JSON(web.OkResp("success_delete_comment", nil))
}

func (h *Handler) GetCommentsByProduct(c *fiber.Ctx) error {
	productID, err := strconv.Atoi(c.Params("product_id"))
	if err != nil {
		logger.Errorf("[GetCommentsByProduct] Error parsing product_id: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(web.ErrorResp("error", "Invalid product_id"))
	}

	comments, err := h.commentService.GetCommentsByProduct(c.Context(), productID)
	if err != nil {
		logger.Errorf("[GetCommentsByProduct] Error getting comments: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(web.ErrorResp("error", "Failed to get comments"))
	}

	return c.JSON(web.OkResp("success_get_comments_by_product", comments))
}
