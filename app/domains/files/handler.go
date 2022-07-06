package files

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Handler struct {
	service Service
}

func NewHandler(db *gorm.DB) Handler {
	return Handler{
		service: NewService(NewRepository(db)),
	}
}

func (h Handler) GetAllFolders(ctx *fiber.Ctx) error {
	return nil
}

func (h Handler) GetFolder(ctx *fiber.Ctx) error {
	return nil
}

func (h Handler) AddFolderToTrack(ctx *fiber.Ctx) error {
	return nil
}

func (h Handler) DeleteTrackedFolder(ctx *fiber.Ctx) error {
	return nil
}

func (h Handler) GetFile(ctx *fiber.Ctx) error {
	return nil
}
