package adapters

import (
	config 			"crypto-server/internal"
	dtos			"crypto-server/internal/crypto/dtos"
)

type CryptoAdapter interface {
	GetList() []*dtos.CoinGeckoCryptoResponse
	GetCrypto(symbol string) *dtos.CoinGeckoCryptoResponse
}

type cryptoAdapter struct {
	configAPI config.APIConfig
}


func New (configAPI config.APIConfig) *cryptoAdapter {
	return &cryptoAdapter{
		configAPI: configAPI,
	}
}