package api

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

func (api Api) listProfile(w http.ResponseWriter, r *http.Request) {
	logrus.Infof("%v %v %v %v\n", r.Method, r.Proto, r.URL.String(), r.RemoteAddr)
	id := r.URL.Query().Get("id")
	list, err := api.redisClient.HGetAll(id).Result()
	if err != nil {
		logrus.Warningln(err)
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	header := w.Header()
	header.Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(list); err != nil {
		logrus.Warningln(err)
	}
}

func (api Api) getProfile(w http.ResponseWriter, r *http.Request) {
	logrus.Infof("%v %v %v %v\n", r.Method, r.Proto, r.URL.String(), r.RemoteAddr)
	id := r.URL.Query().Get("id")
	list, err := api.redisClient.HGetAll(id).Result()
	if err != nil {
		logrus.Warningln(err)
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	header := w.Header()
	header.Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(list); err != nil {
		logrus.Warningln(err)
	}
}

func (api Api) updateProfile(w http.ResponseWriter, r *http.Request) {
	var postData struct {
		Id   string            `json:"id"`
		Data map[string]string `json:"data"`
	}
	if err := json.NewDecoder(r.Body).Decode(&postData); err != nil {
		logrus.Warningln(err)
		http.Error(w, "sever error", http.StatusInternalServerError)
		return
	}
	for k, v := range postData.Data {
		ok, err := api.redisClient.HSet(postData.Id, k, v).Result()
		if err != nil {
			logrus.Warningln(err)
		} else if !ok {
			logrus.Infof("can't hset %v -> %v\n", k, v)
		}
	}
}

func (api Api) newProfile(w http.ResponseWriter, r *http.Request) {
	var postData struct {
		Id   string            `json:"id"`
		Data map[string]string `json:"data"`
	}
	if err := json.NewDecoder(r.Body).Decode(&postData); err != nil {
		logrus.Warningln(err)
		http.Error(w, "sever error", http.StatusInternalServerError)
		return
	}
	for k, v := range postData.Data {
		ok, err := api.redisClient.HSet(postData.Id, k, v).Result()
		if err != nil {
			logrus.Warningln(err)
		} else if !ok {
			logrus.Infof("can't hset %v -> %v\n", k, v)
		}
	}
}
