package shared

import (
	"net/http"

	config "crypto-server/internal"
	securecookie "github.com/gorilla/securecookie"
)

// TODO добавить РЕФРЕШ КУК!!!!!

func CreateCookies(w http.ResponseWriter, userID int, config config.JWTConfig) (*http.Cookie, error)  {
	var s = securecookie.New([]byte(config.HashKey), []byte(config.BlockKey)) 
	payload, err := s.Encode(config.Name, userID)
	if err != nil {
		return nil, err
	}

	cookie := &http.Cookie{
		Name:     config.Name,
		Value:    payload,
		Path: 	  "/",
		Secure:   false,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	}

	err = cookie.Valid()
	if err != nil {
		return nil, err
	}
	return cookie, nil
}

func ParseCookeis(r *http.Request, config config.JWTConfig) error  {
	auth, err := r.Cookie(config.Name)
	if err != nil {
		return err
	}
	codecs := []securecookie.Codec{
		securecookie.New([]byte(config.HashKey), []byte(config.BlockKey)),
        securecookie.New([]byte(config.HashKey), []byte(config.BlockKey)),
	}
	var userID int
	err = securecookie.DecodeMulti(config.Name, auth.Value, &userID, codecs...)
	if err != nil {
		return err
	}
	return nil
}
