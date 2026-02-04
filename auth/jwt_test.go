package auth

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func TestNewJWTManager(t *testing.T) {
	secretKey := "test-secret-key"
	accessDuration := 5 * time.Minute
	refreshDuration := 7 * 24 * time.Hour

	manager := newJWTManager(secretKey, accessDuration, refreshDuration)

	if manager == nil {
		t.Fatal("newJWTManager returned nil")
	}
	if string(manager.secretKey) != secretKey {
		t.Errorf("Expected secret key %s, got %s", secretKey, string(manager.secretKey))
	}
	if manager.accessDuration != accessDuration {
		t.Errorf("Expected access duration %v, got %v", accessDuration, manager.accessDuration)
	}
	if manager.refreshDuration != refreshDuration {
		t.Errorf("Expected refresh duration %v, got %v", refreshDuration, manager.refreshDuration)
	}
	if manager.issuer != "nitewatch" {
		t.Errorf("Expected issuer 'nitewatch', got %s", manager.issuer)
	}
}

func TestGenerateTokenPair(t *testing.T) {
	manager := newJWTManager("test-secret", 5*time.Minute, 24*time.Hour)
	email := "test@example.com"
	deviceID := uuid.New()

	accessToken, refreshToken, err := manager.generateTokenPair(email, deviceID)
	if err != nil {
		t.Fatalf("generateTokenPair failed: %v", err)
	}

	if accessToken == "" {
		t.Error("Access token is empty")
	}
	if refreshToken == "" {
		t.Error("Refresh token is empty")
	}

	// Validate access token
	accessClaims, err := manager.validateToken(accessToken, AccessTokenType)
	if err != nil {
		t.Fatalf("Failed to validate access token: %v", err)
	}
	if accessClaims.Email != email {
		t.Errorf("Expected email %s, got %s", email, accessClaims.Email)
	}
	if accessClaims.Subject != deviceID.String() {
		t.Errorf("Expected subject %s, got %s", deviceID.String(), accessClaims.Subject)
	}
	if accessClaims.Type != AccessTokenType {
		t.Errorf("Expected token type %s, got %s", AccessTokenType, accessClaims.Type)
	}

	// Validate refresh token
	refreshClaims, err := manager.validateToken(refreshToken, RefreshTokenType)
	if err != nil {
		t.Fatalf("Failed to validate refresh token: %v", err)
	}
	if refreshClaims.Email != email {
		t.Errorf("Expected email %s, got %s", email, refreshClaims.Email)
	}
	if refreshClaims.Subject != deviceID.String() {
		t.Errorf("Expected subject %s, got %s", deviceID.String(), refreshClaims.Subject)
	}
	if refreshClaims.Type != RefreshTokenType {
		t.Errorf("Expected token type %s, got %s", RefreshTokenType, refreshClaims.Type)
	}
}

func TestRefreshTokens(t *testing.T) {
	manager := newJWTManager("test-secret", 5*time.Minute, 24*time.Hour)
	email := "test@example.com"
	deviceID := uuid.New()

	accessToken, refreshToken, err := manager.refreshTokens(email, deviceID)
	if err != nil {
		t.Fatalf("refreshTokens failed: %v", err)
	}

	if accessToken == "" {
		t.Error("Access token is empty")
	}
	if refreshToken == "" {
		t.Error("Refresh token is empty")
	}

	// Validate that refresh token has the correct subject
	refreshClaims, err := manager.validateToken(refreshToken, RefreshTokenType)
	if err != nil {
		t.Fatalf("Failed to validate refresh token: %v", err)
	}
	if refreshClaims.Subject != deviceID.String() {
		t.Errorf("Expected subject %s, got %s", deviceID.String(), refreshClaims.Subject)
	}
}

