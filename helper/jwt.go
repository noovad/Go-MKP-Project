package helper

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Claims struct {
	Id       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	jwt.RegisteredClaims
}

func GenerateToken(userId uuid.UUID, username string) (string, error) {
	secret := os.Getenv("GENERATE_ACCESS_TOKEN_SECRET")
	if secret == "" {
		return "", errors.New("JWT secret not configured")
	}

	claims := Claims{
		Id:       userId,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func ValidateToken(tokenString string, secret string) (*Claims, bool) {
	if tokenString == "" || secret == "" {
		return nil, false
	}

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, false
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, true
	}

	return nil, false
}
