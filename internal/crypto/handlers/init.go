package handlers

import (
	"net/http"
	config "crypto-server/internal"
	infrastructure "crypto-server/internal/infrastructure"
)

type handlerCrypto interface {
	Get(w http.ResponseWriter, r *http.Request)
	GetCrypto(w http.ResponseWriter, r *http.Request)
	AddToTrack(w http.ResponseWriter, r *http.Request)
	Refresh(w http.ResponseWriter, r *http.Request)
	GetHistory(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Stats(w http.ResponseWriter, r *http.Request)
}


func InitHandleCrypto(handler handlerCrypto, mux *http.ServeMux, config config.JWTConfig) {	
	mux.Handle("GET /crypto/{symbol}", infrastructure.AuthMiddleware(http.HandlerFunc(handler.GetCrypto), config))
	mux.Handle("GET /crypto", infrastructure.AuthMiddleware(http.HandlerFunc(handler.Get), config))
	mux.Handle("POST /crypto", infrastructure.AuthMiddleware(http.HandlerFunc(handler.AddToTrack), config))
	// принудительно обновить цену криптовалюты
	mux.Handle("GET /crypto/{symbol}/history", infrastructure.AuthMiddleware(http.HandlerFunc(handler.GetHistory), config))
	mux.Handle("GET /crypto/{symbol}/stats", infrastructure.AuthMiddleware(http.HandlerFunc(handler.Stats), config))
	mux.Handle("PUT /crypto/{symbol}/refresh", infrastructure.AuthMiddleware(http.HandlerFunc(handler.Refresh), config))
	mux.Handle("DELETE /crypto/{symbol}", infrastructure.AuthMiddleware(http.HandlerFunc(handler.Delete), config))
}
