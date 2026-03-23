package adapters

import (
	"io"
	"net/http"
	"encoding/json"
	dtos 			"crypto-server/internal/crypto/dtos"
)

func (cryptoAdapter *cryptoAdapter) GetList() []*dtos.CoinGeckoCryptoResponse {
	//TODO check '/' for url
	configAPI := cryptoAdapter.configAPI
	url := configAPI.BaseURL + "/coins/list"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil
	}
	req.Header.Set(configAPI.KeyName, configAPI.KeyAPI)
	
	client := http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return nil
	}
	defer response.Body.Close()
	
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return nil
	}
	
	var list []*dtos.CoinGeckoCryptoResponse
	err = json.Unmarshal(data, &list)
	if err != nil {
		return nil
	}
	
	return list
}
