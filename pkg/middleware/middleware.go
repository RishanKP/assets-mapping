package middleware

import (
	"errors"
	"net/http"
	"strings"

	"asset-mapping/library/api"
	"asset-mapping/library/jwt"
	"asset-mapping/library/utils"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if utils.IsEmpty(authHeader) {
			api.NewError(w, http.StatusUnauthorized, errors.New("missing authorization header"))
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		_, err := jwt.VerifyToken(tokenString)
		if err != nil {
			api.NewError(w, http.StatusUnauthorized, errors.New("invalid token"))
			return
		}

		next.ServeHTTP(w, r)
	})
}
