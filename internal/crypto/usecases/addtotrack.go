package usecases

import (
	dtos "crypto-server/internal/crypto/dtos"
	"crypto-server/internal/crypto/models"
	"time"
)

func (cryptoUC *cryptoUseCase) AddToTrackCrypto(symbol string) *dtos.CryptoItemResponse {
	ID := cryptoUC.repositories.GetCryptoID(symbol)
	if len(ID) == 0 {
		return nil
	}
	crypto := cryptoUC.adapters.GetCrypto(models.CryptoID(ID))
	if crypto == nil {
		return nil
	}
	crypto.LastUpdated = time.Now()
	data := cryptoUC.repositories.Add(crypto)
	response := &dtos.CryptoItemResponse{
		Crypto: &dtos.CryptoResponse{
			Symbol:       data.Symbol,
			Name:         data.Name,
			CurrentPrice: data.CurrentPrice,
			LastUpdated:  data.LastUpdated,
		},
	}
	return response
}
