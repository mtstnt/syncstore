package main

import (
	"log"
	"syncstore/domains/auth"

	"gorm.io/driver/sqlite"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func main() {
	app := fiber.New()

	db, err := gorm.Open(sqlite.Open("database/my.db"))
	if err != nil {
		log.Fatalln("Failed to open database")
	}

	authHandler := auth.NewHandler(db)
	// filesHandler := files.NewHandler()

	app.Route("/auth", func(router fiber.Router) {
		router.Post("/login", authHandler.Login)
		router.Post("/register", authHandler.Register)
	})

	// app.Route("/folders", func(router fiber.Router) {
	// 	router.Get("/", filesHandler.GetAllFolders)
	// 	router.Get("/", filesHandler.GetFolder)

	// 	router.Post("/", filesHandler.AddFolderToTrack)
	// })

	// app.Route("/sync", func(router fiber.Router) {
	// 	router.Get("/folder/{folder_id}")
	// })

	app.Listen(":8080")
}