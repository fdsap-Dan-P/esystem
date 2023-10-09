package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gofiber/fiber"
	"github.com/jackc/pgx/v4"
	"golang.org/x/crypto/bcrypt"

	"github.com/your-username/your-project/domain"
)

type UserHandler struct {
	userRepository domain.UserRepository
}

func NewUserHandler(userRepository domain.UserRepository) *UserHandler {

	return &UserHandler{userRepository: userRepository}
}

func (h *UserHandler) Login(c *fiber.Ctx) error {
	// Parse request body
	var requestBody struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := json.Unmarshal(c.Body(), &requestBody); err != nil {
		return c.Status(http.StatusBadRequest).SendString("Invalid request body")
	}

	// Get the user from the repository
	user, err := h.userRepository.GetByUsername(context.Background(), requestBody.Username)
	if err != nil {
		if err == pgx.ErrNoRows {
			return c.Status(http.StatusUnauthorized).SendString("Invalid username or password")
		}
		return c.Status(http.StatusInternalServerError).SendString("Failed to fetch user")
	}

	// Verify the password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(requestBody.Password)); err != nil {
		return c.Status(http.StatusUnauthorized).SendString("Invalid username or password")
	}

	// Authentication successful, return a token or session cookie
	return c.SendString("Login successful")
}
