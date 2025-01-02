package database

import (
	"github.com/ienjir/ArtaferaBackend/src/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectDatabase() *gorm.DB {
	dsn := "host=localhost user=DBAdmin password=AVerySecurePassword dbname=ArtaferaDB port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Drop all tables
	for _, model := range models.AllModels {
		err := db.Migrator().DropTable(model)
		if err != nil {
			log.Printf("Failed to drop table: %v", err)
		}
	}

	// Make all tables
	err = db.AutoMigrate(models.AllModels...)
	if err != nil {
		log.Fatalf("Failed to migrate tables: %v", err)
	}

	DB = db
	log.Println("Database connected and migrated")

	return DB
}
