package handlers

import (
	"pade-backend/api/presenter"
	"pade-backend/pkg/auth"
	"pade-backend/pkg/entities"
	"pade-backend/pkg/user"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService user.Service
}

type authRequest struct {
	Name            string           `json:"name"`
	Email           string           `json:"email"`
	Password        string           `json:"password"`
	Phone           *string          `json:"phone"`
	ProfileImageURL *string          `json:"profileImageUrl"`
	Roles           []entities.Roles `json:"roles"`
}

type authResponse struct {
	User  *entities.User `json:"user"`
	Token string         `json:"token"`
}

func NewUserHandler(userService user.Service) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// Register godoc
// @Summary Register a new user
// @Description Create a new user account
// @Accept json
// @Produce json
// @Param user body authRequest true "User registration data"
// @Router /auth/register [post]
func (h *UserHandler) Register(c *fiber.Ctx) error {
	var request authRequest
	if err := c.BodyParser(&request); err != nil {
		return presenter.Error(c, fiber.StatusBadRequest, "Invalid request body")
	}

	result, err := h.userService.Register(&entities.User{
		Name:            request.Name,
		Email:           request.Email,
		Password:        request.Password,
		Phone:           request.Phone,
		ProfileImageURL: request.ProfileImageURL,
		Roles:           request.Roles,
	})
	if err != nil {
		return presenter.Error(c, fiber.StatusBadRequest, err.Error())
	}

	token, err := auth.GenerateToken(result.ID, result.Email, rolesToStrings(result.Roles))
	if err != nil {
		return presenter.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return presenter.Success(c, fiber.StatusCreated, "User registered successfully", authResponse{
		User:  result,
		Token: token,
	})
}

// Login godoc
// @Summary Login a user
// @Description Authenticate a user and return a JWT token
// @Accept json
// @Produce json
// @Param credentials body authRequest true "Login credentials"
// @Router /auth/login [post]
func (h *UserHandler) Login(c *fiber.Ctx) error {
	var request authRequest
	if err := c.BodyParser(&request); err != nil {
		return presenter.Error(c, fiber.StatusBadRequest, "Invalid request body")
	}

	result, err := h.userService.Login(request.Email, request.Password)
	if err != nil {
		return presenter.Error(c, fiber.StatusUnauthorized, err.Error())
	}

	token, err := auth.GenerateToken(result.ID, result.Email, rolesToStrings(result.Roles))
	if err != nil {
		return presenter.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return presenter.Success(c, fiber.StatusOK, "Login successful", authResponse{
		User:  result,
		Token: token,
	})
}

// GetUserByID godoc
// @Summary Get user by ID
// @Description Retrieve user information by ID
// @Produce json
// @Param id path string true "User ID"
// @Router /users/{id} [get]
func (h *UserHandler) GetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return presenter.Error(c, fiber.StatusBadRequest, "User ID is required")
	}

	user, err := h.userService.GetUserByID(id)
	if err != nil {
		return presenter.Error(c, fiber.StatusNotFound, "User not found")
	}

	return presenter.Success(c, fiber.StatusOK, "User retrieved successfully", user)
}

// UpdateUser godoc
// @Summary Update user information
// @Description Update existing user data
// @Accept json
// @Produce json
// @Param user body entities.User true "Updated user data"
// @Router /users/{id} [put]
func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	var user entities.User
	if err := c.BodyParser(&user); err != nil {
		return presenter.Error(c, fiber.StatusBadRequest, "Invalid request body")
	}

	user.ID = c.Params("id")
	result, err := h.userService.UpdateUser(&user)
	if err != nil {
		return presenter.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return presenter.Success(c, fiber.StatusOK, "User updated successfully", result)
}

// DeleteUser godoc
// @Summary Delete user
// @Description Remove a user account
// @Param id path string true "User ID"
// @Router /users/{id} [delete]
func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return presenter.Error(c, fiber.StatusBadRequest, "User ID is required")
	}

	if err := h.userService.DeleteUser(id); err != nil {
		return presenter.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return presenter.Success(c, fiber.StatusOK, "User deleted successfully", nil)
}

// GetAllUsers godoc
// @Summary Get all users
// @Description Retrieve all users (admin only)
// @Produce json
// @Router /users [get]
func (h *UserHandler) GetAllUsers(c *fiber.Ctx) error {
	users, err := h.userService.GetAllUsers()
	if err != nil {
		return presenter.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return presenter.Success(c, fiber.StatusOK, "Users retrieved successfully", users)
}

func rolesToStrings(roles []entities.Roles) []string {
	result := make([]string, 0, len(roles))
	for _, role := range roles {
		result = append(result, role.String())
	}

	return result
}
