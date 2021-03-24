package main

// import (
// 	"encoding/gob"
// 	"bytes"
// 	"net"
// 	"log"
// )


// func UDPServer() error {
// 	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:"+connection)
// 	if err != nil { return err }
// 	l, err := net.ListenUDP("udp", addr)
// 	if err != nil { return err }
// 	defer l.Close()
// 	for {
// 		b := make([]byte, 1024)
// 		i, ad, err := l.ReadFrom(b)
// 		go func(){
// 			log.Printf("Client %s connected!", ad.String())
// 			// log.Println(string(b[:i]))
// 			if err != nil {
// 				log.Println(err)
// 			}
			
// 			dec := gob.NewDecoder(bytes.NewBuffer(b[:i]))
// 			dec.Decode(inn)
// 			log.Println(inn)
// 		}()
		
// 	}
// 	// rServer.Accept(l)
// 	// conn, err := net.li
// 	return nil
// }

// func UDPClient() error {
// 	rpcCMD := readSTDIN()
// 	switch rpcCMD[0] {
// 		case "ping" : {
// 			return rClient.ping()
// 		}
// 		case "put" : {
// 			return rClient.put(rpcCMD[1], rpcCMD[2])
// 		}
// 		case "get" : {
// 			return rClient.get(rpcCMD[1])
// 		}
// 		case "del" : {
// 			return rClient.delete(rpcCMD[1])
// 		}
// 		case "ls", "list" : {
// 			return rClient.list()
// 		}
// 		default : {}
// 	}
// 	return nil
// }

