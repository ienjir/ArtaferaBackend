package testutils

import (
	"fmt"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"testing"
	"time"
)

// SetupTestDB creates an in-memory SQLite database for testing
func SetupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             time.Second,
				LogLevel:                  logger.Silent,
				IgnoreRecordNotFoundError: true,
				Colorful:                  false,
			},
		),
	})
	if err != nil {
		t.Fatalf("Failed to connect to test database: %v", err)
	}

	// Migrate all models
	if err := db.AutoMigrate(models.AllModels...); err != nil {
		t.Fatalf("Failed to migrate test database: %v", err)
	}

	return db
}

// SeedTestData creates test data in the database
func SeedTestData(t *testing.T, db *gorm.DB) {
	// Create test roles
	adminRole := models.Role{Name: "admin"}
	userRole := models.Role{Name: "user"}
	
	if err := db.Create(&adminRole).Error; err != nil {
		t.Fatalf("Failed to create admin role: %v", err)
	}
	if err := db.Create(&userRole).Error; err != nil {
		t.Fatalf("Failed to create user role: %v", err)
	}

	// Create test currency
	currency := models.Currency{
		CurrencyCode: "CHF",
		CurrencyName: "Swiss Franc",
	}
	if err := db.Create(&currency).Error; err != nil {
		t.Fatalf("Failed to create currency: %v", err)
	}

	// Create test language
	language := models.Language{
		LanguageName: "English",
		LanguageCode: "en",
	}
	if err := db.Create(&language).Error; err != nil {
		t.Fatalf("Failed to create language: %v", err)
	}

	// Create test users
	adminUser := models.User{
		Firstname: "Admin",
		Lastname:  "User",
		Email:     "admin@example.com",
		Password:  []byte("hashed_password"),
		Salt:      []byte("salt"),
		RoleID:    adminRole.ID,
	}
	regularUser := models.User{
		Firstname: "Regular",
		Lastname:  "User",
		Email:     "user@example.com",
		Password:  []byte("hashed_password"),
		Salt:      []byte("salt"),
		RoleID:    userRole.ID,
	}

	if err := db.Create(&adminUser).Error; err != nil {
		t.Fatalf("Failed to create admin user: %v", err)
	}
	if err := db.Create(&regularUser).Error; err != nil {
		t.Fatalf("Failed to create regular user: %v", err)
	}
}

// CleanupTestDB closes the database connection
func CleanupTestDB(db *gorm.DB) {
	if sqlDB, err := db.DB(); err == nil {
		sqlDB.Close()
	}
}

// AssertNoError is a helper to assert no error occurred
func AssertNoError(t *testing.T, err error) {
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}
}

// AssertError is a helper to assert an error occurred
func AssertError(t *testing.T, err error) {
	if err == nil {
		t.Fatal("Expected error, got nil")
	}
}

// AssertEqual is a helper to assert two values are equal
func AssertEqual[T comparable](t *testing.T, expected, actual T) {
	if expected != actual {
		t.Fatalf("Expected %v, got %v", expected, actual)
	}
}

// AssertNotNil is a helper to assert a value is not nil
func AssertNotNil(t *testing.T, value interface{}) {
	if value == nil {
		t.Fatal("Expected non-nil value, got nil")
	}
}

// AssertNil is a helper to assert a value is nil
func AssertNil(t *testing.T, value interface{}) {
	if value != nil {
		t.Fatalf("Expected nil value, got %v", value)
	}
}

// CreateTestUser creates a test user with default values
func CreateTestUser(db *gorm.DB, email string, roleID int64) *models.User {
	user := &models.User{
		Firstname: "Test",
		Lastname:  "User",
		Email:     email,
		Password:  []byte("hashed_password"),
		Salt:      []byte("salt"),
		RoleID:    roleID,
	}
	db.Create(user)
	return user
}

// CreateTestArt creates a test art with default values
func CreateTestArt(db *gorm.DB, currencyID int64) *models.Art {
	art := &models.Art{
		Price:        10000, // 100.00 in cents
		CurrencyID:   currencyID,
		CreationYear: 2023,
		Available:    true,
		Visible:      true,
	}
	db.Create(art)
	return art
}

// GetTestUserEmail generates unique test email
func GetTestUserEmail(prefix string) string {
	return fmt.Sprintf("%s_%d@test.com", prefix, time.Now().UnixNano())
}