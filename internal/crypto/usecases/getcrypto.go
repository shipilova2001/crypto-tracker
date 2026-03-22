package usecases

import (
	dtos "crypto-server/internal/crypto/dtos"
)

func (cryptoUC *cryptoUseCase) GetCrypto(symbol string) *dtos.CryptoResponse {
	ID := cryptoUC.repositories.GetCryptoID(symbol)
	if len(ID) == 0 {
		return nil
	}
	data := cryptoUC.repositories.GetCrypto(ID)
	if data == nil {
		return nil
	}
	return &dtos.CryptoResponse{
		Symbol:       data.Symbol,
		Name:         data.Name,
		CurrentPrice: data.CurrentPrice,
		LastUpdated:  data.LastUpdated,
	}
}
