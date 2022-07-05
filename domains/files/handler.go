package files

import "github.com/gofiber/fiber/v2"

type Handler struct {
}

func NewHandler() Handler {
	return Handler{}
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

func (h Handler) RemoveFolderFromTrack(ctx *fiber.Ctx) error {
	return nil
}
