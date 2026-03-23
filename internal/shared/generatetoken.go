package shared

import (
	"fmt"
	"time"
	"net/http"

	"strconv"
	"github.com/golang-jwt/jwt/v5"
	config "crypto-server/internal"
)

type CustomClaims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

func generateToken(userID string, config config.JWTConfig) string {
	expirationTime := time.Now().Add(24 * time.Hour)
	var jwtSecret = []byte(config.BlockKey)

	claims := &CustomClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Subject:   "user_authentication",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	
	if err != nil {
		fmt.Printf("Ошибка подписи токена: %v\n", err)
		return ""
	}
	return tokenString

}


func SetToken(w http.ResponseWriter, userID int, config config.JWTConfig) string {
	token := generateToken(strconv.Itoa(userID), config)
	w.Header().Add(config.Name, fmt.Sprintf("Bearer %s", token))
	return token
}
