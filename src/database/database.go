package database

import (
	"fmt"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=DBAdmin password=AVerySecurePassword dbname=ArtaferaDB port=5432 sslmode=disable"

	// Open the SQLite database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Failed to migrate the schema:", err)
	}

	fmt.Println("Table 'users' created successfully!")

	DB = db
	log.Println("Database connected and migrated")
}
