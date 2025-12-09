package repository

import (
	"testing"

	"github.com/ienjir/ArtaferaBackend/src/models"
	"github.com/ienjir/ArtaferaBackend/src/testutils"
)

func TestGormRepository_Create(t *testing.T) {
	db := testutils.SetupTestDB(t)
	defer testutils.CleanupTestDB(db)
	testutils.SeedTestData(t, db)

	repo := NewRepository[models.User](db)

	user := &models.User{
		Firstname: "John",
		Lastname:  "Doe",
		Email:     "john@example.com",
		Password:  []byte("hashed_password"),
		Salt:      []byte("salt"),
		RoleID:    1,
	}

	err := repo.Create(user)
	testutils.AssertNil(t, err)
	testutils.AssertNotNil(t, user)

	// Verify user was created with ID
	if user.ID == 0 {
		t.Error("Expected user ID to be set after creation")
	}
}

func TestGormRepository_GetByID(t *testing.T) {
	db := testutils.SetupTestDB(t)
	defer testutils.CleanupTestDB(db)
	testutils.SeedTestData(t, db)

	repo := NewRepository[models.User](db)

	// Create a test user first
	user := testutils.CreateTestUser(db, "test@example.com", 1)

	// Test getting existing user
	retrieved, err := repo.GetByID(user.ID)
	testutils.AssertNil(t, err)
	testutils.AssertNotNil(t, retrieved)
	testutils.AssertEqual(t, user.Email, retrieved.Email)
	testutils.AssertEqual(t, user.Firstname, retrieved.Firstname)

	// Test getting non-existent user
	_, err = repo.GetByID(99999)
	testutils.AssertNotNil(t, err)
	testutils.AssertEqual(t, 404, err.StatusCode)
}

func TestGormRepository_GetByIDWithPreload(t *testing.T) {
	db := testutils.SetupTestDB(t)
	defer testutils.CleanupTestDB(db)
	testutils.SeedTestData(t, db)

	repo := NewRepository[models.User](db)

	// Create a test user
	user := testutils.CreateTestUser(db, "test@example.com", 1)

	// Test getting user with Role preloaded
	retrieved, err := repo.GetByID(user.ID, "Role")
	testutils.AssertNil(t, err)
	testutils.AssertNotNil(t, retrieved)
	testutils.AssertNotNil(t, retrieved.Role)
	testutils.AssertEqual(t, "admin", retrieved.Role.Name)
}

func TestGormRepository_Update(t *testing.T) {
	db := testutils.SetupTestDB(t)
	defer testutils.CleanupTestDB(db)
	testutils.SeedTestData(t, db)

	repo := NewRepository[models.User](db)

	// Create a test user
	user := testutils.CreateTestUser(db, "test@example.com", 1)

	// Update user
	user.Firstname = "Updated"
	user.Lastname = "Name"

	err := repo.Update(user)
	testutils.AssertNil(t, err)

	// Verify update
	updated, err := repo.GetByID(user.ID)
	testutils.AssertNil(t, err)
	testutils.AssertEqual(t, "Updated", updated.Firstname)
	testutils.AssertEqual(t, "Name", updated.Lastname)
}

func TestGormRepository_UpdateFields(t *testing.T) {
	db := testutils.SetupTestDB(t)
	defer testutils.CleanupTestDB(db)
	testutils.SeedTestData(t, db)

	repo := NewRepository[models.User](db)

	// Create a test user
	user := testutils.CreateTestUser(db, "test@example.com", 1)

	// Update specific fields
	updates := map[string]interface{}{
		"firstname": "FieldUpdated",
		"lastname":  "FieldName",
	}

	updated, err := repo.UpdateFields(user.ID, updates)
	testutils.AssertNil(t, err)
	testutils.AssertNotNil(t, updated)
	testutils.AssertEqual(t, "FieldUpdated", updated.Firstname)
	testutils.AssertEqual(t, "FieldName", updated.Lastname)

	// Test updating non-existent user
	_, err = repo.UpdateFields(99999, updates)
	testutils.AssertNotNil(t, err)
	testutils.AssertEqual(t, 404, err.StatusCode)
}

