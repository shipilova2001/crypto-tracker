package shared
import (
	"github.com/golang-jwt/jwt/v5"
	"fmt"
)

func ParseToken(tokenString, jwtSecret string) (*CustomClaims, error) {
    claims := &CustomClaims{}

    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("ожидаемый метод для подписи: %v", token.Method)
        }
        return []byte(jwtSecret), nil
    })

    if err != nil {
        return nil, err
    }

    if !token.Valid {
        return nil, fmt.Errorf("невалидный токен")
    }
    return claims, nil
}