package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"product/config"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"

	pb "product/pkg/proto"

	"os"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zerolog "github.com/philip-bui/grpc-zerolog"
	"github.com/rs/zerolog"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
)

var (
	zerologger = zerolog.New(os.Stdout).With().Timestamp().Logger()
)

type Deps struct {
	ProductDeps pb.ProductServer
}

type Server struct {
	fiber *fiber.App
	grpc  *grpc.Server
	cfg   *config.Config
	deps  Deps
}

func NewServer(cfg *config.Config) *Server {
	return &Server{
		fiber: fiber.New(fiber.Config{DisableStartupMessage: true}),
		grpc: grpc.NewServer(
			grpc.UnaryInterceptor(
				grpc_middleware.ChainUnaryServer(
					otelgrpc.UnaryServerInterceptor(),
					grpc_zerolog.NewUnaryServerInterceptorWithLogger(&zerologger),
				),
			),
		),
		deps: Deps{},
		cfg:  cfg,
	}
}

func (s *Server) Run(ctx context.Context) error {
	if err := s.MapHandlers(ctx, s.fiber, s.cfg); err != nil {
		log.Printf("Cannot map handlers: %s", err.Error())
		return err
	}
	log.Println("Map handled")

	go func(s *Server) {
		log.Printf("HTTP Server starts on: %s:%s", s.cfg.Server.Host, s.cfg.Server.HTTPPort)
		if err := runHTTP(ctx, s); err != nil {
			log.Println(err)
			return
		}
	}(s)

	go func(s *Server) {
		log.Printf("gRPC Server starts on: %s:%s", s.cfg.Server.Host, s.cfg.Server.GRPCPort)
		if err := runGRPC(ctx, s); err != nil {
			log.Println(err)
			return
		}
	}(s)

	return nil
}

func runHTTP(ctx context.Context, s *Server) error {
	if err := s.fiber.Listen(fmt.Sprintf("%s:%s", s.cfg.Server.Host, s.cfg.Server.HTTPPort)); err != nil {
		log.Fatalf("Error starting Server: %s", err.Error())
	}

	return nil
}

func runGRPC(ctx context.Context, s *Server) error {
	pb.RegisterProductServer(s.grpc, s.deps.ProductDeps)

	l, err := net.Listen("tcp", s.cfg.Server.Host+":"+s.cfg.Server.GRPCPort)
	if err != nil {
		return err
	}

	go func() {
		if err := s.grpc.Serve(l); err != nil {
			log.Fatal(err.Error())
		}
	}()

	return nil
}

func (s *Server) Shutdown() {
	if err := s.fiber.Shutdown(); err != nil {
		log.Println(err)
		return
	}
	log.Println("HTTP Server was stoped properly")

	s.grpc.GracefulStop()
	log.Println("gRPC Server was stoped properly")
}
