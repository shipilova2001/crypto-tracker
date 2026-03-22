package services

import (
	bscrypt "golang.org/x/crypto/bcrypt"
)

func Compare(has string, password string) error {
	err := bscrypt.CompareHashAndPassword([]byte(has), []byte(password))
	if err != nil {
		return err
	}
	return nil
}