func TestGormRepository_Delete(t *testing.T) {
	db := testutils.SetupTestDB(t)
	defer testutils.CleanupTestDB(db)
	testutils.SeedTestData(t, db)

	repo := NewRepository[models.User](db)

	// Create a test user
	user := testutils.CreateTestUser(db, "test@example.com", 1)

	// Delete user
	err := repo.Delete(user.ID)
	testutils.AssertNil(t, err)

	// Verify user is deleted (soft delete)
	_, err = repo.GetByID(user.ID)
	testutils.AssertNotNil(t, err)
	testutils.AssertEqual(t, 404, err.StatusCode)

	// Test deleting non-existent user
	err = repo.Delete(99999)
	testutils.AssertNotNil(t, err)
	testutils.AssertEqual(t, 404, err.StatusCode)
}

func TestGormRepository_DeleteEntity(t *testing.T) {
	db := testutils.SetupTestDB(t)
	defer testutils.CleanupTestDB(db)
	testutils.SeedTestData(t, db)

	repo := NewRepository[models.User](db)

	// Create a test user
	user := testutils.CreateTestUser(db, "test@example.com", 1)

	// Delete user entity
	err := repo.DeleteEntity(user)
	testutils.AssertNil(t, err)

	// Verify user is deleted
	_, err = repo.GetByID(user.ID)
	testutils.AssertNotNil(t, err)
	testutils.AssertEqual(t, 404, err.StatusCode)
}

func TestGormRepository_List(t *testing.T) {
	db := testutils.SetupTestDB(t)
	defer testutils.CleanupTestDB(db)
	testutils.SeedTestData(t, db)

	repo := NewRepository[models.User](db)

	// Create multiple test users
	testutils.CreateTestUser(db, "user1@example.com", 1)
	testutils.CreateTestUser(db, "user2@example.com", 1)
	testutils.CreateTestUser(db, "user3@example.com", 2)

	// Test listing with pagination
	users, err := repo.List(0, 2)
	testutils.AssertNil(t, err)
	testutils.AssertNotNil(t, users)

	// Should have at least 2 users (including seeded ones)
	if len(*users) < 2 {
		t.Errorf("Expected at least 2 users, got %d", len(*users))
	}

	// Test listing with preload
	users, err = repo.List(0, 5, "Role")
	testutils.AssertNil(t, err)
	testutils.AssertNotNil(t, users)

	// Check if Role is preloaded for first user
	if len(*users) > 0 {
		user := (*users)[0]
		testutils.AssertNotNil(t, user.Role)
	}
}

func TestGormRepository_Count(t *testing.T) {
	db := testutils.SetupTestDB(t)
	defer testutils.CleanupTestDB(db)
	testutils.SeedTestData(t, db)

	repo := NewRepository[models.User](db)

	// Get initial count (seeded users)
	initialCount, err := repo.Count()
	testutils.AssertNil(t, err)
	testutils.AssertNotNil(t, initialCount)

	// Create additional users
	testutils.CreateTestUser(db, "count1@example.com", 1)
	testutils.CreateTestUser(db, "count2@example.com", 1)

	// Get new count
	newCount, err := repo.Count()
	testutils.AssertNil(t, err)
	testutils.AssertNotNil(t, newCount)

	// Should have 2 more users
	expected := *initialCount + 2
	testutils.AssertEqual(t, expected, *newCount)
}

func TestGormRepository_FindByField(t *testing.T) {
	db := testutils.SetupTestDB(t)
	defer testutils.CleanupTestDB(db)
	testutils.SeedTestData(t, db)

	repo := NewRepository[models.User](db)

	// Create a test user
	user := testutils.CreateTestUser(db, "find@example.com", 1)

	// Find by email
	found, err := repo.FindByField("email", "find@example.com")
	testutils.AssertNil(t, err)
	testutils.AssertNotNil(t, found)
	testutils.AssertEqual(t, user.ID, found.ID)
	testutils.AssertEqual(t, user.Email, found.Email)

	// Find with preload
	found, err = repo.FindByField("email", "find@example.com", "Role")
	testutils.AssertNil(t, err)
	testutils.AssertNotNil(t, found)
	testutils.AssertNotNil(t, found.Role)

	// Find non-existent user
	_, err = repo.FindByField("email", "nonexistent@example.com")
	testutils.AssertNotNil(t, err)
	testutils.AssertEqual(t, 404, err.StatusCode)
}

