package auth

import (
	"encoding/base64"
	"time"

	"syncstore/helpers"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Handler struct {
	repo Repository
}

func NewHandler(db *gorm.DB) Handler {
	return Handler{
		repo: NewRepository(db),
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

	user, err := h.repo.GetByUsername(req.Username)
	if user == nil {
		return helpers.ErrorUnauthorizedResponse(ctx)
	}
	if err != nil {
		return helpers.ErrorInternalSystemResponse(ctx, err)
	}

	decodedPassword, err := base64.StdEncoding.DecodeString(req.PasswordB64)
	if err != nil {
		return helpers.ErrorUnauthorizedResponse(ctx)
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), decodedPassword); err != nil {
		return helpers.ErrorUnauthorizedResponse(ctx)
	}

	token := ""

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

	password, err := base64.StdEncoding.DecodeString(req.PasswordB64)
	if err != nil {
		return helpers.ErrorResponse(
			ctx,
			fiber.StatusBadRequest,
			"Password is incorrectly sent",
		)
	}

	user, err := h.repo.CreateUser(req.Username, string(password))
	if err != nil {
		return helpers.ErrorInternalSystemResponse(ctx, err)
	}

	return helpers.SuccessResponse(ctx, "Successful registration", fiber.Map{"user": user})
}
