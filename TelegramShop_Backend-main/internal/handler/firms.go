package handler

import (
	"database/sql"
	"strconv"
	"telegramshop_backend/internal/models"
	"telegramshop_backend/pkg/web"

	"github.com/gofiber/fiber/v2"
)

// CreateFirm creates a new firm
// @Summary Create new firm
// @Description Creates a new firm with specified details
// @Tags firms
// @Accept json
// @Produce json
// @Param firm body models.Firm true "Firm creation data"
// @Success 200 {object} models.FirmResponse "Firm successfully created"
// @Failure 400 {object} models.ErrorResponse "Invalid request body"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /api/v1/firms [post]
func (h *Handler) CreateFirm(c *fiber.Ctx) error {
	var input models.Firm
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.ErrorResp("error_invalid_request_body", "Invalid request body"))
	}

	firm, err := h.firmsService.CreateFirm(c.Context(), input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(web.ErrorResp("error_create_firm", err.Error()))
	}

	return c.JSON(web.OkResp("success_firm_created", firm))
}

// GetFirmByID retrieves firm by ID
// @Summary Get firm by ID
// @Description Returns firm details by its ID
// @Tags firms
// @Produce json
// @Param id path int true "Firm ID"
// @Success 200 {object} models.FirmResponse "Firm retrieved successfully"
// @Failure 400 {object} models.ErrorResponse "Invalid firm ID"
// @Failure 404 {object} models.ErrorResponse "Firm not found"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /api/v1/firms/{id} [get]
func (h *Handler) GetFirmByID(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.ErrorResp("error_invalid_id", "Invalid firm ID"))
	}

	firm, err := h.firmsService.GetFirmByID(c.Context(), id)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusNotFound).JSON(web.ErrorResp("error_firm_not_found", "Firm not found"))
		}
		return c.Status(fiber.StatusInternalServerError).JSON(web.ErrorResp("error_get_firm", err.Error()))
	}

	return c.JSON(web.OkResp("success_firm_retrieved", firm))
}

// GetAllFirms retrieves all firms
// @Summary Get all firms
// @Description Returns all firms in the system
// @Tags firms
// @Produce json
// @Success 200 {object} models.FirmListResponse "All firms retrieved successfully"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /api/v1/firms [get]
func (h *Handler) GetAllFirms(c *fiber.Ctx) error {
	firms, err := h.firmsService.GetAllFirms(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(web.ErrorResp("error_get_all_firms", err.Error()))
	}

	return c.JSON(web.OkResp("success_all_firms_retrieved", firms))
}

// UpdateFirm updates firm details
// @Summary Update firm
// @Description Updates firm details by its ID
// @Tags firms
// @Accept json
// @Produce json
// @Param id path int true "Firm ID"
// @Param firm body models.UpdateFirmInput true "Updated firm data"
// @Success 200 {object} models.SuccessResponse "Firm successfully updated"
// @Failure 400 {object} models.ErrorResponse "Invalid request body"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /api/v1/firms/{id} [put]
func (h *Handler) UpdateFirm(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.ErrorResp("error_invalid_id", "Invalid firm ID"))
	}
	var input models.UpdateFirmInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.ErrorResp("error_invalid_request_body", "Invalid request body"))
	}

	if err := h.firmsService.UpdateFirm(c.Context(), id, input); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(web.ErrorResp("error_update_firm", err.Error()))
	}

	return c.JSON(web.OkResp("success_firm_updated", nil))
}

// DeleteFirm deletes a firm
// @Summary Delete firm
// @Description Deletes a firm by its ID
// @Tags firms
// @Produce json
// @Param id path int true "Firm ID"
// @Success 200 {object} models.SuccessResponse "Firm successfully deleted"
// @Failure 400 {object} models.ErrorResponse "Invalid firm ID"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /api/v1/firms/{id} [delete]
func (h *Handler) DeleteFirm(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.ErrorResp("error_invalid_id", "Invalid firm ID"))
	}

	if err := h.firmsService.DeleteFirm(c.Context(), id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(web.ErrorResp("error_delete_firm", err.Error()))
	}

	return c.JSON(web.OkResp("success_firm_deleted", nil))
}
