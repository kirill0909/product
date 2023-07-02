package server

import (
	"context"

	"github.com/gofiber/fiber/v2/middleware/logger"
	"product/config"
	httpProduct "product/internal/product/delivery/http"
	usecaseProduct "product/internal/product/usecase"

	"github.com/gofiber/fiber/v2"
)

func (s *Server) MapHandlers(ctx context.Context, app *fiber.App, cfg *config.Config) error {

	app.Use(logger.New())

	productUC := usecaseProduct.NewProducetUsecase(cfg)
	productHTTP := httpProduct.NewProductHandler(cfg, productUC)

	apiGroup := app.Group("api")
	productGroup := apiGroup.Group("product")

	httpProduct.MapProductRoutes(productGroup, productHTTP)

	return nil
}
