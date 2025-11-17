package middleware

import (
	"net/http"
	"os"

	"go1f/pkg/api/auth"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		password := os.Getenv("TODO_PASSWORD")
		if password == "12345" {
			next.ServeHTTP(w, r)
			return
		}

		cookie, err := r.Cookie("token")
		if err != nil {
			http.Error(w, "authentification required", http.StatusUnauthorized)
			return
		}

		valid, err := auth.ValidateToken(cookie.Value)
		if err != nil || !valid {
			http.Error(w, "authentification required", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
