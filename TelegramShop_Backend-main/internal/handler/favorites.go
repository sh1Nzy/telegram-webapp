package handler

import (
	"strconv"

	"telegramshop_backend/internal/models"
	"telegramshop_backend/pkg/web"

	"github.com/gofiber/fiber/v2"
)

// AddToFavorites adds item to user's favorites
// @Summary Add item to favorites
// @Description Adds a product to user's favorites list
// @Tags favorites
// @Accept json
// @Produce json
// @Param favorite body models.Favorite true "Favorite item data"
// @Success 200 {object} models.FavoriteResponse "Item successfully added to favorites"
// @Failure 400 {object} models.ErrorResponse "Invalid request body"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /api/v1/favorites [post]
func (h *Handler) AddToFavorites(c *fiber.Ctx) error {
	var input models.Favorite
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.ErrorResp("error_invalid_request_body", "Invalid request body"))
	}

	favorite, err := h.favoriteService.AddToFavorites(c.Context(), input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(web.ErrorResp("error_add_to_favorites", err.Error()))
	}

	return c.JSON(web.OkResp("success_item_added_to_favorites", favorite))
}

// GetUserFavorites retrieves user's favorites
// @Summary Get user's favorites
// @Description Returns all items in user's favorites list
// @Tags favorites
// @Produce json
// @Param user_id path int true "User ID"
// @Success 200 {object} models.FavoriteListResponse "User's favorites retrieved"
// @Failure 400 {object} models.ErrorResponse "Invalid user ID"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /api/v1/favorites/{user_id} [get]
func (h *Handler) GetUserFavorites(c *fiber.Ctx) error {
	userIDStr := c.Params("user_id")
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.ErrorResp("error_invalid_user_id", "Invalid user ID"))
	}

	favorites, err := h.favoriteService.GetUserFavorites(c.Context(), userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(web.ErrorResp("error_get_user_favorites", err.Error()))
	}

	return c.JSON(web.OkResp("success_user_favorites_retrieved", favorites))
}

// RemoveFromFavorites removes item from user's favorites
// @Summary Remove item from favorites
// @Description Removes product from user's favorites list
// @Tags favorites
// @Produce json
// @Param user_id path int true "User ID"
// @Param product_id path int true "Product ID"
// @Success 200 {object} models.SuccessResponse "Item successfully removed from favorites"
// @Failure 400 {object} models.ErrorResponse "Invalid parameters"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /api/v1/favorites/{user_id}/{product_id} [delete]
func (h *Handler) RemoveFromFavorites(c *fiber.Ctx) error {
	userIDStr := c.Params("user_id")
	productIDStr := c.Params("product_id")

	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.ErrorResp("error_invalid_user_id", "Invalid user ID"))
	}

	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.ErrorResp("error_invalid_product_id", "Invalid product ID"))
	}

	if err := h.favoriteService.RemoveFromFavorites(c.Context(), userID, productID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(web.ErrorResp("error_remove_from_favorites", err.Error()))
	}

	return c.JSON(web.OkResp("success_item_removed_from_favorites", nil))
}
