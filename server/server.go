package server

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/CudoCommunication/cudocomm/config"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Server struct {
	echo *echo.Echo
	db   *gorm.DB
}

func NewServer(db *gorm.DB) *Server {
	return &Server{
		echo: echo.New(),
		db:   db,
	}
}

func (s *Server) Run() error {

	if err := s.MapHandlers(s.echo); err != nil {
		log.Fatalf("failed to map handlers: %v", err)
	}

	go func() {
		serverAddr := fmt.Sprintf(":%s", config.Env.AppPort)
		log.Printf("Server is listening on PORT %s", config.Env.AppPort)
		if err := s.echo.Start(serverAddr); err != nil {
			log.Fatalf("Error starting Server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")
	return nil
}
