package behavior

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/lapengo/lapengo-go/internal/base/behavior"
	"gorm.io/gorm"
)

type DefaultBehavior interface {
	GetAll(*fiber.Ctx) (string, error)
}

type DefaultBehaviorImpl struct {
	behavior.BaseBehavior
}

func NewDefaultBehavior(DB *gorm.DB, validate *validator.Validate) DefaultBehavior {
	return &DefaultBehaviorImpl{
		BaseBehavior: behavior.BaseBehavior{
			DB:       DB,
			Validate: validate,
			//TableName: "users",
		},
	}
}

func (b *DefaultBehaviorImpl) GetAll(ctx *fiber.Ctx) (res string, err error) {
	res = "Hello World"
	return
}