func TestGormRepository_FindAllByField(t *testing.T) {
	db := testutils.SetupTestDB(t)
	defer testutils.CleanupTestDB(db)
	testutils.SeedTestData(t, db)

	repo := NewRepository[models.User](db)

	// Create multiple users with same role
	testutils.CreateTestUser(db, "role1@example.com", 1)
	testutils.CreateTestUser(db, "role2@example.com", 1)
	testutils.CreateTestUser(db, "role3@example.com", 2)

	// Find all users with role_id = 1
	users, err := repo.FindAllByField("role_id", int64(1), 0, 10)
	testutils.AssertNil(t, err)
	testutils.AssertNotNil(t, users)

	// Should have at least 3 users with role_id = 1 (including seeded admin)
	if len(*users) < 3 {
		t.Errorf("Expected at least 3 users with role_id=1, got %d", len(*users))
	}

	// Verify all users have role_id = 1
	for _, user := range *users {
		testutils.AssertEqual(t, int64(1), user.RoleID)
	}

	// Test with preload
	users, err = repo.FindAllByField("role_id", int64(1), 0, 5, "Role")
	testutils.AssertNil(t, err)
	testutils.AssertNotNil(t, users)

	// Check if Role is preloaded
	if len(*users) > 0 {
		user := (*users)[0]
		testutils.AssertNotNil(t, user.Role)
		testutils.AssertEqual(t, "admin", user.Role.Name)
	}
}

func TestGormRepository_Query(t *testing.T) {
	db := testutils.SetupTestDB(t)
	defer testutils.CleanupTestDB(db)
	testutils.SeedTestData(t, db)

	repo := NewRepository[models.User](db)

	// Test that Query returns a valid GORM DB instance
	query := repo.Query()
	testutils.AssertNotNil(t, query)

	// Test using the query directly
	var count int64
	err := query.Model(&models.User{}).Count(&count).Error
	testutils.AssertNoError(t, err)

	// Count should be greater than 0 (seeded data)
	if count == 0 {
		t.Error("Expected count to be greater than 0")
	}
}

// Test repository with different model types
func TestGormRepository_WithRoleModel(t *testing.T) {
	db := testutils.SetupTestDB(t)
	defer testutils.CleanupTestDB(db)
	testutils.SeedTestData(t, db)

	repo := NewRepository[models.Role](db)

	// Create a new role
	role := &models.Role{
		Name: "moderator",
	}

	err := repo.Create(role)
	testutils.AssertNil(t, err)
	testutils.AssertNotNil(t, role)

	// Get the created role
	retrieved, err := repo.GetByID(role.ID)
	testutils.AssertNil(t, err)
	testutils.AssertNotNil(t, retrieved)
	testutils.AssertEqual(t, "moderator", retrieved.Name)
}

func TestGormRepository_WithArtModel(t *testing.T) {
	db := testutils.SetupTestDB(t)
	defer testutils.CleanupTestDB(db)
	testutils.SeedTestData(t, db)

	repo := NewRepository[models.Art](db)

	// Create a new art piece
	art := &models.Art{
		Price:        50000, // 500.00 CHF in cents
		CurrencyID:   1,
		CreationYear: 2022,
		Available:    true,
		Visible:      true,
	}

	err := repo.Create(art)
	testutils.AssertNil(t, err)
	testutils.AssertNotNil(t, art)

	// Get the created art
	retrieved, err := repo.GetByID(art.ID)
	testutils.AssertNil(t, err)
	testutils.AssertNotNil(t, retrieved)
	testutils.AssertEqual(t, int64(50000), retrieved.Price)
	testutils.AssertEqual(t, 2022, retrieved.CreationYear)
	testutils.AssertEqual(t, true, retrieved.Available)
}