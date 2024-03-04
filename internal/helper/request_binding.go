package helper

import (
	"errors"
	"github.com/gofiber/fiber/v2"
)

func BodyParser(ctx *fiber.Ctx, model interface{}) (err error) {
	err = ctx.BodyParser(model)

	if err != nil {
		err = errors.New("BAD REQUEST | " + err.Error())
	}

	return
}
