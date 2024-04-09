package api

import (
	"context"
	"github.com/jenkinsyoung/web_music_sanctuary/internal/jwt"
	"log"
	"net/http"
	"strings"
)

func Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		headerParts := strings.Split(r.Header.Get("Authorization"), " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			log.Printf("invalid auth header")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if len(headerParts[1]) == 0 {
			log.Printf("token is empty")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		claims, err := jwt.ParseToken(headerParts[1])
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
		}
		ctx := context.WithValue(context.Background(), "userId", claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
