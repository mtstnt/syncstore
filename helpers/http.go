package helpers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func ErrorResponse(ctx *fiber.Ctx, status int, message string) error {
	return ctx.Status(status).JSON(fiber.Map{
		"result":  false,
		"message": message,
		"data":    nil,
	})
}

func ErrorUnauthorizedResponse(ctx *fiber.Ctx) error {
	return ErrorResponse(
		ctx,
		fiber.StatusUnauthorized,
		"Invalid login credentials",
	)
}

func ErrorInternalSystemResponse(ctx *fiber.Ctx, err error) error {
	return ErrorResponse(
		ctx,
		fiber.StatusInternalServerError,
		fmt.Sprintf("Error occurred: %s", err.Error()),
	)
}

func SuccessResponse(ctx *fiber.Ctx, message string, data any) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"result":  true,
		"message": message,
		"data":    data,
	})
}
