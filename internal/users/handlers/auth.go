package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"crypto-server/internal/users/services"
	dtos "crypto-server/internal/users/dtos"
	config "crypto-server/internal"
	usecases "crypto-server/internal/users/usecases"
	shared "crypto-server/internal/shared"
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
		response := &dtos.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		}
		shared.WriteJSON(w, response, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	user := &dtos.Auth{
		Username: auth.Username,
		Password: auth.Password,
	}

	if err := services.Validation(user); err != nil {
		shared.WriteJSON(w, err, http.StatusUnprocessableEntity)
		return
	}

	response := authHandler.authUC.Register(user)
	if response.Code < 200 || response.Code > 299 {
		shared.WriteJSON(w, response, response.Code)
		return
	}

	fmt.Println(response.Data)
	cookie, err := shared.CreateCookies(w, response.Data.ID, *authHandler.CookieConfig)
	if err != nil {
		shared.WriteJSON(w, &dtos.Response{
			Code: http.StatusInternalServerError,
			Message: err.Error(),
			Data: nil,
		}, http.StatusInternalServerError)
		return;
	}

	http.SetCookie(w, cookie)
	shared.WriteJSON(w, response, response.Code)
}



func (authHandler *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var auth dtos.AuthJSON
	err := json.NewDecoder(r.Body).Decode(&auth)
	if err != nil {
		response := &dtos.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		}
		shared.WriteJSON(w, response, response.Code)
		return
	}
	defer r.Body.Close()
	
	user := &dtos.Auth{
		Username: auth.Username,
		Password: auth.Password,
	}
	if err := services.Validation(user); err != nil {
		shared.WriteJSON(w, err, http.StatusUnprocessableEntity)
		return
	}
	
	response := authHandler.authUC.Login(user)
	//TODO убрать в функцию
	if response.Code < 200 || response.Code > 299 {
		shared.WriteJSON(w, response, response.Code)
		return
	}


	// fmt.Println(response.Data)
	cookie, err := shared.CreateCookies(w, response.Data.ID, *authHandler.CookieConfig)
	if err != nil {
		shared.WriteJSON(w, &dtos.Response{
			Code: http.StatusInternalServerError,
			Message: err.Error(),
			Data: nil,
		}, response.Code)
		return;
	}
	http.SetCookie(w, cookie)
	shared.WriteJSON(w, response, response.Code)
}