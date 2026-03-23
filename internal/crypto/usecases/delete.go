package usecases

import (
	modelCrypto "crypto-server/internal/crypto/models"
)

func (cryptoUC *cryptoUseCase) Delete(symbol string) any {
	ID := cryptoUC.repositories.GetCryptoID(symbol)
	if len(ID) == 0 {
		return nil
	}
	cryptoUC.repositories.DeleteHistory(modelCrypto.CryptoID(ID))
	return cryptoUC.repositories.DeleteCrypto(modelCrypto.CryptoID(ID))
}