package middlewares

import (
	"RSOI/src/config"
	"context"
	"github.com/robbert229/jwt"
	"net/http"
	"strings"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		authHeader := r.Header.Get("Authorization")
		if authHeader != "" {
			authToken := strings.Split(authHeader, " ")[1]
			algorithm := jwt.HmacSha256(config.JWTKey)

			claims, err := algorithm.DecodeAndValidate(authToken)
			if err == nil {
				uuid, _ := claims.Get("UUID")
				if err == nil {
					ctx = context.WithValue(ctx, "UUID", uuid)
				}
				username, _ := claims.Get("Username")
				if err == nil {
					ctx = context.WithValue(ctx, "Username", username)
				}
			}
		}
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
