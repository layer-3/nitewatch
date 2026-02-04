package auth

import (
	"regexp"
	"testing"
	"time"
)

func TestNewPINManager(t *testing.T) {
	manager := newPINManager()

	if manager == nil {
		t.Fatal("newPINManager returned nil")
	}
	if manager.expirationDuration != PINExpiration {
		t.Errorf("Expected expiration duration %v, got %v", PINExpiration, manager.expirationDuration)
	}
}

func TestGeneratePIN(t *testing.T) {
	manager := newPINManager()

	// Test generating multiple PINs
	generatedPINs := make(map[string]bool)
	pinRegex := regexp.MustCompile(`^\d{6}$`)

	for i := 0; i < 100; i++ {
		pin, err := manager.generatePIN()
		if err != nil {
			t.Fatalf("generatePIN failed: %v", err)
		}

		// Check PIN format (6 digits)
		if !pinRegex.MatchString(pin) {
			t.Errorf("PIN does not match expected format (6 digits): %s", pin)
		}

		// Check length
		if len(pin) != PINLength {
			t.Errorf("Expected PIN length %d, got %d", PINLength, len(pin))
		}

		generatedPINs[pin] = true
	}

	// Check that we got some variety in PINs (at least 50 unique PINs out of 100)
	if len(generatedPINs) < 50 {
		t.Errorf("Expected more variety in generated PINs, got only %d unique PINs out of 100", len(generatedPINs))
	}
}

func TestHashPIN(t *testing.T) {
	manager := newPINManager()

	tests := []struct {
		name string
		pin  string
	}{
		{"6 digit PIN", "123456"},
		{"PIN with leading zeros", "000123"},
		{"All zeros", "000000"},
		{"All nines", "999999"},
		{"Random PIN 1", "456789"},
		{"Random PIN 2", "987654"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hash1 := manager.hashPIN(tt.pin)
			hash2 := manager.hashPIN(tt.pin)

			// Same PIN should produce same hash
			if hash1 != hash2 {
				t.Errorf("Same PIN produced different hashes: %d vs %d", hash1, hash2)
			}

			// Hash should not be zero
			if hash1 == 0 {
				t.Error("Hash should not be zero")
			}
		})
	}

	// Different PINs should produce different hashes
	hash1 := manager.hashPIN("123456")
	hash2 := manager.hashPIN("123457")
	if hash1 == hash2 {
		t.Error("Different PINs produced same hash")
	}
}

func TestValidatePIN(t *testing.T) {
	manager := newPINManager()

	tests := []struct {
		name     string
		pin      string
		expected bool
	}{
		{"Valid PIN", "123456", true},
		{"Invalid PIN", "123457", false},
		{"Empty PIN", "", false},
		{"Short PIN", "12345", false},
		{"Long PIN", "1234567", false},
		{"PIN with letters", "12345a", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Generate hash for the correct PIN
			correctPIN := "123456"
			hashedPIN := manager.hashPIN(correctPIN)

			// Validate the test PIN
			result := manager.validatePIN(tt.pin, hashedPIN)

			if result != tt.expected {
				t.Errorf("Expected %v, got %v for PIN %s", tt.expected, result, tt.pin)
			}
		})
	}
}

func TestIsExpired(t *testing.T) {
	manager := newPINManager()

	tests := []struct {
		name      string
		createdAt time.Time
		isExpired bool
	}{
		{
			name:      "Just created",
			createdAt: time.Now(),
			isExpired: false,
		},
		{
			name:      "5 minutes old",
			createdAt: time.Now().Add(-5 * time.Minute),
			isExpired: false,
		},
		{
			name:      "9 minutes old",
			createdAt: time.Now().Add(-9 * time.Minute),
			isExpired: false,
		},
		{
			name:      "11 minutes old",
			createdAt: time.Now().Add(-11 * time.Minute),
			isExpired: true,
		},
		{
			name:      "1 hour old",
			createdAt: time.Now().Add(-1 * time.Hour),
			isExpired: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := manager.isExpired(tt.createdAt)
			if result != tt.isExpired {
				t.Errorf("Expected isExpired=%v, got %v", tt.isExpired, result)
			}
		})
	}
}

