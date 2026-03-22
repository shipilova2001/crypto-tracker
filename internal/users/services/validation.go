package services

import (
	"net/http"
	dtos "crypto-server/internal/users/dtos"
)
func Validation (user *dtos.Auth) *dtos.Response {
	var errorResponse *dtos.Response
	lenPassword := len(user.Password)
	lenUsername := len(user.Username)
	
	if lenPassword < 5 || lenPassword > 50 {
		errorResponse = &dtos.Response{
			Code:    http.StatusUnprocessableEntity,
			Message: "Invalid password",
			Data:    nil,
		}
		return errorResponse
	}
	
	if lenUsername < 1 || lenUsername > 50 {
		errorResponse = &dtos.Response{
			Code:    http.StatusUnprocessableEntity,
			Message: "Invalid username",
			Data:    nil,
		}
		return errorResponse
	}
	
	return nil
}