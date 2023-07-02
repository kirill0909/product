package http

import (
	"product/internal/product"

	"github.com/gofiber/fiber/v2"
)

func MapProductRoutes(productRoutes fiber.Router, h product.Handler) {
	productRoutes.Get("/get_product", h.GetProduct())
}
