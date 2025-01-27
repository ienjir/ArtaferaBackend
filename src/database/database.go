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

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	sslmode := os.Getenv("DB_SSL_MODE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		host, user, password, dbname, port, sslmode)

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
		log.Printf("Failed to migrate tables: %v", err)
		return err
	}

	// Set the default value for role_id in users table
	DB.Exec("ALTER TABLE users ALTER COLUMN role_id SET DEFAULT 1")

	// Create initial roles
	err = createInitialRoles()
	if err != nil {
		log.Printf("Failed to create initial roles: %v", err)
		return err
	}

	log.Println("Database connected and migrated successfully")
	return nil
}

func createInitialRoles() error {
	userRole := models.Role{
		Role: "user",
	}
	result := DB.Create(&userRole)
	if result.Error != nil {
		return result.Error
	}

	adminRole := models.Role{
		Role: "admin",
	}
	result = DB.Create(&adminRole)
	if result.Error != nil {
		return result.Error
	}
	
	artistRole := models.Role{
		Role: "artist",
	}
	result = DB.Create(&artistRole)

	return nil
}
