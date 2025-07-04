package handler

import (
	"strconv"
	"telegramshop_backend/pkg/logger"
	"telegramshop_backend/pkg/web"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) GetAvgMark(c *fiber.Ctx) error {
	productID, err := strconv.Atoi(c.Params("product_id"))
	if err != nil {
		logger.Errorf("[GetAvgMark] Error parsing product_id: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(web.ErrorResp("error", "Invalid product_id"))
	}

	avgMark, err := h.avgMarksService.GetAvgMark(c.Context(), productID)
	if err != nil {
		logger.Errorf("[GetAvgMark] Error getting average mark: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(web.ErrorResp("error", "Failed to get average mark"))
	}

	return c.JSON(web.OkResp("success_get_avg_mark", avgMark))
}

func (h *Handler) GetAllAvgMarks(c *fiber.Ctx) error {
	avgMarks, err := h.avgMarksService.GetAllAvgMarks(c.Context())
	if err != nil {
		logger.Errorf("[GetAllAvgMarks] Error getting all average marks: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(web.ErrorResp("error", "Failed to get average marks"))
	}

	return c.JSON(web.OkResp("success_get_all_avg_marks", avgMarks))
}

func (h *Handler) RecalculateAvgMark(c *fiber.Ctx) error {
	productID, err := strconv.Atoi(c.Params("product_id"))
	if err != nil {
		logger.Errorf("[RecalculateAvgMark] Error parsing product_id: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(web.ErrorResp("error", "Invalid product_id"))
	}

	err = h.avgMarksService.RecalculateAvgMark(c.Context(), productID)
	if err != nil {
		logger.Errorf("[RecalculateAvgMark] Error recalculating average mark: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(web.ErrorResp("error", "Failed to recalculate average mark"))
	}

	return c.JSON(web.OkResp("success_recalculate", nil))
}
