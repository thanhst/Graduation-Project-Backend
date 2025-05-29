package helper

import (
	"log"
	"net/http"
	"time"
)

func SetTokenCookies(w http.ResponseWriter, accessToken, refreshToken string, accessExpire string, refreshExpire string) {
	accExpireTime, err := time.Parse(time.RFC3339, accessExpire)
	if err != nil {
		log.Printf("Invalid accessExpire format: %v\n", err)
		accExpireTime = time.Now().Add(1 * time.Hour)
	}
	refExpireTime, err := time.Parse(time.RFC3339, refreshExpire)
	if err != nil {
		log.Printf("Invalid refreshExpire format: %v\n", err)
		refExpireTime = time.Now().Add(7 * 24 * time.Hour)
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    accessToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
		Expires:  accExpireTime,
		MaxAge:   int(time.Until(accExpireTime).Seconds()),
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Path:     "/api/auth",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
		Expires:  refExpireTime,
		MaxAge:   int(time.Until(refExpireTime).Seconds()),
	})
}

func RemoveCookies(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
		Expires:  time.Unix(0, 0),
	})
	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    "",
		Path:     "/api/auth",
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
		Expires:  time.Unix(0, 0),
	})
}
