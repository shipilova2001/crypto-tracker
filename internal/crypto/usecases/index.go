package usecases

import (
	dtos 			"crypto-server/internal/crypto/dtos"
	adapters 		"crypto-server/internal/crypto/adapters"
	repositories 	"crypto-server/internal/crypto/repositories"
)

type CryptoUseCase interface {
	InitCoinMapping()
	
 	AddToTrackCrypto(symbol string) *dtos.CryptoItemResponse
	Get() *dtos.CryptoListResponse
	GetCrypto(symbol string) *dtos.CryptoResponse
}

type cryptoUseCase struct {
	adapters adapters.CryptoAdapter
	repositories repositories.CryptoRepository
}

func New (adapters adapters.CryptoAdapter, repositories repositories.CryptoRepository) *cryptoUseCase {
	return &cryptoUseCase{
		adapters: adapters,
		repositories: repositories,
	}
}
