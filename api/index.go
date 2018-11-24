package api

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

// page principale
func (h Api) index(w http.ResponseWriter, r *http.Request) {
	logrus.Infof("%v %v %v %v\n", r.Method, r.Proto, r.URL.String(), r.RemoteAddr)
	http.Error(w, "not found", http.StatusNotFound)
}
