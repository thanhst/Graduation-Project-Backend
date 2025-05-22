package jwtutil

import (
	"errors"
	"fmt"
	"log"
	"server/internal/utils/dotenv"
	CustomHash "server/internal/utils/hash"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var (
	accessSecret  []byte
	refreshSecret []byte
	accessExpire  time.Duration
	refreshExpire time.Duration
)

func init() {
	accessSecret = CustomHash.GenerateSecretKey(dotenv.GetDotEnv("JWT_ACCESS_SECRET"))
	refreshSecret = CustomHash.GenerateSecretKey(dotenv.GetDotEnv("JWT_REFRESH_SECRET"))

	if len(accessSecret) == 0 {
		panic("JWT_ACCESS_SECRET is not set")
	}
	if len(refreshSecret) == 0 {
		panic("JWT_REFRESH_SECRET is not set")
	}
	accessExpireHoursStr := dotenv.GetDotEnv("JWT_ACCESS_EXPIRE_HOURS")
	accessHours, err := strconv.Atoi(accessExpireHoursStr)
	if err != nil {
		log.Fatalf("Invalid JWT_ACCESS_EXPIRE_HOURS: %v", err)
	}
	accessExpire = time.Duration(accessHours) * time.Hour

	// Parse refresh expire days
	refreshExpireDaysStr := dotenv.GetDotEnv("JWT_REFRESH_EXPIRE_DAYS")
	refreshDays, err := strconv.Atoi(refreshExpireDaysStr)
	if err != nil {
		log.Fatalf("Invalid JWT_REFRESH_EXPIRE_DAYS: %v", err)
	}
	refreshExpire = time.Duration(refreshDays) * 24 * time.Hour
}

type CustomClaims struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func GetAccessToken() []byte {
	return accessSecret
}
func GetRefreshToken() []byte {
	return refreshSecret
}

func GetAccessExpire() time.Duration {
	return accessExpire
}
func GetRefreshExpire() time.Duration {
	return refreshExpire
}

func CreateToken(userID string, secret []byte, expire time.Duration) (string, error) {
	claims := CustomClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expire)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}

func RefreshTokens(refreshToken string) (map[string]string, error) {
	token, err := jwt.ParseWithClaims(refreshToken, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return refreshSecret, nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid refresh token")
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok || claims.ExpiresAt.Time.Before(time.Now()) {
		return nil, errors.New("refresh token expired")
	}

	accessToken, err := CreateToken(claims.UserID, accessSecret, accessExpire)
	if err != nil {
		return nil, err
	}

	newRefreshToken, err := CreateToken(claims.UserID, refreshSecret, refreshExpire)
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"access_token":    accessToken,
		"refresh_token":   newRefreshToken,
		"access_exprise":  time.Now().Add(GetAccessExpire()).String(),
		"refresh_exprise": time.Now().Add(GetRefreshExpire()).String(),
	}, nil
}

func VerifyAccessToken(tokenStr string) (*CustomClaims, error) {
	return verifyToken(tokenStr, accessSecret)
}

func VerifyRefreshToken(tokenStr string) (*CustomClaims, error) {
	return verifyToken(tokenStr, refreshSecret)
}

func verifyToken(tokenStr string, secret []byte) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
