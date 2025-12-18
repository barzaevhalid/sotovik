package user

import (
	"fmt"
	"time"

	"github.com/barzaevhalid/sotovik/utils"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService *UserService
}

func NewUserHandler(service *UserService) *UserHandler {
	return &UserHandler{
		userService: service,
	}
}
func (h *UserHandler) RegisterRoutes(api fiber.Router) {

	users := api.Group("/users")
	users.Post("/register", h.Register)
	users.Post("/login", h.Login)
	users.Post("/refresh", h.Refresh)

}
func (h *UserHandler) Register(c *fiber.Ctx) error {

	var req RegisterRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	if err := utils.Validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": utils.ValidationError(req, err)})
	}

	token, refreshToken, err := h.userService.Register(c.Context(), req.Username, req.Email, req.Password, req.Phone)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		HTTPOnly: true,
		Secure:   false,
		Path:     "/",
		Expires:  time.Now().Add(30 * 24 * time.Hour),
	})

	return c.JSON(fiber.Map{
		"access_token": token,
	})
}
func (h *UserHandler) Login(c *fiber.Ctx) error {
	var req LoginRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "wrong email or password"})
	}
	token, err := h.userService.Login(c.Context(), req.Email, req.Password)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot find the user"})
	}

	return c.JSON(fiber.Map{
		"access_token": token,
	})

}

func (h *UserHandler) Refresh(c *fiber.Ctx) error {
	cookie := c.Cookies("refresh_token")

	if cookie == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "missing refresh token"})
	}

	userId, err := h.userService.VerifyRefreshToken(cookie)

	fmt.Println(userId, "=----")

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid refresh token"})
	}
	user, err := h.userService.GetById(c.Context(), userId)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "user not found"})
	}

	accessToken, err := utils.GenerateJWT(userId, user.Role)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot generate access token"})
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{"access_token": accessToken})
}
