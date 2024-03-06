package behavior

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/lapengo/lapengo-go/internal/base/behavior"
	"gorm.io/gorm"
)

type AuthBehavior interface {
	GetAll(*fiber.Ctx) (string, error)
}

type AuthBehaviorImpl struct {
	behavior.BaseBehavior
}

func NewAuthBehavior(DB *gorm.DB, validate *validator.Validate) AuthBehavior {
	return &AuthBehaviorImpl{
		BaseBehavior: behavior.BaseBehavior{
			DB:       DB,
			Validate: validate,
			//TableName: "users",
		},
	}
}

func (b *AuthBehaviorImpl) GetAll(ctx *fiber.Ctx) (res string, err error) {
	res = "Hello World"
	return
}
