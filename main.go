package main

import (
	"fmt"
	"net/http"
	"github.com/joho/godotenv"
	config "crypto-server/internal"
	infrastructure "crypto-server/internal/infrastructure"
	
	usershandlers "crypto-server/internal/users/handlers"
	usersusecases "crypto-server/internal/users/usecases"
	userrepositories "crypto-server/internal/users/repositories"
	
	cryptohandlers "crypto-server/internal/crypto/handlers"
	cryptousecases "crypto-server/internal/crypto/usecases"
	cryptorepositories "crypto-server/internal/crypto/repositories"
	cryptoadapters "crypto-server/internal/crypto/adapters"
)

func init() {
    if err := godotenv.Load(); err != nil {
        fmt.Print("No .env file found")
    }
}
func main() {
	configAll := config.Load()
	mux := http.NewServeMux()
	localstorage := infrastructure.NewLocalStorage()

	// пользователь
	userRepo := userrepositories.New(localstorage)
	userUsecase := usersusecases.New(userRepo)
	authHandler := usershandlers.New(userUsecase, &configAll.JWT)
	
	//крипта
	cryptoRepo := cryptorepositories.New(localstorage)
	cryptoAdapter := cryptoadapters.New(configAll.API)
	cryptoUsecase := cryptousecases.New(cryptoAdapter, cryptoRepo)
	cryptoHandler := cryptohandlers.New(cryptoUsecase)
	cryptoUsecase.InitCoinMapping()
	
	
	cryptohandlers.InitHandleCrypto(cryptoHandler, mux, configAll.JWT)
	usershandlers.InitHandleUser(authHandler, mux)
	
	// for i, item := range localstorage.NamingMapSymbolID {
	// 	fmt.Println(i, "        ", item)
	// }
	
	initServer(mux)
}
