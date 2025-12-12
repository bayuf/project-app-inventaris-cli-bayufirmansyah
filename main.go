package main

import (
	"fmt"
	"os"

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
		fmt.Println(err)
		os.Exit(1)
	}
	defer dbConn.Close()

	// Init
	repo := repository.NewInventoryRepository(dbConn)
	svc := service.NewInventoryService(repo)
	InvHandler := handler.NewInventoryHandler(svc)

	cmd.Homepage(InvHandler)
}
