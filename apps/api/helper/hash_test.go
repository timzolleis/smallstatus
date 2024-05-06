package helper

import (
	"testing"
)

// TestHashPassword checks if the HashPassword function successfully generates hashes
func TestHashPassword(t *testing.T) {
	tests := []struct {
		name     string
		password string
		wantErr  bool
	}{
		{"Valid Password", "securePassword123!", false},
		{"Empty Password", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := HashPassword(tt.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("HashPassword() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// TestCheckPassword verifies the correctness of password checks using known passwords and hashes
func TestCheckPassword(t *testing.T) {
	// Test passwords
	originalPassword := "securePassword123!"
	invalidPassword := "wrongPassword"

	// Hash the original password
	hashedPassword, err := HashPassword(originalPassword)
	if err != nil {
		t.Fatalf("HashPassword() failed: %v", err)
	}

	tests := []struct {
		name     string
		password string
		hash     string
		want     bool
	}{
		{"Matching Password", originalPassword, hashedPassword, true},
		{"Non-Matching Password", invalidPassword, hashedPassword, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckPassword(tt.password, tt.hash); got != tt.want {
				t.Errorf("CheckPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}
