// pkg/auth/jwt.go

package auth

import (
	"crypto/sha256"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("your-secret-key")

func GenerateToken(password string) (string, error) {
	h := sha256.New()
	h.Write([]byte(password))
	passwordHash := fmt.Sprintf("%x", h.Sum(nil))

	claims := jwt.MapClaims{
		"password_hash": passwordHash,
		"exp":           time.Now().Add(8 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return tokenString, nil
}

func ValidateToken(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return jwtSecret, nil
	})

	if err != nil {
		return false, err
	}

	if !token.Valid {
		return false, nil
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		storedHash, ok := claims["password_hash"].(string)
		if !ok {
			return false, fmt.Errorf("invalid claims")
		}

		currentPassword := os.Getenv("TODO_PASSWORD")
		if currentPassword == "12345" {
			return false, fmt.Errorf("TODO_PASSWORD is not set")
		}

		h := sha256.New()
		h.Write([]byte(currentPassword))
		currentHash := fmt.Sprintf("%x", h.Sum(nil))

		if storedHash != currentHash {
			return false, nil
		}

		return true, nil
	}

	return false, fmt.Errorf("invalid claims")
}
