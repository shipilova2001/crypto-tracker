package usecases

import (
	"net/http"
	dtos "crypto-server/internal/users/dtos"
	repositories "crypto-server/internal/users/repositories"
	passwordServices "crypto-server/internal/users/services/password"
)

type AuthUseCase interface {
	Register(user *dtos.Auth) *dtos.Response
	Login(user *dtos.Auth) *dtos.Response
}

type authUseCase struct {
	UserRepository repositories.UserRepository
}

func New (UserRepository repositories.UserRepository) *authUseCase {
	return &authUseCase{
		UserRepository: UserRepository,
	}
}

func (authUC *authUseCase) Register(user *dtos.Auth) *dtos.Response {
	doesExist := authUC.UserRepository.FindByUsername(user.Username)
	if doesExist != nil {
		return &dtos.Response{
			Code:    http.StatusUnprocessableEntity,
			Message: "User already exist",
			Data: &dtos.UserResponse{
				ID:        doesExist.ID,
				Username:  doesExist.Username,
				CreatedAt: doesExist.CreatedAt,
				UpdatedAt: doesExist.UpdatedAt,
			},
		}
	}

	hash := passwordServices.GetHash(user.Password)
	if hash == nil {
		return &dtos.Response{
			Code:    http.StatusInternalServerError,
			Message: "Password hashing error",
			Data:    nil,
		}
	}
	
	user.Password = string(hash)
	response := authUC.UserRepository.Create(user)
	answer := &dtos.Response{
		Code:    http.StatusCreated,
		Message: "User created success",
		Data: &dtos.UserResponse{
			ID:        response.ID,
			Username:  response.Username,
			CreatedAt: response.CreatedAt,
			UpdatedAt: response.UpdatedAt,
		},
	}
	if response == nil {
		answer.Message = "User created faild"
		answer.Code = http.StatusInternalServerError
	}
	return answer
}


func (authUC *authUseCase) Login(user *dtos.Auth) *dtos.Response {
	doesExist := authUC.UserRepository.FindByUsername(user.Username)
	if doesExist == nil {
		return &dtos.Response{
			Code:    http.StatusNotFound,
			Message: "User not found",
			Data: nil,
		}
	}

	err := passwordServices.Compare(doesExist.Password, user.Password)

	if err != nil {
		return &dtos.Response{
			Code:    http.StatusUnprocessableEntity,
			Message: "Password not compare",
			Data:    nil,
		}
	}

	return &dtos.Response{
		Code:    http.StatusAccepted,
		Message: "User login succesfully",
		Data:    &dtos.UserResponse{
			ID:        doesExist.ID,
			Username:  doesExist.Username,
			CreatedAt: doesExist.CreatedAt,
			UpdatedAt: doesExist.UpdatedAt,
		},
	}
}
