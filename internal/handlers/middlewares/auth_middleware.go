package middleware

import (
	"context"
	"net/http"
	jwtutil "server/internal/utils/jwt"
)

type contextKey string

const UserIDKey = contextKey("userId")

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("access_token")
		if err != nil {
			http.Error(w, "Missing access token cookie", http.StatusUnauthorized)
			return
		}

		tokenStr := cookie.Value

		claims, err := jwtutil.VerifyAccessToken(tokenStr)
		if err != nil {
			http.Error(w, "Invalid token: "+err.Error(), http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), UserIDKey, claims.UserID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
