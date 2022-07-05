package helpers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func MakeError(code int, err error) *fiber.Error {
	return fiber.NewError(code, fmt.Sprintf("error: %s", err.Error()))
}
