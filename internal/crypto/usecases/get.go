package usecases

import (
	modelCrypto "crypto-server/internal/crypto/models"
	dtos "crypto-server/internal/crypto/dtos"
)

func toCryptoItemResponse(item *modelCrypto.Crypto) *dtos.CryptoResponse {
	if item == nil {
		return nil
	}

	return &dtos.CryptoResponse{
		Symbol:       item.Symbol,
		Name:         item.Name,
		CurrentPrice: item.CurrentPrice,
		LastUpdated:  item.LastUpdated,
	}
}

func (cryptoUC *cryptoUseCase) Get() *dtos.CryptoListResponse {
	data := cryptoUC.repositories.Get()

	response := &dtos.CryptoListResponse{
		Cryptos: make([]*dtos.CryptoResponse, 0, len(data)),
	}

	for _, item := range data {
		response.Cryptos = append(response.Cryptos, toCryptoItemResponse(item))
	}

	return response
}
