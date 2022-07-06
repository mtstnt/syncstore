package main

import (
	"log"
	"syncstore/domains/files"
	"syncstore/domains/sync"
	"syncstore/domains/user"
	"syncstore/session"

	"gorm.io/driver/sqlite"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func initDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("database/my.db"))
	if err != nil {
		return nil, err
	}

	if err = db.AutoMigrate(user.User{}, files.Folder{}); err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	app := fiber.New()

	db, err := initDB()
	if err != nil {
		log.Fatalln(err)
	}

	session := session.GetSessionMap()

	userHandler := user.NewHandler(db, session)
	filesHandler := files.NewHandler(db)
	syncHandler := sync.NewHandler()
	app.Route("/api", func(app fiber.Router) {
		app.Route("/auth", func(router fiber.Router) {
			router.Post("/login", userHandler.Login)
			router.Post("/register", userHandler.Register)
		})

		app.Route("/folders", func(router fiber.Router) {
			router.Get("/", filesHandler.GetAllFolders)
			router.Get("/:folder_id", filesHandler.GetFolder)

			router.Post("/", filesHandler.AddFolderToTrack)
			router.Delete("/:folder_id", filesHandler.DeleteTrackedFolder)

			// Sync sends the checksum of the local folder contents
			//  and returns the required to download and required to delete contents.
			router.Post("/sync", syncHandler.SyncFolder)

			router.Get("/:folder_id/download/:file_name", filesHandler.GetFile)
		})
	})

	app.Listen(":8080")
}
