package repositories

import (
	adapters "crypto-server/internal/crypto/adapters"
	dtos "crypto-server/internal/crypto/dtos"
	modelCrypto "crypto-server/internal/crypto/models"
	modelHistory "crypto-server/internal/crypto/histories/models"
	infrastructure "crypto-server/internal/infrastructure"
)

type CryptoRepository interface {
	GetCryptoID(symbol string) modelCrypto.CryptoID
	SetNamingMapSymbolID(data []*dtos.CoinGeckoCryptoResponse)
	
	Get() []*modelCrypto.Crypto
	GetCrypto(ID modelCrypto.CryptoID) *modelCrypto.Crypto
	Add(crypto *dtos.CoinGeckoCryptoResponse) *modelCrypto.Crypto
	RefreshCrypto(ID modelCrypto.CryptoID) *modelCrypto.Crypto
	RefreshCryptoAndHistory(ID modelCrypto.CryptoID) *modelCrypto.Crypto
	GetHistory(ID modelCrypto.CryptoID) []modelHistory.History
	DeleteHistory(ID modelCrypto.CryptoID) any
	DeleteCrypto(ID modelCrypto.CryptoID) any
}


type cryptoLocalStorageRepository struct {
	storage *infrastructure.LocalStorage
	adapters adapters.CryptoAdapter
}

func New (storage *infrastructure.LocalStorage, adapters adapters.CryptoAdapter) *cryptoLocalStorageRepository {
	return &cryptoLocalStorageRepository{
		storage: storage,
		adapters: adapters,
	}
}