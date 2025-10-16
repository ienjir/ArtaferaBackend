package database

import (
	"fmt"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"github.com/ienjir/ArtaferaBackend/src/repository"
	"github.com/ienjir/ArtaferaBackend/src/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"errors"
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

	log.Println(host)
	log.Println(user)
	log.Println(password)
	log.Println(dbname)
	log.Println(port)
	log.Println(sslmode)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		host, user, password, dbname, port, sslmode)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	// Drop all tables if not prod
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
	tempRepos := repository.NewRepositoryManager(DB)

	roles := []string{"user", "admin", "artist"}

	for _, roleName := range roles {
		var existing models.Role
		// Check if role already exists
		err := DB.Where("name = ?", roleName).First(&existing).Error

		if err == nil {
			// Role already exists, skip
			continue
		}
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			// Unexpected DB error
			return fmt.Errorf("failed to check role %s: %v", roleName, err)
		}

		// Role doesn't exist → create it
		newRole := models.Role{Name: roleName}
		if createErr := tempRepos.Role.Create(&newRole); createErr != nil {
			return fmt.Errorf("failed to create %s role: %v", roleName, createErr.Message)
		}
	}

	return nil
}
