package main

import (
	// "bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
)

var (
	name = getName()
	mode = ""
	transport = ""
	connection = ""
)

func main() {
	parseArgs()
	err := (error)(nil)
	switch mode {
		case "server" : {
			err = server()
		}
		case "client" : {
			err = client()
		}
		default : {
			err = errors.New("mode not supported: "+mode)
		}
	}
	if err != nil {
		log.Fatalln(err)
	}
}

func parseArgs() {
	mode = os.Args[1]
	transport = strings.TrimPrefix(os.Args[2], "-")
	connection = strings.TrimPrefix(os.Args[3], ":")
	if !(transport == "tcp" || transport == "http") {
		log.Fatalln("transport not supported: " + transport)
	}
	if mode == "client" {
		err := rClient.init(transport, connection)
		if err != nil {
			log.Fatalln(err)
		}
	}
	if mode == "server" {
		rServer = makeRPCServer()
	}
}

func getName() string {
	_, name := path.Split(os.Args[0])
	if len(os.Args) != 4 {
		log.Fatalln(fmt.Sprintf(`
Please use:
$ %s <mode> -<transport> <connection>
eg:
- %s server -tcp :1234
- %s client -tcp :1234
or
- %s server -http :1234
- %s client -http :1234`, name, name, name, name, name))
	}
	return name
}
