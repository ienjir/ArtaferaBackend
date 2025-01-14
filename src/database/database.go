package database

import (
	"fmt"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func ConnectDatabase() error {
	var err error

	host := os.Getenv("host")
	user := os.Getenv("user")
	password := os.Getenv("password")
	dbname := os.Getenv("dbname")
	port := os.Getenv("port")
	sslmode := os.Getenv("sslmode")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, user, password, dbname, port, sslmode)

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
