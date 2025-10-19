package models

import (
	"testing"
	"time"
)

func TestUserModel(t *testing.T) {
	user := User{
		Firstname:   "John",
		Lastname:    "Doe",
		Email:       "john.doe@example.com",
		Phone:       stringPtr("+41796123456"),
		PhoneRegion: stringPtr("CH"),
		Address1:    stringPtr("123 Main St"),
		Address2:    stringPtr("Apt 4B"),
		City:        stringPtr("Zurich"),
		PostalCode:  stringPtr("8001"),
		Password:    []byte("hashed_password"),
		Salt:        []byte("salt"),
		RoleID:      1,
	}

	if user.Firstname != "John" {
		t.Errorf("Expected firstname 'John', got '%s'", user.Firstname)
	}

	if user.Email != "john.doe@example.com" {
		t.Errorf("Expected email 'john.doe@example.com', got '%s'", user.Email)
	}

	if user.Phone == nil || *user.Phone != "+41796123456" {
		t.Errorf("Expected phone '+41796123456', got %v", user.Phone)
	}

	if user.RoleID != 1 {
		t.Errorf("Expected RoleID 1, got %d", user.RoleID)
	}
}

func TestUserModelWithNilOptionalFields(t *testing.T) {
	user := User{
		Firstname: "Jane",
		Lastname:  "Smith",
		Email:     "jane.smith@example.com",
		Password:  []byte("hashed_password"),
		Salt:      []byte("salt"),
		RoleID:    2,
	}

	if user.Phone != nil {
		t.Errorf("Expected phone to be nil, got %v", user.Phone)
	}

	if user.Address1 != nil {
		t.Errorf("Expected address1 to be nil, got %v", user.Address1)
	}

	if user.City != nil {
		t.Errorf("Expected city to be nil, got %v", user.City)
	}
}

func TestCreateUserRequest(t *testing.T) {
	request := CreateUserRequest{
		Firstname:   "Test",
		Lastname:    "User",
		Email:       "test@example.com",
		Phone:       stringPtr("+41791234567"),
		PhoneRegion: stringPtr("CH"),
		Password:    "SecurePassword123!",
	}

	if request.Firstname != "Test" {
		t.Errorf("Expected firstname 'Test', got '%s'", request.Firstname)
	}

	if request.Email != "test@example.com" {
		t.Errorf("Expected email 'test@example.com', got '%s'", request.Email)
	}

	if request.Password != "SecurePassword123!" {
		t.Errorf("Expected password 'SecurePassword123!', got '%s'", request.Password)
	}
}

func TestUpdateUserRequest(t *testing.T) {
	firstname := "Updated"
	email := "updated@example.com"
	phone := "+41799876543"

	request := UpdateUserRequest{
		UserID:    1,
		UserRole:  "admin",
		TargetID:  2,
		Firstname: &firstname,
		Email:     &email,
		Phone:     &phone,
	}

	if request.UserID != 1 {
		t.Errorf("Expected UserID 1, got %d", request.UserID)
	}

	if request.UserRole != "admin" {
		t.Errorf("Expected UserRole 'admin', got '%s'", request.UserRole)
	}

	if request.Firstname == nil || *request.Firstname != "Updated" {
		t.Errorf("Expected firstname 'Updated', got %v", request.Firstname)
	}

	if request.Email == nil || *request.Email != "updated@example.com" {
		t.Errorf("Expected email 'updated@example.com', got %v", request.Email)
	}
}

func TestGetUserByIDRequest(t *testing.T) {
	request := GetUserByIDRequest{
		UserID:   1,
		UserRole: "user",
		TargetID: 2,
	}

	if request.UserID != 1 {
		t.Errorf("Expected UserID 1, got %d", request.UserID)
	}

	if request.UserRole != "user" {
		t.Errorf("Expected UserRole 'user', got '%s'", request.UserRole)
	}

	if request.TargetID != 2 {
		t.Errorf("Expected TargetID 2, got %d", request.TargetID)
	}
}

func TestGetUserByEmailRequest(t *testing.T) {
	request := GetUserByEmailRequest{
		UserID:   1,
		UserRole: "admin",
		Email:    "search@example.com",
	}

	if request.UserID != 1 {
		t.Errorf("Expected UserID 1, got %d", request.UserID)
	}

	if request.UserRole != "admin" {
		t.Errorf("Expected UserRole 'admin', got '%s'", request.UserRole)
	}

	if request.Email != "search@example.com" {
		t.Errorf("Expected Email 'search@example.com', got '%s'", request.Email)
	}
}

func TestListUserRequest(t *testing.T) {
	request := ListUserRequest{
		UserID:   1,
		UserRole: "admin",
		Offset:   5,
	}

	if request.UserID != 1 {
		t.Errorf("Expected UserID 1, got %d", request.UserID)
	}

	if request.UserRole != "admin" {
		t.Errorf("Expected UserRole 'admin', got '%s'", request.UserRole)
	}

	if request.Offset != 5 {
		t.Errorf("Expected Offset 5, got %d", request.Offset)
	}
}

func TestDeleteUserRequest(t *testing.T) {
	request := DeleteUserRequest{
		UserID:   1,
		UserRole: "admin",
		TargetID: 3,
	}

	if request.UserID != 1 {
		t.Errorf("Expected UserID 1, got %d", request.UserID)
	}

	if request.UserRole != "admin" {
		t.Errorf("Expected UserRole 'admin', got '%s'", request.UserRole)
	}

	if request.TargetID != 3 {
		t.Errorf("Expected TargetID 3, got %d", request.TargetID)
	}
}

func TestUserWithLastAccess(t *testing.T) {
	now := time.Now()
	user := User{
		Firstname:  "Test",
		Lastname:   "User",
		Email:      "test@example.com",
		Password:   []byte("hashed"),
		Salt:       []byte("salt"),
		RoleID:     1,
		LastAccess: &now,
	}

	if user.LastAccess == nil {
		t.Error("Expected LastAccess to be set, got nil")
	}

	if !user.LastAccess.Equal(now) {
		t.Errorf("Expected LastAccess to be %v, got %v", now, *user.LastAccess)
	}
}

// Helper function to create string pointers
func stringPtr(s string) *string {
	return &s
}