func TestGenerateToken(t *testing.T) {
	manager := newJWTManager("test-secret", 5*time.Minute, 24*time.Hour)
	email := "test@example.com"
	deviceID := uuid.New()
	tokenType := AccessTokenType
	duration := 10 * time.Minute

	tokenString, jti, err := manager.generateToken(email, deviceID, tokenType, duration)
	if err != nil {
		t.Fatalf("generateToken failed: %v", err)
	}

	if tokenString == "" {
		t.Error("Token string is empty")
	}
	if jti == "" {
		t.Error("JTI is empty")
	}

	// Validate the generated token
	claims, err := manager.validateToken(tokenString, tokenType)
	if err != nil {
		t.Fatalf("Failed to validate token: %v", err)
	}
	if claims.Email != email {
		t.Errorf("Expected email %s, got %s", email, claims.Email)
	}
	if claims.Subject != deviceID.String() {
		t.Errorf("Expected subject %s, got %s", deviceID.String(), claims.Subject)
	}
	if claims.ID != jti {
		t.Errorf("Expected JTI %s, got %s", jti, claims.ID)
	}
}

func TestValidateToken(t *testing.T) {
	manager := newJWTManager("test-secret", 5*time.Minute, 24*time.Hour)
	email := "test@example.com"
	deviceID := uuid.New()

	tests := []struct {
		name          string
		tokenType     TokenType
		expectedType  TokenType
		duration      time.Duration
		expectError   bool
		expectedError error
	}{
		{
			name:         "Valid access token",
			tokenType:    AccessTokenType,
			expectedType: AccessTokenType,
			duration:     5 * time.Minute,
			expectError:  false,
		},
		{
			name:         "Valid refresh token",
			tokenType:    RefreshTokenType,
			expectedType: RefreshTokenType,
			duration:     24 * time.Hour,
			expectError:  false,
		},
		{
			name:          "Wrong token type",
			tokenType:     AccessTokenType,
			expectedType:  RefreshTokenType,
			duration:      5 * time.Minute,
			expectError:   true,
			expectedError: ErrInvalidTokenType,
		},
		{
			name:          "Expired token",
			tokenType:     AccessTokenType,
			expectedType:  AccessTokenType,
			duration:      -1 * time.Hour,
			expectError:   true,
			expectedError: nil, // jwt library returns its own error
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token, _, _ := manager.generateToken(email, deviceID, tt.tokenType, tt.duration)

			claims, err := manager.validateToken(token, tt.expectedType)

			if tt.expectError {
				if err == nil {
					t.Error("Expected error but got none")
				}
				if tt.expectedError != nil && err != tt.expectedError {
					t.Errorf("Expected error %v, got %v", tt.expectedError, err)
				}
			} else {
				if err != nil {
					t.Fatalf("Unexpected error: %v", err)
				}
				if claims.Email != email {
					t.Errorf("Expected email %s, got %s", email, claims.Email)
				}
			}
		})
	}
}

func TestValidateTokenWithInvalidSignature(t *testing.T) {
	manager1 := newJWTManager("secret1", 5*time.Minute, 24*time.Hour)
	manager2 := newJWTManager("secret2", 5*time.Minute, 24*time.Hour)
	deviceID := uuid.New()

	// Generate token with manager1
	token, _, _ := manager1.generateToken("test@example.com", deviceID, AccessTokenType, 5*time.Minute)

	// Try to validate with manager2 (different secret)
	_, err := manager2.validateToken(token, AccessTokenType)
	if err == nil {
		t.Error("Expected error for invalid signature, got none")
	}
}

func TestExtractJTI(t *testing.T) {
	manager := newJWTManager("test-secret", 5*time.Minute, 24*time.Hour)
	email := "test@example.com"
	deviceID := uuid.New()
	expectedJTI := "test-jti-456"

	token := manager.generateTokenWithJTI(email, deviceID, AccessTokenType, 5*time.Minute, expectedJTI)

	extractedJTI, err := manager.extractJTI(token)
	if err != nil {
		t.Fatalf("extractJTI failed: %v", err)
	}

	if extractedJTI != expectedJTI {
		t.Errorf("Expected JTI %s, got %s", expectedJTI, extractedJTI)
	}
}

