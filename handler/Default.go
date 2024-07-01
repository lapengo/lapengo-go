package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lapengo/lapengo-go/behavior/_interface"
	"github.com/lapengo/lapengo-go/internal/helper"
)

type DefaultHandler interface {
	GetAll(*fiber.Ctx) error
}

type DefaultHandlerImpl struct {
	behavior _interface.DefaultInterface
}

func NewDefaultHandler(b _interface.DefaultInterface) DefaultHandler {
	return &DefaultHandlerImpl{
		behavior: b,
	}
}

func (h *DefaultHandlerImpl) GetAll(ctx *fiber.Ctx) error {
	result, err := h.behavior.GetAll(ctx)

	return helper.GetWebResponse(ctx, result, err)
}
