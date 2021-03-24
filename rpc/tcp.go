package main

import (
	"log"
	"net"
)

func TCPServer() error {
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:"+connection)
	if err != nil { return err }
	l, err := net.ListenTCP("tcp", addr)
	if err != nil { return err }
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println(err)
		} else {
			go func(){
				log.Printf("Client %s connected!", conn.RemoteAddr())
				rServer.ServeConn(conn)
			}()
		}
	}
}
