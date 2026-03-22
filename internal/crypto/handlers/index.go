package handlers

import (
	"encoding/json"
	// "fmt"
	dtos "crypto-server/internal/crypto/dtos"
	usecases "crypto-server/internal/crypto/usecases"
	shared "crypto-server/internal/shared"
	"net/http"
)

type CryptoHandler interface {
	GetCrypto(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	AddToTrack(w http.ResponseWriter, r *http.Request)
	// Refresh (w http.ResponseWriter, r *http.Request)
}

type cryptoHandler struct {
	cryptoUseCase usecases.CryptoUseCase
}

func New(cryptoUseCase usecases.CryptoUseCase) *cryptoHandler {
	return &cryptoHandler{
		cryptoUseCase: cryptoUseCase,
	}
}

func (handler *cryptoHandler) GetCrypto(w http.ResponseWriter, r *http.Request) {
	symbol := r.PathValue("symbol")
	response := handler.cryptoUseCase.GetCrypto(symbol)
	if response == nil {
		shared.WriteJSON(w, response, http.StatusNotFound)
		return
	}
	shared.WriteJSON(w, response, http.StatusOK)
}

func (handler *cryptoHandler) Get(w http.ResponseWriter, r *http.Request) {
	response := handler.cryptoUseCase.Get()
	if response == nil {
		shared.WriteJSON(w, response, http.StatusNotFound)
		return
	}
	shared.WriteJSON(w, response, http.StatusOK)
}

func (handler *cryptoHandler) AddToTrack(w http.ResponseWriter, r *http.Request) {
	var body dtos.SymbolCrypto
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		shared.WriteJSON(w, err.Error(), http.StatusBadRequest)
		return
	}
	response := handler.cryptoUseCase.AddToTrackCrypto(body.Symbol)
	if response == nil {
		shared.WriteJSON(w, response, http.StatusNotFound)
		return
	}
	// fmt.Println("response  ", response)
	shared.WriteJSON(w, response, http.StatusOK)
}
