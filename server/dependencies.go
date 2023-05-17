package server

import (
	"nickPay/wallet/internal/db"
	"nickPay/wallet/internal/service"
)


type Dependencies struct {
	NikPay service.WalletService
}

func InitDependencies() (deps *Dependencies, err error) {
	store, err := db.Init()
	if err != nil {
		return
	}
	return &Dependencies{
		NikPay: service.NewWalletService(store),
	}, nil
}
