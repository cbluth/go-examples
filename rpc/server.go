package main

import (
	"errors"
	"net/rpc"
	"log"
)

type DB struct {
	m map[string]string
	// TODO: should this have a mutex?
}

var (
	db = DB{m:map[string]string{}}
	rServer = &rpc.Server{}
)

func server() error {
	log.Printf("Starting %s in %s-mode", name, mode)
	log.Printf("Using %s localhost connection on %s", transport, connection)
	switch transport {
		case "tcp" : {
			return TCPServer()
		}
		case "http" : {
			return HTTPServer()
		}
		default : {
			return errors.New("transport not supported: "+transport)
		}
	}
}

func makeRPCServer() *rpc.Server {
	r := rpc.NewServer()
	err := r.Register(&DB{})
	if err != nil {
		log.Fatalln(err)
	}
	return r
}

func (d *DB) Ping(r []byte, w *[]byte) error {
	if r == nil || string(r) == "" {
		return errors.New("empty request")
	}
	*w = r
	log.Println("Ping from client >", string(r))
	return nil
}

func (d *DB) Get(k string, res *string) error {
	if v, ok := db.m[k] ; ok {
		*res = v
		log.Println("GET:", k, "==", v)
		return nil
	}
	return errors.New("key not found: "+k)	
}

func (d *DB) Delete(k string, ok *bool) error {
	if v, o := db.m[k] ; o {
		delete(db.m, k)
		log.Println("DELETE:", k, "==", v)
		*ok = true
		return nil
	}
	return errors.New("key not found: "+k)	
}

// TODO: redo this?
func (d *DB) Put(m map[string]string, ok *bool) error {
	k, v  := "", ""
	for k, v = range m {
		db.m[k] = v
		log.Println("PUT:", k, "==", v)
		*ok = true
	}
	if *ok {
		return nil
	}
	return errors.New("key not found: "+k)
}

func (d *DB) List(_ int, keys *[]string) error {
	ok := false
	for k := range db.m {
		*keys = append(*keys, k)
		ok = true
	}
	if ok {
		log.Println("LIST:", *keys)
		return nil
	}
	return errors.New("keys not found")	
}
