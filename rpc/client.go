package main

import (
	"bytes"
	"log"
	"sort"

	// "log"
	"bufio"
	"errors"
	"fmt"
	"net/rpc"
	"os"
	"strings"
)

type (
	// mClient represents an rpc
	// client for multiple transports
	mClient struct {
		*rpc.Client
	}
)

var (
	rClient = mClient{&rpc.Client{}}
)

func client() error {
	stdin := bufio.NewReader(os.Stdin)
	stdout := bufio.NewWriter(os.Stdout)
	err := (error)(nil)
	for {
		err = nil
		_, err = stdout.WriteString("rpc > ")
		if err != nil {
			break
		}
		err = stdout.Flush()
		if err != nil {
			break
		}
		line, err := stdin.ReadString('\n')
		if err != nil {
			break
		}
		cmd := strings.Fields(line)
		switch cmd[0] {
			case "ping" : {
				err = rClient.ping()
			}
			case "put" : {
				err = rClient.put(cmd[1], cmd[2])
			}
			case "get" : {
				err =  rClient.get(cmd[1])
			}
			case "del" : {
				err =  rClient.delete(cmd[1])
			}
			case "ls", "list" : {
				err =  rClient.list()
			}
			default : {
				fmt.Printf("command %s not supported\n", cmd[0])
			}
		}
		if err != nil {
			log.Println(err)
		}
	}
	return err
}

func (c *mClient) init(transport, conn string) error {
	err := (error)(nil)
	conn = strings.TrimPrefix(conn, ":")
	switch transport {
		case "tcp" : {
			c.Client, err = rpc.Dial(transport, "127.0.0.1:"+conn)
			if err != nil { return err }
			return nil
		}
		case "http" : {
			c.Client, err = rpc.DialHTTP("tcp", "127.0.0.1:"+conn)
			if err != nil { return err }
			return nil
		}
		default : {
			return errors.New("transport "+transport+" not supported")
		}
	}
	// return errors.New("transport "+transport+" not supported")
}

func (c *mClient) ping() error {
	pong := []byte("pong")
	res := []byte{}
	err := c.Call("DB.Ping", pong, &res)
	if err != nil { return err }
	if bytes.Equal(res, pong) {
		fmt.Println(string(res))
		return nil
	}
	return errors.New("ping failed")
}

func (c *mClient) put(k, v string) error {
	res := false
	err := c.Call("DB.Put", map[string]string{k:v}, &res)
	if err != nil { return err }
	if res {
		fmt.Println("OK")
		return nil
	}
	return errors.New("ping failed")
}

func (c *mClient) get(get string) error {
	res := ""
	err := c.Call("DB.Get", get, &res)
	if err != nil { return err }
	if res != "" {
		fmt.Println(res)
		return nil
	}
	return errors.New("get failed")
}

func (c *mClient) delete(del string) error {
	res := false
	err := c.Call("DB.Delete", del, &res)
	if err != nil { return err }
	if res {
		fmt.Println("OK")
		return nil
	}
	return errors.New("delete failed")
}

func (c *mClient) list() error {
	res := []string{}
	err := c.Call("DB.List", 0, &res)
	if err != nil { return err }
	if len(res) > 0 {
		sort.Strings(res)
		fmt.Println("Key List:")
		for _, k := range res {
			fmt.Println("-", k)
		}
		return nil
	}
	return errors.New("list failed")
}
