package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/lapengo/lapengo-go/internal/config"
	"github.com/lapengo/lapengo-go/route"
)

func setup(app *fiber.App) {
	config.InitDB()

	db := config.DBConn
	validate := validator.New()
	api := app.Group("/api/v1")

	route.DefaultRoute(api.Group("/default"), db, validate)

}
