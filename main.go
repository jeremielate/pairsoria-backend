package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	apiModule "pairsoria.com/server/api"
	configModule "pairsoria.com/server/config"
)

const (
	DefaultConfigFile = "config.toml"
)

func main() {
	logrus.SetReportCaller(true)

	c, err := configModule.ReadConfig(DefaultConfigFile)
	if err != nil {
		logrus.Fatalln(err)
	}

	api := apiModule.NewApi(c)

	r := mux.NewRouter()
	api.Route(r)

	srv := http.Server{
		Addr:    c.Address,
		Handler: r,
	}

	if err := srv.ListenAndServe(); err != nil {
		logrus.Fatalln(err)
	}
}
