package main

import (
	"log"

	"github.com/bayuf/project-app-inventaris-cli-bayufirmansyah/cmd"
	"github.com/bayuf/project-app-inventaris-cli-bayufirmansyah/db"
	"github.com/bayuf/project-app-inventaris-cli-bayufirmansyah/handler"
	"github.com/bayuf/project-app-inventaris-cli-bayufirmansyah/repository"
	"github.com/bayuf/project-app-inventaris-cli-bayufirmansyah/service"
	"github.com/joho/godotenv"
)

func main() {
	// load env file
	_ = godotenv.Load()

	// init db
	dbConn, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer dbConn.Close()

	// Init
	repo := repository.NewInventoryRepository(dbConn)
	svc := service.NewInventoryService(repo)
	InvHandler := handler.NewInventoryHandler(svc)

	cmd.Homepage(InvHandler)
}
