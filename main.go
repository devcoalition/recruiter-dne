package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/api", notImplementedHandler)

	s := &http.Server{
		Handler:      r,
		Addr:         ":4000",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Fatalln(s.ListenAndServe())
}

func notImplementedHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("success"))
}
