package helpers

import (
	"github.com/gofiber/fiber/v2"
)

func ErrorResponse(ctx *fiber.Ctx, status int, message string) error {
	return ctx.Status(status).JSON(fiber.Map{
		"result":  false,
		"message": message,
		"data":    nil,
	})
}

func ErrorResponseFromErr(ctx *fiber.Ctx, err *fiber.Error) error {
	return ctx.Status(err.Code).JSON(fiber.Map{
		"result":  false,
		"message": err.Message,
		"data":    nil,
	})
}

func SuccessResponse(ctx *fiber.Ctx, message string, data any) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"result":  true,
		"message": message,
		"data":    data,
	})
}
