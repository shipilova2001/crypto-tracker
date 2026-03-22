package repositories

import (
	modelCrypto	"crypto-server/internal/crypto/models"
	dtos		"crypto-server/internal/crypto/dtos"
	infrastructure 	"crypto-server/internal/infrastructure"
)

type CryptoRepository interface {
	GetCryptoID(symbol string) string
	SetNamingMapSymbolID(data []*dtos.CoinGeckoCryptoResponse)
	
	Get() []*modelCrypto.Crypto
	GetCrypto(ID string) *modelCrypto.Crypto
	Add(crypto *dtos.CoinGeckoCryptoResponse) *modelCrypto.Crypto
}


type cryptoLocalStorageRepository struct {
	storage *infrastructure.LocalStorage
}

func New (storage *infrastructure.LocalStorage) *cryptoLocalStorageRepository {
	return &cryptoLocalStorageRepository{
		storage: storage,
	}
}