func TestValidatePINWithExpiry(t *testing.T) {
	manager := newPINManager()

	correctPIN := "123456"
	hashedPIN := manager.hashPIN(correctPIN)

	tests := []struct {
		name          string
		pin           string
		createdAt     time.Time
		expectedError error
	}{
		{
			name:          "Valid PIN, not expired",
			pin:           correctPIN,
			createdAt:     time.Now(),
			expectedError: nil,
		},
		{
			name:          "Valid PIN, almost expired",
			pin:           correctPIN,
			createdAt:     time.Now().Add(-9 * time.Minute),
			expectedError: nil,
		},
		{
			name:          "Valid PIN, expired",
			pin:           correctPIN,
			createdAt:     time.Now().Add(-11 * time.Minute),
			expectedError: ErrPINExpired,
		},
		{
			name:          "Invalid PIN, not expired",
			pin:           "654321",
			createdAt:     time.Now(),
			expectedError: ErrInvalidPIN,
		},
		{
			name:          "Invalid PIN, expired",
			pin:           "654321",
			createdAt:     time.Now().Add(-11 * time.Minute),
			expectedError: ErrPINExpired, // Expiry is checked first
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := manager.validatePINWithExpiry(tt.pin, hashedPIN, tt.createdAt)

			if err != tt.expectedError {
				t.Errorf("Expected error %v, got %v", tt.expectedError, err)
			}
		})
	}
}

func TestPINConsistency(t *testing.T) {
	manager := newPINManager()

	// Generate a PIN
	pin, err := manager.generatePIN()
	if err != nil {
		t.Fatalf("Failed to generate PIN: %v", err)
	}

	// Hash it
	hashedPIN := manager.hashPIN(pin)

	// Validate it immediately
	if !manager.validatePIN(pin, hashedPIN) {
		t.Error("Failed to validate a PIN that was just generated and hashed")
	}

	// Validate with expiry (should pass as it's fresh)
	err = manager.validatePINWithExpiry(pin, hashedPIN, time.Now())
	if err != nil {
		t.Errorf("Failed to validate fresh PIN with expiry: %v", err)
	}
}

func TestPINHashUniqueness(t *testing.T) {
	manager := newPINManager()

	// Test that different PINs produce different hashes
	pins := []string{"000000", "111111", "222222", "333333", "444444", "555555", "666666", "777777", "888888", "999999", "123456", "654321"}
	hashes := make(map[uint64]string)

	for _, pin := range pins {
		hash := manager.hashPIN(pin)
		if existingPIN, exists := hashes[hash]; exists {
			t.Errorf("Hash collision: PIN %s and %s produced same hash %d", pin, existingPIN, hash)
		}
		hashes[hash] = pin
	}
}

func TestExpirationDurationConstant(t *testing.T) {
	// Verify the constant is set correctly
	if PINExpiration != 10*time.Minute {
		t.Errorf("Expected PINExpiration to be 10 minutes, got %v", PINExpiration)
	}

	if PINLength != 6 {
		t.Errorf("Expected PINLength to be 6, got %d", PINLength)
	}
}

func TestConcurrentPINGeneration(t *testing.T) {
	manager := newPINManager()

	// Test concurrent PIN generation
	numGoroutines := 10
	pinsPerGoroutine := 10
	results := make(chan string, numGoroutines*pinsPerGoroutine)
	errors := make(chan error, numGoroutines*pinsPerGoroutine)

	for i := 0; i < numGoroutines; i++ {
		go func() {
			for j := 0; j < pinsPerGoroutine; j++ {
				pin, err := manager.generatePIN()
				if err != nil {
					errors <- err
				} else {
					results <- pin
				}
			}
		}()
	}

	// Collect results
	pinRegex := regexp.MustCompile(`^\d{6}$`)
	for i := 0; i < numGoroutines*pinsPerGoroutine; i++ {
		select {
		case err := <-errors:
			t.Errorf("Error generating PIN: %v", err)
		case pin := <-results:
			if !pinRegex.MatchString(pin) {
				t.Errorf("Invalid PIN format: %s", pin)
			}
		case <-time.After(5 * time.Second):
			t.Fatal("Timeout waiting for PIN generation")
		}
	}
}
