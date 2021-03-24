package main

import (
	"log"
	"net/http"
)

func HTTPServer() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		log.Printf("Client %s connected.", r.RemoteAddr)
		rServer.ServeHTTP(w, r)
	})
	return http.ListenAndServe("127.0.0.1:"+connection, mux)
}

