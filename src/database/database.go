package database

import (
	"github.com/ienjir/ArtaferaBackend/src/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectDatabase() error {
	var err error

	dsn := "host=localhost user=DBAdmin password=AVerySecurePassword dbname=ArtaferaDB port=5432 sslmode=disable"

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	// Drop all tables
	for _, model := range models.AllModels {
		err := DB.Migrator().DropTable(model)
		if err != nil {
			log.Printf("Failed to drop table: %v", err)
		}
	}

	// Make all tables
	err = DB.AutoMigrate(models.AllModels...)
	if err != nil {
		log.Fatalf("Failed to migrate tables: %v", err)
		return err
	}

	log.Println("Database connected and migrated")
	return err
}
