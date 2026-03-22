package repositories

import (
	"crypto-server/internal/crypto/dtos"
	"strings"
	"time"

	maps "golang.org/x/exp/maps"
	modelCrypto "crypto-server/internal/crypto/models"
)




func (localStorage *cryptoLocalStorageRepository) GetCryptoID(symbol string) string {
	if items, ok := localStorage.storage.NamingMapSymbolID[symbol]; ok {
		return items[0]
	}
	symbol = strings.ToLower(symbol)
	if items, ok := localStorage.storage.NamingMapSymbolID[symbol]; ok {
		return items[0]
	}
	return ""
}

func (localStorage *cryptoLocalStorageRepository) SetNamingMapSymbolID(data []*dtos.CoinGeckoCryptoResponse) {
 	for _, s := range data {
        localStorage.storage.NamingMapSymbolID[s.Symbol] = append(localStorage.storage.NamingMapSymbolID[s.Symbol], s.ID)
    }
}



func (localStorage *cryptoLocalStorageRepository) Get() []*modelCrypto.Crypto {
	return maps.Values(localStorage.storage.Cryptos)
}

func (localStorage *cryptoLocalStorageRepository) GetCrypto(ID string) *modelCrypto.Crypto {
	if item, ok := localStorage.storage.Cryptos[ID]; ok {
		return item
	}
	return nil
}

func (localStorage *cryptoLocalStorageRepository) Add(crypto *dtos.CoinGeckoCryptoResponse) *modelCrypto.Crypto {
	storage := localStorage.storage
	storage.Mutex.Lock()
	defer storage.Mutex.Unlock()
	
	res := &modelCrypto.Crypto{
		ID:           crypto.ID,
		Symbol:       crypto.Symbol,
		Name:         crypto.Name,
		CurrentPrice: crypto.MarketData.CurrentPrice.RUB,
		LastUpdated:  time.Now(),
		Platforms:    crypto.Platforms,
	}
	storage.Cryptos[crypto.ID] = res
	return storage.Cryptos[crypto.ID]
}
