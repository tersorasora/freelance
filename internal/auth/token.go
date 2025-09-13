package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var token_key = []byte("ForRealThisIsASecretKeyVerySecret")

func GenerateToken(userID string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // expires in 24h
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(token_key)
}

func ParseToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		// Check if token's signing method is HMAC
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return token_key, nil
	})
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if userID, ok := claims["user_id"].(string); ok {
			return userID, nil
		}
		return "", errors.New("user_id not found in token")
	}

	return "", errors.New("invalid token")
}