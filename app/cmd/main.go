package main

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
	"os"
	"os/signal"
	"product/config"
	grpcProduct "product/internal/product/delivery/grpc"
	httpProduct "product/internal/product/delivery/http"
	usecaseProduct "product/internal/product/usecase"
	"product/internal/server"
	"syscall"
)

func main() {
	viper, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	cfg, err := config.ParseConfig(viper)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Config loaded")

	app, deps := mapHandler(cfg)
	server := server.NewServer(app, deps, cfg)

	ctx := context.Background()
	if err := server.Run(ctx); err != nil {
		log.Println(err)
		return
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	server.Shutdown()
}

func mapHandler(cfg *config.Config) (*fiber.App, server.Deps) {
	// create App
	app := fiber.New(fiber.Config{DisableStartupMessage: true})
	app.Use(logger.New())

	// usecase
	productUC := usecaseProduct.NewProducetUsecase(cfg)

	// handler
	productHTTP := httpProduct.NewProductHandler(cfg, productUC)
	productGRPC := grpcProduct.NewProductHandler(productUC)

	// groups
	apiGroup := app.Group("api")
	productGroup := apiGroup.Group("product")

	// routes
	httpProduct.MapProductRoutes(productGroup, productHTTP)

	// create grpc dependencyes
	deps := server.Deps{ProductDeps: productGRPC}
	return app, deps
}
