package database

import (
	"github.com/DarioKnezovic/campaign-service/internal/campaign"
	"gorm.io/gorm"
	"log"
)

func PerformAutoMigrations(db *gorm.DB) error {
	err := db.AutoMigrate(&campaign.Campaign{})
	if err != nil {
		return err
	}
	log.Print("Auto migrations have been successfully finished")

	return nil
}
