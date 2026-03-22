package models

import "time"

type Crypto struct {
	ID 				string
	Symbol 			string
	Name 			string
	CurrentPrice 	float64
	LastUpdated		time.Time	
	Platforms 		interface{}
	// UserID 			string
}

