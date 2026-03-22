package usecases

import (
	dtos "crypto-server/internal/crypto/dtos"
	modelCrypto "crypto-server/internal/crypto/models"
)

func (cryptoUC *cryptoUseCase) Refresh(symbol string) *dtos.CryptoResponse {
	ID := cryptoUC.repositories.GetCryptoID(symbol)
	if len(ID) == 0 {
		return nil
	}
	data := cryptoUC.repositories.RefreshCryptoAndHistory(modelCrypto.CryptoID(ID))
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
