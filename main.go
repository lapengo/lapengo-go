package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lapengo/lapengo-go/internal/helper"
	"os"
)

func main() {
	// Custom config
	app := fiber.New(fiber.Config{
		//Prefork:       true,
		//CaseSensitive: true,
		//StrictRouting: true,
		//ServerHeader: "Fiber",
		AppName:   "lapengo",
		BodyLimit: 1024 * 1024 * 200,
	})

	app.Static("/", "./public")
	setup(app)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	err := app.Listen(":" + port)

	// handle error
	if err != nil {
		helper.PanicIfError(err)
	}

}
