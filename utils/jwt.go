package utils

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// replace with env secret in production
var jwtSecret = []byte("2c476f896200489ebcee594ad481cd38")

// CustomClaims stores user id and standard claims
type CustomClaims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

// GenerateToken creates a JWT for a user ID
func GenerateToken(userID uint) (string, error) {
	claims := CustomClaims{UserID: userID, RegisteredClaims: jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(time.Now().Add(72 * time.Hour))}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// ParseToken validates token and returns claims
func ParseToken(tokenStr string) (*CustomClaims, error) {
	var claims CustomClaims
	parsed, err := jwt.ParseWithClaims(tokenStr, &claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil || !parsed.Valid {
		return nil, errors.New("invalid token")
	}
	return &claims, nil
}

// ExtractTokenFromHeader removes "Bearer " prefix
func ExtractTokenFromHeader(h string) string {
	// naive extraction; expects "Bearer <token>"
	if len(h) > 7 && h[:7] == "Bearer " {
		return h[7:]
	}
	return h
}

// GenerateTokenString returns a URL-safe random token for share links
func GenerateTokenString() string {
	b := make([]byte, 32)
	_, _ = rand.Read(b)
	return base64.RawURLEncoding.EncodeToString(b)
}
