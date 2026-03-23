package usecases

import (
	modelCrypto "crypto-server/internal/crypto/models"
	dtos "crypto-server/internal/crypto/dtos"
)
func (cryptoUC *cryptoUseCase) Stats(symbol string) *dtos.StatsResponse {
	ID := cryptoUC.repositories.GetCryptoID(symbol)
	if len(ID) == 0 {
		return nil
	}

	histories := cryptoUC.repositories.GetHistory(modelCrypto.CryptoID(ID))
	count := len(histories)
	if count == 0 {
		return nil
	}

	first := histories[0]
	last := histories[count-1]

	min := first.Price
	max := first.Price
	sum := 0.0

	for _, item := range histories {
		price := item.Price
		sum += price

		if price < min {
			min = price
		}
		if price > max {
			max = price
		}
	}

	avg := sum / float64(count)
	change := last.Price - first.Price

	changePercent := 0.0
	if first.Price != 0 {
		changePercent = (change / first.Price) * 100
	}

	return &dtos.StatsResponse{
		Symbol:       symbol,
		CurrentPrice: last.Price,
		Stats: dtos.StatsItemResponse{
			MinPrice:           min,
			MaxPrice:           max,
			AVGPrice:           avg,
			ChangePrice:        change,
			PriceChangePercent: changePercent,
			RecordsCount:       count,
		},
	}
}