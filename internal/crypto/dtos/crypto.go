package dtos

import "time"

type SymbolCrypto struct {
	Symbol string `json:"symbol"`
}

type CryptoJSON struct {
	ID string			    `json:"id"`	
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
	Crypto []*CryptoResponse `json:"crypto"`
}


type CoinGeckoCryptoResponse struct {
	ID          string      `json:"id"`
	Symbol      string      `json:"symbol"`
	Name        string      `json:"name"`
	Platforms   interface{} `json:"platforms"`
	LastUpdated time.Time      `json:"last_updated"`
	MarketData  MarketData  `json:"market_data"`
}

type MarketData struct {
	CurrentPrice CurrentPrice `json:"current_price"`
}

type CurrentPrice struct {
	RUB float64 `json:"rub"`
}