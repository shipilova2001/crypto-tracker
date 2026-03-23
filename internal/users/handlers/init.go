package handlers

import (
	"net/http"
)

func InitHandleUser(handler UserHandler, mux *http.ServeMux) {
	mux.HandleFunc("POST /auth/register", handler.Register)
	mux.HandleFunc("POST /auth/login", handler.Login)
}
