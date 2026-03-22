package infrastructure

import (
	config "crypto-server/internal"
	shared "crypto-server/internal/shared"
	"fmt"
	"net/http"
)


func AuthMiddleware(next http.Handler, config config.JWTConfig) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := shared.ParseCookeis(r, config)
		fmt.Println(err)
		if err != nil {
			http.Error(w, "Невалидные куки", 401)
			return
		}
		next.ServeHTTP(w, r)
	})
}