package infrastructure

import (
	"sync"
	modelUser "crypto-server/internal/users/models"
	modelCrypto "crypto-server/internal/crypto/models"
	modelHistory "crypto-server/internal/crypto/histories/models"
)


type LocalStorage struct {
	Users 				map[int]*modelUser.User
	Cryptos 			map[modelCrypto.CryptoID]*modelCrypto.Crypto
	History 			map[modelCrypto.CryptoID][]modelHistory.History
	NamingMapSymbolID	map[string][]modelCrypto.CryptoID
	Mutex 				*sync.RWMutex
}

func NewLocalStorage() *LocalStorage {
	return &LocalStorage{
		Users: 				make(map[int]*modelUser.User, 1024),
		Cryptos: 			make(map[modelCrypto.CryptoID]*modelCrypto.Crypto, 1024),
		History: 			make(map[modelCrypto.CryptoID][]modelHistory.History, 1024),
		NamingMapSymbolID: 	make(map[string][]modelCrypto.CryptoID, 1024),
		Mutex: 				new(sync.RWMutex),
	}
}
