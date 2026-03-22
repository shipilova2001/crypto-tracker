package adapters

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	dtos "crypto-server/internal/crypto/dtos"
)

func (cryptoAdapter *cryptoAdapter) GetCrypto(id string) *dtos.CoinGeckoCryptoResponse {
	configAPI := cryptoAdapter.configAPI
	url := configAPI.BaseURL + "/coins/" + id
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil
	}
	fmt.Println("req ", req)
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
	var crypto *dtos.CoinGeckoCryptoResponse
	err = json.Unmarshal(data, &crypto)
	if err != nil {
		return nil
	}
	
	return crypto
}
