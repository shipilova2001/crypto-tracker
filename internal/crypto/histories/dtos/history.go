package models

import "time"

type HistoryItemResponse struct {
	Price			float64		 	`json:"price"`	
	Timestamp		time.Time		`json:"timestamp"`	
}

type HistoryResponse struct {
	Symbol			string					`json:"symbol"`
	History 		[]HistoryItemResponse 	`json:"history"`
}
