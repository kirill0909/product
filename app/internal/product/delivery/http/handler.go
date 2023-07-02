package http

import (
	"log"
	"product/config"
	models "product/internal/models/product"
	"product/internal/product"

	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

type ProductHandler struct {
	cfg       *config.Config
	productUC product.Usecase
}

func NewProductHandler(cfg *config.Config, productUC product.Usecase) product.Handler {
	return &ProductHandler{cfg: cfg, productUC: productUC}
}

func (h *ProductHandler) GetProduct() fiber.Handler {
	return func(c *fiber.Ctx) error {

		var getProductRequest models.GetProductRequest
		if err := c.QueryParser(&getProductRequest); err != nil {
			log.Println(err)
			return err
		}
		if getProductRequest.ID == 0 {
			c.SendStatus(fiber.ErrBadRequest.Code)
			return errors.New("Unable to parse query")
		}

		result, err := h.productUC.GetProduct(getProductRequest)
		if err != nil {
			log.Println(err)
			return err
		}

		return c.JSON(result)
	}
}
