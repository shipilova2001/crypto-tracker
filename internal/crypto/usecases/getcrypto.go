package usecases

import (
	modelCrypto "crypto-server/internal/crypto/models"
	dtos "crypto-server/internal/crypto/dtos"
)

func (cryptoUC *cryptoUseCase) GetCrypto(symbol string) *dtos.CryptoResponse {
	ID := cryptoUC.repositories.GetCryptoID(symbol)
	if len(ID) == 0 {
		return nil
	}
	//modelCrypto.CryptoID
	data := cryptoUC.repositories.GetCrypto(modelCrypto.CryptoID(ID))
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
