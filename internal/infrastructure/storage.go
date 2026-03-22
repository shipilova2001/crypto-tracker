package infrastructure

import (
	"sync"
	modelUser "crypto-server/internal/users/models"
	modelCrypto "crypto-server/internal/crypto/models"
)

type LocalStorage struct {
	Users 				map[int]*modelUser.User
	Cryptos 			map[string]*modelCrypto.Crypto
	NamingMapSymbolID	map[string][]string
	Mutex 				*sync.RWMutex
}

func NewLocalStorage() *LocalStorage {
	return &LocalStorage{
		Users: 				make(map[int]*modelUser.User, 1024),
		Cryptos: 			make(map[string]*modelCrypto.Crypto, 1024),
		NamingMapSymbolID: 	make(map[string][]string, 1024),
		Mutex: 				new(sync.RWMutex),
	}
}
