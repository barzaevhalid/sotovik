package middleware

import (
	"errors"
	"log"

	"github.com/barzaevhalid/sotovik/internal/api/user"
	"github.com/barzaevhalid/sotovik/internal/domain"
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	var fiberErr *fiber.Error

	if errors.As(err, &fiberErr) {
		return c.Status(fiberErr.Code).JSON(fiber.Map{
			"error": fiberErr.Message,
		})
	}

	switch {
	case errors.Is(err, user.ErrInvalidCredentials):
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid email or password",
		})
	case errors.Is(err, user.ErrForbidden):
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "forbidden",
		})
	case errors.Is(err, domain.ErrNotFound):
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "not found",
		})
	}
	log.Printf("INTERNAL ERROR: %+v", err)
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"error": "internal server error",
	})
}
