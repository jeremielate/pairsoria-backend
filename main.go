package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

const (
	DefaultConfigFile = "config.toml"
)

type MainHandler struct {
	redisClient *redis.Client
}

// page principale
func (h MainHandler) rootPage(w http.ResponseWriter, r *http.Request) {
	logrus.Infof("%v %v %v %v\n", r.Method, r.Proto, r.URL.String(), r.RemoteAddr)
	http.Error(w, "not found", http.StatusNotFound)
}

// page profil, GET pour recuperer le profil, POST pour le mettre a jour
func (h MainHandler) profile(w http.ResponseWriter, r *http.Request) {
	logrus.Infof("%v %v %v %v\n", r.Method, r.Proto, r.URL.String(), r.RemoteAddr)
	switch r.Method {
	case "GET":
		id := r.URL.Query().Get("id")
		list, err := h.redisClient.HGetAll(id).Result()
		if err != nil {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}
		if err := json.NewEncoder(w).Encode(list); err != nil {
			logrus.Warningln(err)
		}
	case "POST":
	default:
		http.Error(w, "bad request", http.StatusBadRequest)
	}

}

func NewMainHandler(c *Config) MainHandler {
	return MainHandler{
		redisClient: redis.NewClient(c.RedisOptions()),
	}
}

func main() {
	logrus.SetReportCaller(true)

	c, err := ReadConfig(DefaultConfigFile)
	if err != nil {
		logrus.Fatalln(err)
	}

	mh := NewMainHandler(c)

	mux := http.ServeMux{}
	mux.HandleFunc("/", mh.rootPage)
	mux.HandleFunc("/profile", mh.profile)

	srv := http.Server{
		Addr:    c.Address,
		Handler: &mux,
	}

	if err := srv.ListenAndServe(); err != nil {
		logrus.Fatalln(err)
	}
}
