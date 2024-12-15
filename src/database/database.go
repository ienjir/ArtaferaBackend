package database

import (
	"github.com/ienjir/ArtaferaBackend/src/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func ConnectDatabase() {

	// Get the database DSN (default to "database.db" if not set)
	dsn := os.Getenv("DATABASE_DSN")
	if dsn == "" {
		dsn = "database.db"
	}

	// Open the SQLite database
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Automatically migrate models (creates or updates tables)
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	DB = db
	log.Println("Database connected and migrated")
}
