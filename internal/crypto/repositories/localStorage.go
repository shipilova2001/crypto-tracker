package repositories

import (
	"crypto-server/internal/crypto/dtos"
	"crypto-server/internal/crypto/histories/models"
	"strings"
	"time"

	modelCrypto "crypto-server/internal/crypto/models"
	modelHistory "crypto-server/internal/crypto/histories/models"
	maps "golang.org/x/exp/maps"
)




func (localStorage *cryptoLocalStorageRepository) GetCryptoID(symbol string) modelCrypto.CryptoID {
	localStorage.storage.Mutex.RLock()
	defer localStorage.storage.Mutex.RUnlock()
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
	localStorage.storage.Mutex.Lock()
	defer localStorage.storage.Mutex.Unlock()
	for _, s := range data {
        localStorage.storage.NamingMapSymbolID[s.Symbol] = append(localStorage.storage.NamingMapSymbolID[s.Symbol], s.ID)
    }
}



func (localStorage *cryptoLocalStorageRepository) Get() []*modelCrypto.Crypto {
	localStorage.storage.Mutex.RLock()
	defer localStorage.storage.Mutex.RUnlock()
	return maps.Values(localStorage.storage.Cryptos)
}

func (localStorage *cryptoLocalStorageRepository) GetCrypto(ID modelCrypto.CryptoID) *modelCrypto.Crypto {
	localStorage.storage.Mutex.RLock()
	defer localStorage.storage.Mutex.RUnlock()
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

func (localStorage *cryptoLocalStorageRepository) RefreshCrypto(ID modelCrypto.CryptoID) *modelCrypto.Crypto {
	localStorage.storage.Mutex.Lock()
	defer localStorage.storage.Mutex.Unlock()
	
	if item, ok := localStorage.storage.Cryptos[ID]; ok {
		newData := localStorage.adapters.GetCrypto(ID)
		item.CurrentPrice = newData.MarketData.CurrentPrice.RUB
		item.LastUpdated = time.Now()
		return item
	}
	return nil
}


func (localStorage *cryptoLocalStorageRepository) RefreshCryptoAndHistory(ID modelCrypto.CryptoID) *modelCrypto.Crypto {
	newData := localStorage.RefreshCrypto(ID)
	if newData == nil {
		return nil
	}

	history := models.History{
		Price:     newData.CurrentPrice,
		Timestamp: time.Now(),
	}

	localStorage.storage.Mutex.Lock()
	defer localStorage.storage.Mutex.Unlock()

	items := localStorage.storage.History[ID]
	items = append(items, history)

	// Каждая криптовалюта хранит историю из максимум 100 последних цен
	if len(items) > 100 {
		items = items[len(items)-100:]
	}
	
	localStorage.storage.History[ID] = items
	return newData
}


func (localStorage *cryptoLocalStorageRepository) GetHistory(ID modelCrypto.CryptoID) []modelHistory.History {
	localStorage.storage.Mutex.RLock()
	defer localStorage.storage.Mutex.RUnlock()
	if items, ok := localStorage.storage.History[ID]; ok {
		return items
	}
	
	return []modelHistory.History{}
}

func (localStorage *cryptoLocalStorageRepository) DeleteHistory(ID modelCrypto.CryptoID) any {
	localStorage.storage.Mutex.Lock()
	defer localStorage.storage.Mutex.Unlock()
	delete(localStorage.storage.History, ID)
	return struct{}{}
}

func (localStorage *cryptoLocalStorageRepository) DeleteCrypto(ID modelCrypto.CryptoID) any {
	localStorage.storage.Mutex.Lock()
	defer localStorage.storage.Mutex.Unlock()
	delete(localStorage.storage.Cryptos, ID)
	return struct{}{}
}

