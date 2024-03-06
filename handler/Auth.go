package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lapengo/lapengo-go/behavior"
	"github.com/lapengo/lapengo-go/internal/helper"
)

type AuthHandler interface {
	GetAll(*fiber.Ctx) error
}

type AuthHandlerImpl struct {
	behavior behavior.AuthBehavior
}

func NewAuthHandler(b behavior.AuthBehavior) AuthHandler {
	return &AuthHandlerImpl{
		behavior: b,
	}
}

func (h *AuthHandlerImpl) GetAll(ctx *fiber.Ctx) error {
	result, err := h.behavior.GetAll(ctx)

	return helper.GetWebResponse(ctx, result, err)
}
