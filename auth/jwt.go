package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// TokenType represents the type of JWT token
type TokenType string

const (
	AccessTokenType  TokenType = "access"
	RefreshTokenType TokenType = "refresh"
)

// Claims represents the JWT claims structure
type Claims struct {
	Type  TokenType `json:"type"`
	Email string    `json:"email"`
	jwt.RegisteredClaims
}

// jwtManager handles JWT token operations
type jwtManager struct {
	secretKey       []byte
	accessDuration  time.Duration
	refreshDuration time.Duration
	issuer          string
}

// newJWTManager creates a new JWT manager instance
func newJWTManager(secretKey string, accessDuration, refreshDuration time.Duration) *jwtManager {
	return &jwtManager{
		secretKey:       []byte(secretKey),
		accessDuration:  accessDuration,
		refreshDuration: refreshDuration,
		issuer:          "nitewatch",
	}
}

// generateTokenPair generates both access and refresh tokens
func (j *jwtManager) generateTokenPair(email string, deviceID uuid.UUID) (accessToken, refreshToken string, err error) {
	// Generate access token
	accessToken, _, err = j.generateToken(email, deviceID, AccessTokenType, j.accessDuration)
	if err != nil {
		return "", "", err
	}

	// Generate refresh token
	refreshToken, _, err = j.generateToken(email, deviceID, RefreshTokenType, j.refreshDuration)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

// refreshTokens generates new tokens using the existing refresh token's subject
func (j *jwtManager) refreshTokens(email string, deviceID uuid.UUID) (accessToken, refreshToken string, err error) {
	// Generate new access token
	accessToken, _, err = j.generateToken(email, deviceID, AccessTokenType, j.accessDuration)
	if err != nil {
		return "", "", err
	}

	// Generate new refresh token
	refreshToken, _, err = j.generateToken(email, deviceID, RefreshTokenType, j.refreshDuration)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

// generateToken creates a JWT token
func (j *jwtManager) generateToken(email string, deviceID uuid.UUID, tokenType TokenType, duration time.Duration) (tokenString string, jti string, err error) {
	jti = uuid.New().String()
	tokenString = j.generateTokenWithJTI(email, deviceID, tokenType, duration, jti)
	return tokenString, jti, nil
}

// generateTokenWithJTI creates a JWT token with a specific JTI
func (j *jwtManager) generateTokenWithJTI(email string, deviceID uuid.UUID, tokenType TokenType, duration time.Duration, jti string) string {
	now := time.Now()
	claims := Claims{
		Type:  tokenType,
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    j.issuer,
			Subject:   deviceID.String(),
			ExpiresAt: jwt.NewNumericDate(now.Add(duration)),
			IssuedAt:  jwt.NewNumericDate(now),
			ID:        jti,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(j.secretKey)
	return tokenString
}

// validateToken validates a JWT token and returns the claims
func (j *jwtManager) validateToken(tokenString string, expectedType TokenType) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidToken
		}
		return j.secretKey, nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, ErrInvalidToken
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, ErrInvalidToken
	}

	// Verify token type
	if claims.Type != expectedType {
		return nil, ErrInvalidTokenType
	}

	// Check expiration
	if claims.ExpiresAt != nil && claims.ExpiresAt.Before(time.Now()) {
		return nil, ErrTokenExpired
	}

	return claims, nil
}

// extractJTI extracts the JTI from a token without full validation
func (j *jwtManager) extractJTI(tokenString string) (string, error) {
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, &Claims{})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return "", ErrInvalidToken
	}

	return claims.ID, nil
}
