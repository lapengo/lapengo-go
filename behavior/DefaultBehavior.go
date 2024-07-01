package behavior

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/lapengo/lapengo-go/behavior/_interface"
	"github.com/lapengo/lapengo-go/internal/base/behavior"
	ModelResUsersDTO "github.com/lapengo/lapengo-go/models/response/users"
	"gorm.io/gorm"
)

type DefaultBehaviorImpl struct {
	behavior.BaseBehavior
}

func NewDefaultBehavior(DB *gorm.DB, validate *validator.Validate) _interface.DefaultInterface {
	return &DefaultBehaviorImpl{
		BaseBehavior: behavior.BaseBehavior{
			DB:        DB,
			Validate:  validate,
			TableName: "public.users",
		},
	}
}

func (b *DefaultBehaviorImpl) GetAll(ctx *fiber.Ctx) (res ModelResUsersDTO.UserDTO, err error) {
	err = b.SimpleGetAll(ctx, &res, ModelResUsersDTO.UserDTO{})
	return
}
