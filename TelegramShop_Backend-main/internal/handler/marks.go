package handler

import (
	"strconv"
	"telegramshop_backend/pkg/logger"
	"telegramshop_backend/pkg/web"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) GetUserMarks(c *fiber.Ctx) error {
	userID, err := strconv.ParseInt(c.Params("user_id"), 10, 64)
	if err != nil {
		logger.Errorf("[GetUserMarks] Error parsing user_id: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(web.ErrorResp("error", "Invalid user_id"))
	}

	marks, err := h.marksService.GetUserMarks(c.Context(), userID)
	if err != nil {
		logger.Errorf("[GetUserMarks] Error getting user marks: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(web.ErrorResp("error_get_marks", "Failed to get user marks"))
	}

	return c.JSON(web.OkResp("success_get_user_marks", marks))
}

func (h *Handler) GetProductUserMark(c *fiber.Ctx) error {
	userID, err := strconv.ParseInt(c.Params("user_id"), 10, 64)
	if err != nil {
		logger.Errorf("[GetProductUserMark] Error parsing user_id: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(web.ErrorResp("error", "Invalid user_id"))
	}

	productID, err := strconv.Atoi(c.Params("product_id"))
	if err != nil {
		logger.Errorf("[GetProductUserMark] Error parsing product_id: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(web.ErrorResp("error", "Invalid product_id"))
	}

	mark, err := h.marksService.GetProductUserMark(c.Context(), userID, productID)
	if err != nil {
		logger.Errorf("[GetProductUserMark] Error getting product user mark: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(web.ErrorResp("error", "Failed to get product user mark"))
	}

	return c.JSON(web.OkResp("success_get_product_user_marks", mark))
}

func (h *Handler) AddMark(c *fiber.Ctx) error {
	userID, err := strconv.ParseInt(c.Params("user_id"), 10, 64)
	if err != nil {
		logger.Errorf("[AddMark] Error parsing user_id: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(web.ErrorResp("error", "Invalid user_id"))
	}

	productID, err := strconv.Atoi(c.Params("product_id"))
	if err != nil {
		logger.Errorf("[AddMark] Error parsing product_id: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(web.ErrorResp("error", "Invalid product_id"))
	}

	var request struct {
		Mark float64 `json:"mark"`
	}

	if err := c.BodyParser(&request); err != nil {
		logger.Errorf("[AddMark] Error parsing request body: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(web.ErrorResp("error", "Invalid request body"))
	}

	mark, err := h.marksService.AddMark(c.Context(), userID, productID, request.Mark)
	if err != nil {
		logger.Errorf("[AddMark] Error adding mark: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(web.ErrorResp("error", "Failed to add mark"))
	}

	if err := h.avgMarksService.RecalculateAvgMark(c.Context(), productID); err != nil {
		logger.Errorf("[AddMark] Error recalculating average mark: %v", err)
	}

	return c.JSON(web.OkResp("success_add_mark", mark))
}

func (h *Handler) DeleteMark(c *fiber.Ctx) error {
	userID, err := strconv.ParseInt(c.Params("user_id"), 10, 64)
	if err != nil {
		logger.Errorf("[DeleteMark] Error parsing user_id: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(web.ErrorResp("error", "Invalid user_id"))
	}

	productID, err := strconv.Atoi(c.Params("product_id"))
	if err != nil {
		logger.Errorf("[DeleteMark] Error parsing product_id: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(web.ErrorResp("error", "Invalid product_id"))
	}

	err = h.marksService.DeleteMark(c.Context(), userID, productID)
	if err != nil {
		logger.Errorf("[DeleteMark] Error deleting mark: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(web.ErrorResp("error", "Failed to delete mark"))
	}

	// Recalculate average mark
	if err := h.avgMarksService.RecalculateAvgMark(c.Context(), productID); err != nil {
		logger.Errorf("[DeleteMark] Error recalculating average mark: %v", err)
	}

	return c.JSON(web.OkResp("success_delete_mark", nil))
}
