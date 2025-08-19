package database

import (
	"fmt"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"github.com/ienjir/ArtaferaBackend/src/repository"
	"github.com/ienjir/ArtaferaBackend/src/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB
var Repositories *repository.RepositoryManager

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
	if utils.GinMode != 2 {
		for _, model := range models.AllModels {
			err := DB.Migrator().DropTable(model)
			if err != nil {
				log.Printf("Failed to drop table: %v", err)
			}
		}
	}

	// Make all tables
	err = DB.AutoMigrate(models.AllModels...)
	if err != nil {
		log.Printf("Failed to migrate tables: %v", err)
		return err
	}

	// Create initial roles
	err = createInitialRoles()
	if err != nil {
		log.Printf("Failed to create initial roles: %v", err)
		return err
	}

	// Initialize repository manager
	Repositories = repository.NewRepositoryManager(DB)

	log.Println("Database connected and migrated successfully")
	return nil
}

func createInitialRoles() error {
	// Initialize repository manager first for this function
	tempRepos := repository.NewRepositoryManager(DB)

	userRole := models.Role{
		Name: "user",
	}
	if err := tempRepos.Role.Create(&userRole); err != nil {
		return fmt.Errorf("failed to create user role: %v", err.Message)
	}

	adminRole := models.Role{
		Name: "admin",
	}
	if err := tempRepos.Role.Create(&adminRole); err != nil {
		return fmt.Errorf("failed to create admin role: %v", err.Message)
	}

	artistRole := models.Role{
		Name: "artist",
	}
	if err := tempRepos.Role.Create(&artistRole); err != nil {
		return fmt.Errorf("failed to create artist role: %v", err.Message)
	}

	return nil
}
