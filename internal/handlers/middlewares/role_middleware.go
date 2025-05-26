package middleware

import (
	"net/http"
	service "server/internal/handlers/services"
)

func RoleSwitchMiddleware(accountService *service.AccountService, handlers map[string]http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId, ok := r.Context().Value(UserIDKey).(string)
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		accounts, err := accountService.GetAccountsByUser(userId)
		if err != nil {
			http.Error(w, "Forbidden: role not allowed", http.StatusForbidden)
			return
		}
		if len(accounts) == 0 {
			http.Error(w, "No account found for user", http.StatusForbidden)
			return
		}
		role := accounts[0].Role
		handler, exists := handlers[role]
		if !exists {
			http.Error(w, "Forbidden: role not allowed", http.StatusForbidden)
			return
		}
		handler(w, r)
	}
}
