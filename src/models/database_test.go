package models

import (
	"testing"
	"time"
)

func TestOrderStatus(t *testing.T) {
	tests := []struct {
		status   OrderStatus
		expected string
	}{
		{OrderStatusPending, "pending"},
		{OrderStatusPaid, "paid"},
		{OrderStatusShipped, "shipped"},
		{OrderStatusDelivered, "delivered"},
		{OrderStatusCancelled, "cancelled"},
	}

	for _, test := range tests {
		if string(test.status) != test.expected {
			t.Errorf("Expected %s, got %s", test.expected, string(test.status))
		}
	}
}

func TestSoftDeleteModel(t *testing.T) {
	model := SoftDeleteModel{
		ID:        1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if model.ID != 1 {
		t.Errorf("Expected ID 1, got %d", model.ID)
	}

	if model.DeletedAt.Valid {
		t.Error("Expected DeletedAt to be invalid (nil), but it was valid")
	}
}

func TestModel(t *testing.T) {
	now := time.Now()
	model := Model{
		ID:        1,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if model.ID != 1 {
		t.Errorf("Expected ID 1, got %d", model.ID)
	}

	if !model.CreatedAt.Equal(now) {
		t.Errorf("Expected CreatedAt %v, got %v", now, model.CreatedAt)
	}

	if !model.UpdatedAt.Equal(now) {
		t.Errorf("Expected UpdatedAt %v, got %v", now, model.UpdatedAt)
	}
}

func TestUserModel(t *testing.T) {
	user := User{
		SoftDeleteModel: SoftDeleteModel{ID: 1},
		Firstname:       "John",
		Lastname:        "Doe",
		Email:          "john@example.com",
		Password:       []byte("hashed"),
		Salt:          []byte("salt"),
		RoleID:        1,
	}

	if user.ID != 1 {
		t.Errorf("Expected ID 1, got %d", user.ID)
	}

	if user.Firstname != "John" {
		t.Errorf("Expected firstname 'John', got '%s'", user.Firstname)
	}

	if user.Email != "john@example.com" {
		t.Errorf("Expected email 'john@example.com', got '%s'", user.Email)
	}
}

func TestArtModel(t *testing.T) {
	width := float32(100.5)
	height := float32(150.0)
	depth := float32(5.5)

	art := Art{
		SoftDeleteModel: SoftDeleteModel{ID: 1},
		Price:          10000, // 100.00 CHF in cents
		CurrencyID:     1,
		CreationYear:   2023,
		Width:          &width,
		Height:         &height,
		Depth:          &depth,
		Available:      true,
		Visible:        true,
	}

	if art.ID != 1 {
		t.Errorf("Expected ID 1, got %d", art.ID)
	}

	if art.Price != 10000 {
		t.Errorf("Expected price 10000, got %d", art.Price)
	}

	if art.CreationYear != 2023 {
		t.Errorf("Expected creation year 2023, got %d", art.CreationYear)
	}

	if art.Width == nil || *art.Width != 100.5 {
		t.Errorf("Expected width 100.5, got %v", art.Width)
	}

	if !art.Available {
		t.Error("Expected art to be available")
	}

	if !art.Visible {
		t.Error("Expected art to be visible")
	}
}

func TestRoleModel(t *testing.T) {
	role := Role{
		Model: Model{ID: 1},
		Name:  "admin",
	}

	if role.ID != 1 {
		t.Errorf("Expected ID 1, got %d", role.ID)
	}

	if role.Name != "admin" {
		t.Errorf("Expected name 'admin', got '%s'", role.Name)
	}
}

func TestLanguageModel(t *testing.T) {
	language := Language{
		Model:        Model{ID: 1},
		LanguageName: "English",
		LanguageCode: "en",
	}

	if language.ID != 1 {
		t.Errorf("Expected ID 1, got %d", language.ID)
	}

	if language.LanguageName != "English" {
		t.Errorf("Expected language name 'English', got '%s'", language.LanguageName)
	}

	if language.LanguageCode != "en" {
		t.Errorf("Expected language code 'en', got '%s'", language.LanguageCode)
	}
}

func TestSavedModel(t *testing.T) {
	saved := Saved{
		Model:  Model{ID: 1},
		UserID: 1,
		ArtID:  1,
	}

	if saved.ID != 1 {
		t.Errorf("Expected ID 1, got %d", saved.ID)
	}

	if saved.UserID != 1 {
		t.Errorf("Expected UserID 1, got %d", saved.UserID)
	}

	if saved.ArtID != 1 {
		t.Errorf("Expected ArtID 1, got %d", saved.ArtID)
	}
}

func TestOrderModel(t *testing.T) {
	orderDate := time.Now()
	order := Order{
		Model:     Model{ID: 1},
		UserID:    1,
		ArtID:     1,
		OrderDate: orderDate,
		Status:    OrderStatusPending,
	}

	if order.ID != 1 {
		t.Errorf("Expected ID 1, got %d", order.ID)
	}

	if order.UserID != 1 {
		t.Errorf("Expected UserID 1, got %d", order.UserID)
	}

	if order.ArtID != 1 {
		t.Errorf("Expected ArtID 1, got %d", order.ArtID)
	}

	if order.Status != OrderStatusPending {
		t.Errorf("Expected status 'pending', got '%s'", order.Status)
	}

	if !order.OrderDate.Equal(orderDate) {
		t.Errorf("Expected order date %v, got %v", orderDate, order.OrderDate)
	}
}

func TestArtTranslationModel(t *testing.T) {
	translation := ArtTranslation{
		Model:       Model{ID: 1},
		ArtID:       1,
		LanguageID:  1,
		Title:       "Beautiful Painting",
		Description: "A beautiful painting description",
		Text:        "Detailed text about the artwork",
	}

	if translation.ID != 1 {
		t.Errorf("Expected ID 1, got %d", translation.ID)
	}

	if translation.ArtID != 1 {
		t.Errorf("Expected ArtID 1, got %d", translation.ArtID)
	}

	if translation.LanguageID != 1 {
		t.Errorf("Expected LanguageID 1, got %d", translation.LanguageID)
	}

	if translation.Title != "Beautiful Painting" {
		t.Errorf("Expected title 'Beautiful Painting', got '%s'", translation.Title)
	}
}

func TestArtPictureModel(t *testing.T) {
	priority := 1
	artPicture := ArtPicture{
		Model:     Model{ID: 1},
		ArtID:     1,
		PictureID: 1,
		Name:      "main-image.jpg",
		Priority:  &priority,
	}

	if artPicture.ID != 1 {
		t.Errorf("Expected ID 1, got %d", artPicture.ID)
	}

	if artPicture.ArtID != 1 {
		t.Errorf("Expected ArtID 1, got %d", artPicture.ArtID)
	}

	if artPicture.PictureID != 1 {
		t.Errorf("Expected PictureID 1, got %d", artPicture.PictureID)
	}

	if artPicture.Name != "main-image.jpg" {
		t.Errorf("Expected name 'main-image.jpg', got '%s'", artPicture.Name)
	}

	if artPicture.Priority == nil || *artPicture.Priority != 1 {
		t.Errorf("Expected priority 1, got %v", artPicture.Priority)
	}
}

func TestPictureModel(t *testing.T) {
	priority := int64(1)
	picture := Picture{
		Model:    Model{ID: 1},
		Name:     "image.jpg",
		Priority: &priority,
		IsPublic: true,
		Type:     "image/jpeg",
	}

	if picture.ID != 1 {
		t.Errorf("Expected ID 1, got %d", picture.ID)
	}

	if picture.Name != "image.jpg" {
		t.Errorf("Expected name 'image.jpg', got '%s'", picture.Name)
	}

	if picture.Priority == nil || *picture.Priority != 1 {
		t.Errorf("Expected priority 1, got %v", picture.Priority)
	}

	if !picture.IsPublic {
		t.Error("Expected picture to be public")
	}

	if picture.Type != "image/jpeg" {
		t.Errorf("Expected type 'image/jpeg', got '%s'", picture.Type)
	}
}

func TestCurrencyModel(t *testing.T) {
	currency := Currency{
		Model:        Model{ID: 1},
		CurrencyCode: "CHF",
		CurrencyName: "Swiss Franc",
	}

	if currency.ID != 1 {
		t.Errorf("Expected ID 1, got %d", currency.ID)
	}

	if currency.CurrencyCode != "CHF" {
		t.Errorf("Expected currency code 'CHF', got '%s'", currency.CurrencyCode)
	}

	if currency.CurrencyName != "Swiss Franc" {
		t.Errorf("Expected currency name 'Swiss Franc', got '%s'", currency.CurrencyName)
	}
}

func TestTranslationModel(t *testing.T) {
	translation := Translation{
		Model:      Model{ID: 1},
		EntityID:   1,
		LanguageID: 1,
		Context:    "button",
		Text:       "Click me",
	}

	if translation.ID != 1 {
		t.Errorf("Expected ID 1, got %d", translation.ID)
	}

	if translation.EntityID != 1 {
		t.Errorf("Expected EntityID 1, got %d", translation.EntityID)
	}

	if translation.LanguageID != 1 {
		t.Errorf("Expected LanguageID 1, got %d", translation.LanguageID)
	}

	if translation.Context != "button" {
		t.Errorf("Expected context 'button', got '%s'", translation.Context)
	}

	if translation.Text != "Click me" {
		t.Errorf("Expected text 'Click me', got '%s'", translation.Text)
	}
}