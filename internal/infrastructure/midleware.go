package infrastructure

import (
	config "crypto-server/internal"
	shared "crypto-server/internal/shared"
	"fmt"
	"net/http"
)

func AuthMiddleware(next http.Handler, config config.JWTConfig) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get(config.Name)
		tokenString, err := shared.ExtractBearerToken(auth)
		fmt.Println(tokenString, err, config.Name)
		if err != nil {
			http.Error(w, "Невалидный токен", http.StatusUnauthorized)
			return
		}
		_, err = shared.ParseToken(tokenString, config.BlockKey)
		fmt.Println(err)
		if err != nil {
			http.Error(w, "Невалидный токен", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
