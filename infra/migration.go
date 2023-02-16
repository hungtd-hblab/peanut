package infra

import (
	"log"
	"peanut/domain"

	"gorm.io/gorm"
)

func Migration(db *gorm.DB) {
	err := db.AutoMigrate(
		&domain.User{},
	)

	if err != nil {
		log.Printf("Migrate failed: %v", err)
	}
}
