package usecases

import (
	dtos "crypto-server/internal/crypto/dtos"
	"fmt"
	"time"
)

func (cryptoUC *cryptoUseCase) AddToTrackCrypto(symbol string) *dtos.CryptoItemResponse {
	ID := cryptoUC.repositories.GetCryptoID(symbol)
	fmt.Println("ID ", ID)
	if len(ID) == 0 {
		return nil
	}
	crypto := cryptoUC.adapters.GetCrypto(ID)
	fmt.Println("crypto ", crypto)
	if crypto == nil {
		return nil
	}
	crypto.LastUpdated = time.Now()
	data := cryptoUC.repositories.Add(crypto)
	response := &dtos.CryptoItemResponse{
		Crypto: make([]*dtos.CryptoResponse, 0, 1),
	}
	response.Crypto = append(response.Crypto, &dtos.CryptoResponse{
		Symbol:       data.Symbol,
		Name:         data.Name,
		CurrentPrice: data.CurrentPrice,
		LastUpdated:  data.LastUpdated,
	})
	fmt.Println("response ", response)
	return response
}
