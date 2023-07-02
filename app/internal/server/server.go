package server

import (
	"context"
	"fmt"
	"log"
	"product/config"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	fiber *fiber.App
	cfg   *config.Config
}

func NewServer(cfg *config.Config) *Server {
	return &Server{
		fiber: fiber.New(fiber.Config{DisableStartupMessage: true}),
		cfg:   cfg,
	}
}

func (s *Server) Run(ctx context.Context) error {
	if err := s.MapHandlers(ctx, s.fiber); err != nil {
		log.Fatalf("Cannot map handlers: %s", err.Error())
	}
	if err := s.fiber.Listen(fmt.Sprintf("%s:%s", s.cfg.Server.Host, s.cfg.Server.HTTPPort)); err != nil {
		log.Fatalf("Error starting Server: %s", err.Error())
	}

	return nil
}
