package usecases

import (
	dtos "crypto-server/internal/crypto/dtos"
	modelCrypto "crypto-server/internal/crypto/models"
)

func (cryptoUC *cryptoUseCase) Refresh(symbol string) *dtos.CryptoItemResponse {
	ID := cryptoUC.repositories.GetCryptoID(symbol)
	if len(ID) == 0 {
		return nil
	}
	data := cryptoUC.repositories.RefreshCryptoAndHistory(modelCrypto.CryptoID(ID))
	if data == nil {
		return nil
	}

	return &dtos.CryptoItemResponse{
		Crypto: &dtos.CryptoResponse{
			Symbol:       data.Symbol,
			Name:         data.Name,
			CurrentPrice: data.CurrentPrice,
			LastUpdated:  data.LastUpdated,
		},
	}
}

func (cryptoUC *cryptoUseCase) RefreshAll() {
	cryptos := cryptoUC.repositories.Get()
	// response := &dtos.CryptoListResponse{
	// 	Cryptos: make([]*dtos.CryptoResponse, 0),
	// }
	for _, item := range cryptos {
		cryptoUC.Refresh(item.Symbol)
		// response.Cryptos = append(response.Cryptos, crypto.Crypto)
	}
	// return response
}
