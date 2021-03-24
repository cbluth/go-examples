package main

import (
	"log"
	"net"
	"net/http"
	"strings"
)

func HTTPServer() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		log.Printf("Client %s connected.", r.RemoteAddr)
		rServer.ServeHTTP(w, r)
	})
	return http.ListenAndServe("127.0.0.1:"+connection, mux)
}

func HTTPServerOLD(conn string) error {
	conn = strings.TrimPrefix(conn, ":")
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:"+conn)
	if err != nil { return err }
	l, err := net.ListenTCP("tcp", addr)
	if err != nil { return err }
	defer l.Close()
	rServer.Accept(l)
	return nil
}
