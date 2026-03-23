package usecases

import (
	modelCrypto "crypto-server/internal/crypto/models"
	historydtos "crypto-server/internal/crypto/histories/dtos"
)

func (cryptoUC *cryptoUseCase) GetHistory(symbol string) *historydtos.HistoryResponse {
	ID := cryptoUC.repositories.GetCryptoID(symbol)
	if len(ID) == 0 {
		return nil
	}

	histories := cryptoUC.repositories.GetHistory(modelCrypto.CryptoID(ID))
	items := make([]historydtos.HistoryItemResponse, 0, len(histories))
	for _, item := range histories {
		items = append(items, historydtos.HistoryItemResponse{
			Price:     item.Price,
			Timestamp: item.Timestamp,
		})
	}

	response := &historydtos.HistoryResponse{
		Symbol:  symbol,
		History: items,
	}
	return response
}