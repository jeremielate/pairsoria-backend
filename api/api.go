package api

import (
	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	configModule "pairsoria.com/server/config"
)

type Api struct {
	redisClient *redis.Client
}

func NewApi(c *configModule.Config) Api {
	return Api{
		redisClient: redis.NewClient(c.RedisOptions()),
	}
}

func (a Api) Route(mux *mux.Router) {
	mux.HandleFunc("/", a.index)
	mux.HandleFunc("/profile", a.listProfile).Methods("GET")
	mux.HandleFunc("/profile", a.newProfile).Methods("POST")
	mux.HandleFunc("/profile/{id}", a.getProfile).Methods("GET")
	mux.HandleFunc("/profile/{id}", a.updateProfile).Methods("PUT")
}
