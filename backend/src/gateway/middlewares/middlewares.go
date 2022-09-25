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
		var id = 0

		authHeader := r.Header.Get("Authorization")
		if authHeader != "" {
			authToken := strings.Split(authHeader, " ")[1]
			algorithm := jwt.HmacSha256(config.JWTKey)

			claims, err := algorithm.DecodeAndValidate(authToken)
			if err == nil {
				identificator, err := claims.Get("Identificator")
				if err == nil {
					id = int(identificator.(float64))
				}
			}
		}

		ctx = context.WithValue(ctx, "id", id)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
