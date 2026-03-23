package main

import (
	"context"
	config "crypto-server/internal"
	scheduler "crypto-server/internal/infrastructure/scheduler"
	infrastructure "crypto-server/internal/infrastructure"
	"fmt"
	"net/http"

	"github.com/joho/godotenv"

	usershandlers "crypto-server/internal/users/handlers"
	userrepositories "crypto-server/internal/users/repositories"
	usersusecases "crypto-server/internal/users/usecases"

	cryptoadapters "crypto-server/internal/crypto/adapters"
	cryptohandlers "crypto-server/internal/crypto/handlers"
	cryptorepositories "crypto-server/internal/crypto/repositories"
	cryptousecases "crypto-server/internal/crypto/usecases"
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
	cryptoAdapter := cryptoadapters.New(configAll.API)
	cryptoRepo := cryptorepositories.New(localstorage, cryptoAdapter)
	cryptoUsecase := cryptousecases.New(cryptoAdapter, cryptoRepo)
	cryptoHandler := cryptohandlers.New(cryptoUsecase)
	cryptoUsecase.InitCoinMapping()

	cryptohandlers.InitHandleCrypto(cryptoHandler, mux, configAll.JWT)
	usershandlers.InitHandleUser(authHandler, mux)

	// for i, item := range localstorage.NamingMapSymbolID {
	// 	fmt.Println(i, "        ", item)
	// }
	// 
	ctx, _ := context.WithCancel(context.Background())
	go scheduler.UpdatePrices(ctx, cryptoUsecase)

	func() {
		err := http.ListenAndServe(":8080", mux)
		if err != nil {
			fmt.Println(err)
			return
		}
	}()

}
