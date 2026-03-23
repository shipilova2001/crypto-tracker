package usecases

func (cryptoUC *cryptoUseCase) InitCoinMapping () {
	cryptos := cryptoUC.adapters.GetList()
	cryptoUC.repositories.SetNamingMapSymbolID(cryptos)
}
