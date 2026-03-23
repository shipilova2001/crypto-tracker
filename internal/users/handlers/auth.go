package handlers

import (
	config "crypto-server/internal"
	shared "crypto-server/internal/shared"
	dtos "crypto-server/internal/users/dtos"
	"crypto-server/internal/users/services"
	usecases "crypto-server/internal/users/usecases"
	"encoding/json"
	"net/http"
)

type UserHandler interface {
	Register(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
}

type AuthHandler struct {
	authUC       usecases.AuthUseCase
	CookieConfig *config.JWTConfig
}

func New(authUC usecases.AuthUseCase, cookieConfig *config.JWTConfig) *AuthHandler {
	return &AuthHandler{
		authUC:       authUC,
		CookieConfig: cookieConfig,
	}
}

func (authHandler *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var auth dtos.AuthJSON
	err := json.NewDecoder(r.Body).Decode(&auth)
	if err != nil {
		resp := &dtos.ResponseError{
			Error: err.Error(),
		}
		shared.WriteJSON(w, resp, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	user := &dtos.Auth{
		Username: auth.Username,
		Password: auth.Password,
	}

	if err := services.Validation(user); err != nil {
		resp := &dtos.ResponseError{
			Error: err.Message,
		}
		shared.WriteJSON(w, resp, http.StatusUnprocessableEntity)
		return
	}

	response := authHandler.authUC.Register(user)
	if response.Code < 200 || response.Code > 299 {
		resp := &dtos.ResponseError{
			Error: response.Message,
		}
		shared.WriteJSON(w, resp, response.Code)
		return
	}

	// fmt.Println(response.Data)
	// cookie, err := shared.CreateCookies(w, response.Data.ID, *authHandler.CookieConfig)
	// if err != nil {
	// 	shared.WriteJSON(w, &dtos.Response{
	// 		Code: http.StatusInternalServerError,
	// 		Message: err.Error(),
	// 		Data: nil,
	// 	}, http.StatusInternalServerError)
	// 	return;
	// }
	// http.SetCookie(w, cookie)
	token := shared.SetToken(w, response.Data.ID, *authHandler.CookieConfig)
	resp := &dtos.ResponseToken{
		Token: token,
	}
	shared.WriteJSON(w, resp, response.Code)
}

func (authHandler *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var auth dtos.AuthJSON
	err := json.NewDecoder(r.Body).Decode(&auth)
	if err != nil {
		resp := &dtos.ResponseError{
			Error: err.Error(),
		}
		shared.WriteJSON(w, resp, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	user := &dtos.Auth{
		Username: auth.Username,
		Password: auth.Password,
	}
	if err := services.Validation(user); err != nil {
		resp := &dtos.ResponseError{
			Error: err.Message,
		}
		shared.WriteJSON(w, resp, http.StatusUnprocessableEntity)
		return
	}

	response := authHandler.authUC.Login(user)

	if response.Code < 200 || response.Code > 299 {
		resp := &dtos.ResponseError{
			Error: response.Message,
		}
		shared.WriteJSON(w, resp, response.Code)
		return
	}

	// cookie, err := shared.CreateCookies(w, response.Data.ID, *authHandler.CookieConfig)
	// if err != nil {
	// 	shared.WriteJSON(w, &dtos.Response{
	// 		Code: http.StatusInternalServerError,
	// 		Message: err.Error(),
	// 		Data: nil,
	// 	}, response.Code)
	// 	return;
	// }
	// http.SetCookie(w, cookie)
	token := shared.SetToken(w, response.Data.ID, *authHandler.CookieConfig)
	resp := &dtos.ResponseToken{
		Token: token,
	}
	shared.WriteJSON(w, resp, response.Code)
}
