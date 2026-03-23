package internal

import "os"

type Config struct {
	JWT    			JWTConfig
	API				APIConfig
}


type APIConfig struct {
	BaseURL 	string
	KeyName 	string
	KeyAPI 		string
}

type JWTConfig struct {
	HashKey		string
	BlockKey 	string
	Name		string
}


func Load() *Config {
	return &Config{
		JWT: JWTConfig{
			HashKey: 	os.Getenv("COOKIE_HASH"),
			BlockKey: 	os.Getenv("COOKIE_BLOCK_KEY"),
			Name: 		os.Getenv("NAME_COOKIE"),
		},
		API: APIConfig{
			BaseURL: 	os.Getenv("API_BASE_URL"),
			KeyName: 	os.Getenv("API_KEY_NAME"),
			KeyAPI: 	os.Getenv("API_KEY"),
		},
	}
}