func TestExtractJTIFromInvalidToken(t *testing.T) {
	manager := newJWTManager("test-secret", 5*time.Minute, 24*time.Hour)

	// Test with invalid token
	_, err := manager.extractJTI("invalid-token")
	if err == nil {
		t.Error("Expected error for invalid token, got none")
	}
}

func TestGenerateTokenWithJTI(t *testing.T) {
	manager := newJWTManager("test-secret", 5*time.Minute, 24*time.Hour)
	email := "test@example.com"
	deviceID := uuid.New()
	tokenType := AccessTokenType
	duration := 10 * time.Minute
	specificJTI := "specific-jti-789"

	tokenString := manager.generateTokenWithJTI(email, deviceID, tokenType, duration, specificJTI)

	if tokenString == "" {
		t.Error("Token string is empty")
	}

	// Validate and check JTI
	claims, err := manager.validateToken(tokenString, tokenType)
	if err != nil {
		t.Fatalf("Failed to validate token: %v", err)
	}

	if claims.ID != specificJTI {
		t.Errorf("Expected JTI %s, got %s", specificJTI, claims.ID)
	}
}

func TestTokenExpiration(t *testing.T) {
	manager := newJWTManager("test-secret", 5*time.Minute, 24*time.Hour)
	email := "test@example.com"
	deviceID := uuid.New()

	// Test 1: Create already expired token (negative duration)
	expiredToken, _, _ := manager.generateToken(email, deviceID, AccessTokenType, -1*time.Hour)
	_, err := manager.validateToken(expiredToken, AccessTokenType)
	if err == nil {
		t.Error("Expected error for expired token, got none")
	}

	// Test 2: Create token that expires in the past by manipulating time in claims
	now := time.Now()
	claims := Claims{
		Type:  AccessTokenType,
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    manager.issuer,
			Subject:   deviceID.String(),
			ExpiresAt: jwt.NewNumericDate(now.Add(-1 * time.Hour)), // 1 hour ago
			IssuedAt:  jwt.NewNumericDate(now.Add(-2 * time.Hour)), // 2 hours ago
			ID:        "test-jti",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(manager.secretKey)

	_, err = manager.validateToken(tokenString, AccessTokenType)
	if err == nil {
		t.Error("Expected error for manually created expired token, got none")
	}

	// Test 3: Valid token (future expiration)
	validToken, _, _ := manager.generateToken(email, deviceID, AccessTokenType, 1*time.Hour)
	_, err = manager.validateToken(validToken, AccessTokenType)
	if err != nil {
		t.Errorf("Expected valid token, got error: %v", err)
	}
}

func TestClaimsStructure(t *testing.T) {
	manager := newJWTManager("test-secret", 5*time.Minute, 24*time.Hour)
	email := "test@example.com"
	deviceID := uuid.New()

	token, jti, _ := manager.generateToken(email, deviceID, AccessTokenType, 5*time.Minute)

	// Parse token to check claims
	parsedToken, _ := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return manager.secretKey, nil
	})

	claims := parsedToken.Claims.(*Claims)

	// Check all claim fields
	if claims.Type != AccessTokenType {
		t.Errorf("Expected token type %s, got %s", AccessTokenType, claims.Type)
	}
	if claims.Issuer != "nitewatch" {
		t.Errorf("Expected issuer 'nitewatch', got %s", claims.Issuer)
	}
	if claims.Email != email {
		t.Errorf("Expected email %s, got %s", email, claims.Email)
	}
	if claims.Subject != deviceID.String() {
		t.Errorf("Expected subject %s, got %s", deviceID.String(), claims.Subject)
	}
	if claims.ID != jti {
		t.Errorf("Expected JTI %s, got %s", jti, claims.ID)
	}
	if claims.IssuedAt == nil {
		t.Error("IssuedAt is nil")
	}
	if claims.ExpiresAt == nil {
		t.Error("ExpiresAt is nil")
	}
}
