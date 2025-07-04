package handler

import (
	"strconv"

	"telegramshop_backend/internal/models"
	"telegramshop_backend/pkg/web"

	"github.com/gofiber/fiber/v2"
)

// CreateCategory creates a new category
// @Summary Create new category
// @Description Creates a new category with specified details
// @Tags categories
// @Accept json
// @Produce json
// @Param category body models.Category true "Category creation data"
// @Success 200 {object} models.CategoryResponse "Category successfully created"
// @Failure 400 {object} models.ErrorResponse "Invalid request body"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /api/v1/categories [post]
func (h *Handler) CreateCategory(c *fiber.Ctx) error {
	var input models.Category
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.ErrorResp("error_invalid_request_body", "Invalid request body"))
	}

	category, err := h.categoryService.CreateCategory(c.Context(), input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(web.ErrorResp("error_create_category", err.Error()))
	}

	return c.JSON(web.OkResp("success_category_created", category))
}

// GetCategoryByID retrieves category by ID
// @Summary Get category by ID
// @Description Returns category details by its ID
// @Tags categories
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} models.CategoryResponse "Category retrieved successfully"
// @Failure 400 {object} models.ErrorResponse "Invalid category ID"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /api/v1/categories/{id} [get]
func (h *Handler) GetCategoryByID(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.ErrorResp("error_invalid_id", "Invalid category ID"))
	}

	category, err := h.categoryService.GetCategoryByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(web.ErrorResp("error_get_category", err.Error()))
	}

	return c.JSON(web.OkResp("success_category_retrieved", category))
}

// GetAllCategories retrieves all categories
// @Summary Get all categories
// @Description Returns all categories in the system
// @Tags categories
// @Produce json
// @Success 200 {object} models.CategoryListResponse "All categories retrieved successfully"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /api/v1/categories [get]
func (h *Handler) GetAllCategories(c *fiber.Ctx) error {
	categories, err := h.categoryService.GetAllCategories(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(web.ErrorResp("error_get_all_categories", err.Error()))
	}

	return c.JSON(web.OkResp("success_categories_retrieved", categories))
}

// UpdateCategory updates category details
// @Summary Update category
// @Description Updates category details by its ID
// @Tags categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Param category body models.UpdateCategoryInput true "Updated category data"
// @Success 200 {object} models.SuccessResponse "Category successfully updated"
// @Failure 400 {object} models.ErrorResponse "Invalid request body"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /api/v1/categories/{id} [put]
func (h *Handler) UpdateCategory(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.ErrorResp("error_invalid_id", "Invalid category ID"))
	}
	var input models.UpdateCategoryInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.ErrorResp("error_invalid_request_body", "Invalid request body"))
	}

	if err := h.categoryService.UpdateCategory(c.Context(), id, input); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(web.ErrorResp("error_update_category", err.Error()))
	}

	return c.JSON(web.OkResp("success_category_updated", nil))
}

// DeleteCategory deletes a category
// @Summary Delete category
// @Description Deletes a category by its ID
// @Tags categories
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} models.SuccessResponse "Category successfully deleted"
// @Failure 400 {object} models.ErrorResponse "Invalid category ID"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /api/v1/categories/{id} [delete]
func (h *Handler) DeleteCategory(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.ErrorResp("error_invalid_id", "Invalid category ID"))
	}

	if err := h.categoryService.DeleteCategory(c.Context(), id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(web.ErrorResp("error_delete_category", err.Error()))
	}

	return c.JSON(web.OkResp("success_category_deleted", nil))
}

// SetCategoryImage sets category image
// @Summary Set category image
// @Description Sets or updates the image for a category
// @Tags categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Param image body models.ImageInput true "Image data"
// @Success 200 {object} models.SuccessResponse "Category image successfully set"
// @Failure 400 {object} models.ErrorResponse "Invalid request body"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /api/v1/categories/{id}/image [put]
func (h *Handler) SetCategoryImage(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.ErrorResp("error_invalid_id", "Invalid category ID"))
	}

	var input models.ImageInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.ErrorResp("error_invalid_body", "Invalid body"))
	}

	err = h.categoryService.SetImage(c.Context(), id, input.Image)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(web.ErrorResp("error_set_image", err.Error()))
	}

	return c.JSON(web.OkResp("success_category_image_set", nil))
}

// RemoveCategoryImage removes category image
// @Summary Remove category image
// @Description Removes the image from a category
// @Tags categories
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} models.SuccessResponse "Category image successfully removed"
// @Failure 400 {object} models.ErrorResponse "Invalid category ID"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /api/v1/categories/{id}/image [delete]
func (h *Handler) RemoveCategoryImage(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.ErrorResp("error_invalid_id", "Invalid category ID"))
	}

	err = h.categoryService.RemoveImage(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(web.ErrorResp("error_remove_image", err.Error()))
	}

	return c.JSON(web.OkResp("success_category_image_removed", nil))
}
