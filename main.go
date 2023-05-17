package main

import (
	"fmt"
	"nickPay/wallet/config"
	"nickPay/wallet/internal/controller"
	"nickPay/wallet/server"
	"strconv"

	logrus "github.com/sirupsen/logrus"
	negroni "github.com/urfave/negroni"
)

func main() {
	config.Load()

	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	deps, err := server.InitDependencies()
	if err != nil {
		logrus.WithError(err).Error("Failed to initialize dependencies")
		return
	}

	router := controller.InitRouter(deps)
	server := negroni.Classic()
	server.UseHandler(router)

	port := config.AppPort()
	address := fmt.Sprintf(":%s", strconv.Itoa(port))
	server.Run(address)
}
