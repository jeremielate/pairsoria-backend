package main

import (
	"log"
	"net/http"
)

func rootPage(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "not found", http.StatusNotFound)
}

func main() {
	mux := http.ServeMux{}

	mux.HandleFunc("/", rootPage)
	srv := http.Server{Handler: &mux}

	if err := srv.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
