package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"product/config"
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

	ctx := context.Background()
	httpServer := server.NewServer(cfg)
	go func() {
		log.Printf("HTTP Server starts on: %s:%s", cfg.Server.Host, cfg.Server.HTTPPort)
		if err := httpServer.Run(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
}
