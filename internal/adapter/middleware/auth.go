package middleware

import (
	"context"
	"strings"
	"net/http"
	"github.com/MaryneZa/backend-challenge/internal/core/util"
)

func NewAuthMiddleware(secret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				util.SendErrorResponse(w, "Missing access token", http.StatusUnauthorized)
				return
			}

			tokenParts := strings.Split(authHeader, " ")
			if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
				util.SendErrorResponse(w, "Invalid token format", http.StatusUnauthorized)
				return
			}
			accessToken := tokenParts[1]

			userID, err := util.VerifyToken(accessToken, secret)
			if err != nil {
				util.SendErrorResponse(w, "Invalid or expired token", http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), "userID", userID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
