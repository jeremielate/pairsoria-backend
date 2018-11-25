package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	apiModule "pairsoria.com/server/api"
	chatModule "pairsoria.com/server/chat"
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
	hub := chatModule.NewHub()

	r := mux.NewRouter()
	r.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		chatModule.ServeWs(hub, w, r)
	})
	api.Route(r)

	srv := http.Server{
		Addr:    c.Address,
		Handler: r,
	}

	if err := srv.ListenAndServe(); err != nil {
		logrus.Fatalln(err)
	}
}
