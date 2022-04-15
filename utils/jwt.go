package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

const (
	AccessTokenDuration   = time.Hour * 2
	AccessTokenSecretKey  = "access"
	RefreshTokenDuration  = time.Hour * 60 * 14
	RefreshTokenSecretKey = "refresh"
)

func generateToken(userId int64, secretKey string, duration time.Duration) (string, int64, error) {
	exp := time.Now().Add(duration).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userId,
		"exp":     exp,
	})
	t, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", 0, err
	}
	return t, exp, nil
}

func GenerateAccessToken(userId int64) (string, int64, error) {
	return generateToken(userId, AccessTokenSecretKey, AccessTokenDuration)
}

func GenerateRefreshToken(userId int64) (string, int64, error) {
	return generateToken(userId, RefreshTokenSecretKey, RefreshTokenDuration)
}

func GetUserIdFromToken(token *jwt.Token) int64 {
	claims := token.Claims.(jwt.MapClaims)
	return int64(claims["user_id"].(float64))
}
