package auth

import (
	"testing"
)

func TestGenerateToken(t *testing.T) {
	service := NewService()
	token, err := service.GenerateToken(1)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if token == "" {
		t.Fatal("Expected token to not be empty")
	}
}
