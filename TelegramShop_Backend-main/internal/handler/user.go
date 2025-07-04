package handler

import (
	"strconv"

	"telegramshop_backend/internal/models"
	"telegramshop_backend/pkg/web"

	"github.com/gofiber/fiber/v2"
)

// CreateUser creates a new user
// @Summary Create new user
// @Description Creates a new user in the system
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.CreateUser true "User creation data"
// @Success 200 {object} models.UserResponse "User successfully created"
// @Failure 400 {object} models.ErrorResponse "Invalid request body"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /api/v1/users [post]
func (h *Handler) CreateUser(c *fiber.Ctx) error {
	var input models.CreateUser
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.ErrorResp("error_invalid_request_body", "Invalid request body"))
	}

	user, err := h.userService.CreateUser(c.Context(), input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(web.ErrorResp("error_create_user", err.Error()))
	}

	return c.JSON(web.OkResp("success_user_created", user))
}

// GetUser retrieves user by ID
// @Summary Get user by ID
// @Description Returns user details by ID
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} models.UserResponse "User retrieved successfully"
// @Failure 400 {object} models.ErrorResponse "Invalid user ID"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /api/v1/users/{id} [get]
func (h *Handler) GetUser(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.ErrorResp("error_invalid_user_id", "Invalid user ID"))
	}

	user, err := h.userService.GetUserByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(web.ErrorResp("error_get_user", err.Error()))
	}

	return c.JSON(web.OkResp("success_user_retrieved", user))
}

// DeleteUser deletes user by ID
// @Summary Delete user
// @Description Deletes a user from the system
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} models.SuccessResponse "User successfully deleted"
// @Failure 400 {object} models.ErrorResponse "Invalid user ID"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /api/v1/users/{id} [delete]
func (h *Handler) DeleteUser(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.ErrorResp("error_invalid_user_id", "Invalid user ID"))
	}

	if err := h.userService.DeleteUser(c.Context(), id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(web.ErrorResp("error_delete_user", err.Error()))
	}

	return c.JSON(web.OkResp("success_user_deleted", nil))
}

// GetAllUsers retrieves all users
// @Summary Get all users
// @Description Returns all users in the system
// @Tags users
// @Produce json
// @Success 200 {object} models.UserListResponse "All users retrieved successfully"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /api/v1/users [get]
func (h *Handler) GetAllUsers(c *fiber.Ctx) error {
	users, err := h.userService.GetAll(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(web.ErrorResp("error_get_all_users", err.Error()))
	}
	return c.JSON(web.OkResp("success_all_users_retrieved", users))
}
