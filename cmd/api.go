package main

import (
	"log"

	"github.com/CudoCommunication/cudocomm/config"
	database "github.com/CudoCommunication/cudocomm/pkg/database/gorm"
	"github.com/CudoCommunication/cudocomm/server"
)

func main() {
	log.Println("Starting api server")

	config.LoadEnvironmentFile(".env")

	db, err := database.NewGorm()
	if err != nil {
		log.Fatalf("db init: %s", err)
	}

	s := server.NewServer(db)

	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
