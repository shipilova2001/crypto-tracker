package dtos

import (
	models "crypto-server/internal/crypto/models"
	"time"
)

type SymbolCrypto struct {
	Symbol string `json:"symbol"`
}

type CryptoJSON struct {
	ID models.CryptoID		`json:"id"`	
	Symbol string			`json:"symbol"`
	Name string				`json:"name"`
	Platforms interface{}	`json:"platforms"`
}

type CryptoResponse struct {
	Symbol       string      `json:"symbol"`
	Name         string      `json:"name"`
	CurrentPrice float64     `json:"current_price"`
	LastUpdated  time.Time   `json:"last_updated"`
	Platforms    interface{} `json:"platforms,omitempty"`
}

type CryptoListResponse struct {
	Cryptos []*CryptoResponse `json:"cryptos"`
}
type CryptoItemResponse struct {
	Crypto *CryptoResponse    `json:"crypto"`
}


type CoinGeckoCryptoResponse struct {
	ID          models.CryptoID     `json:"id"`
	Symbol      string      		`json:"symbol"`
	Name        string      		`json:"name"`
	Platforms   interface{} 		`json:"platforms"`
	LastUpdated time.Time   		`json:"last_updated"`
	MarketData  MarketData  		`json:"market_data"`
}

type MarketData struct {
	CurrentPrice CurrentPrice `json:"current_price"`
}

type CurrentPrice struct {
	RUB float64 `json:"rub"`
}


type StatsItemResponse struct {
	MinPrice    		float64      `json:"min_price"`
	MaxPrice      		float64      `json:"max_price"`
	AVGPrice        	float64      `json:"avg_price"`
	ChangePrice   		float64 	 `json:"price_change"`
	PriceChangePercent  float64      `json:"price_change_percent"`
	RecordsCount   		int 		 `json:"records_count"`
}

type StatsResponse struct {
	Symbol       string      		`json:"symbol"`
	CurrentPrice float64      		`json:"current_price"`
	Stats        StatsItemResponse  `json:"stats"`
}