package sync

import "github.com/gofiber/fiber/v2"

type SyncHandler struct {
}

func NewHandler() SyncHandler {
	return SyncHandler{}
}

func (h SyncHandler) SyncFolder(ctx *fiber.Ctx) error {
	return nil
}
