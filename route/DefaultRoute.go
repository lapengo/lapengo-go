package route

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/lapengo/lapengo-go/behavior"
	"github.com/lapengo/lapengo-go/handler"
	"gorm.io/gorm"
)

func DefaultRoute(route fiber.Router, db *gorm.DB, validate *validator.Validate) {
	defaultBehavior := behavior.NewDefaultBehavior(db, validate)
	defaultHandler := handler.NewDefaultHandler(defaultBehavior)

	route.Get("/", defaultHandler.GetAll)
}
