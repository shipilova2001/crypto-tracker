package models

import "time"

type CryptoID string

type Crypto struct {
	ID 				CryptoID
	Symbol 			string
	Name 			string
	CurrentPrice 	float64
	LastUpdated		time.Time	
	Platforms 		interface{}
	// UserID 			string
}

