package repositories

import (
	"time"
	infrastructure "crypto-server/internal/infrastructure"
	dtos "crypto-server/internal/users/dtos"
	modelUser "crypto-server/internal/users/models"
)

type userLocalStorageRepository struct { 
	storage *infrastructure.LocalStorage 
}

func New (storage *infrastructure.LocalStorage) *userLocalStorageRepository {
	return &userLocalStorageRepository{
		storage: storage,
	}
}


func (userRepository *userLocalStorageRepository) IsExistUser(userID int) bool {
	localstorage := userRepository.storage
	localstorage.Mutex.RLock()
	defer localstorage.Mutex.RUnlock()
	if _, ok := localstorage.Users[userID]; !ok {
		return false
	}
	return true
}

func (userRepository *userLocalStorageRepository) FindByUsername (userName string) *modelUser.User {
	localstorage := userRepository.storage
	localstorage.Mutex.RLock()
	defer localstorage.Mutex.RUnlock()
	for _, item := range localstorage.Users {
		if item.Username == userName {
			return &modelUser.User{
				ID: 		item.ID,
				Username: 	item.Username,
				Password: 	item.Password,
				CreatedAt: 	item.CreatedAt,
				UpdatedAt: 	item.UpdatedAt,
			}
		}
	}
	return nil
}


func (userRepository *userLocalStorageRepository) Create (data *dtos.Auth) *modelUser.User {
	localstorage := userRepository.storage
	lenUsers := len(localstorage.Users) + 1
	localstorage.Mutex.Lock()
	defer localstorage.Mutex.Unlock()
	
	localstorage.Users[lenUsers] = &modelUser.User{
		ID:		   lenUsers,
		Username:  data.Username,
		Password:  data.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return localstorage.Users[lenUsers]
}
