package _interface

import (
	"github.com/gofiber/fiber/v2"
	ModelResUsersDTO "github.com/lapengo/lapengo-go/models/response/users"
)

type DefaultInterface interface {
	GetAll(*fiber.Ctx) (ModelResUsersDTO.UserDTO, error)
}
