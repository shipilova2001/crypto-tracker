package repositories

import (
	dtos "crypto-server/internal/users/dtos"
	modelUser "crypto-server/internal/users/models"
)

type UserRepository interface {
	IsExistUser(userID int) 		bool
	FindByUsername(userName string) *modelUser.User
	Create(auth *dtos.Auth)  		*modelUser.User
}
