// main.go
package main

import (
	"fmt"
	"github.com/DarioKnezovic/campaign-service/api"
	"github.com/DarioKnezovic/campaign-service/config"
	"github.com/DarioKnezovic/campaign-service/internal/campaign/repository"
	"github.com/DarioKnezovic/campaign-service/internal/campaign/service"
	"github.com/DarioKnezovic/campaign-service/pkg/database"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	cfg := config.LoadConfig()

	// Connect to the database
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	err = database.PerformAutoMigrations(db)
	if err != nil {
		log.Fatalf("Failed to perform auto migrations: %v", err)
	}

	campaignRepo := repository.NewCampaignRepository(db)
	campaignService := service.NewCampaignService(campaignRepo)

	// Create a new Gorilla Mux router
	router := mux.NewRouter()

	// Register routes using Gorilla Mux
	api.RegisterRoutes(router, campaignService)

	log.Printf("Server listening on port %s", cfg.APIPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", cfg.APIPort), router))
}
