package user

import (
	"time"

	"syncstore/helpers"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Handler struct {
	service Service
}

func NewHandler(db *gorm.DB) Handler {
	return Handler{
		NewService(NewRepository(db)),
	}
}

func (h Handler) Login(ctx *fiber.Ctx) error {
	req := LoginRequest{}
	if err := ctx.BodyParser(&req); err != nil {
		return helpers.ErrorResponse(
			ctx,
			fiber.StatusBadRequest,
			"Request must only consist of \"username\" and \"password\" only",
		)
	}

	token, err := h.service.Login(req.Username, req.PasswordB64)
	if err != nil {
		return helpers.ErrorResponseFromErr(ctx, err)
	}

	ctx.Cookie(&fiber.Cookie{
		Name:     "SyncSessID",
		Value:    token,
		Expires:  time.Now().Add(1 * time.Hour),
		HTTPOnly: true,
	})

	return helpers.SuccessResponse(ctx, "Login successful", fiber.Map{"token": token})
}

func (h Handler) Register(ctx *fiber.Ctx) error {
	req := LoginRequest{}
	if err := ctx.BodyParser(&req); err != nil {
		return helpers.ErrorResponse(
			ctx,
			fiber.StatusBadRequest,
			"Request must only consist of \"username\" and \"password\" only",
		)
	}

	user, token, err := h.service.Register(req.Username, req.PasswordB64)
	if err != nil {
		return helpers.ErrorResponseFromErr(ctx, err)
	}

	ctx.Cookie(&fiber.Cookie{
		Name:     "SyncSessID",
		Value:    token,
		Expires:  time.Now().Add(1 * time.Hour),
		HTTPOnly: true,
	})

	return helpers.SuccessResponse(ctx, "Successful registration", fiber.Map{"user": user, "token": token})
}
