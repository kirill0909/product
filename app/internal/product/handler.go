package product

import "github.com/gofiber/fiber/v2"

type Handler interface {
	GetProduct() fiber.Handler
}
