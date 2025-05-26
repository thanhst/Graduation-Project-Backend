package jwtutil

import (
	"errors"
	"server/internal/handlers/dto"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type GitHubUserClaims struct {
	RedirectURI string `json:"redirect_uri"`
	Nonce       string `json:"nonce"`
	ExpiresAt   int64  `json:"expires_at"`
	jwt.RegisteredClaims
}

func CreateGitHubStateToken(githubRequest *dto.GitHubUserRequest) (string, error) {
	exp := time.Now().Add(time.Duration(githubRequest.ExpiresAt))
	claims := GitHubUserClaims{
		RedirectURI: githubRequest.RedirectURI,
		Nonce:       githubRequest.Nonce,
		ExpiresAt:   exp.Unix(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(GetAccessToken())
}

func VerifyAccessTokenGithub(tokenStr string) (*GitHubUserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &GitHubUserClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(GetAccessToken()), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*GitHubUserClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token or claims")
}
