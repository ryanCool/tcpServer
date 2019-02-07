package main

/*
	Create a tcp server which can takes in any request text per line and send a query to an external API
	server can accept multiple connections at the same time.
	client can send quit to close connection.
*/

import (
	"flag"
	"fmt"
	"net"
	"strconv"

	"github.com/honestbeeHomeTest/server/tcp"
)

var (
	port    = flag.Int("port", 8000, "tcp server listen port")
	network = "tcp"
)

func main() {
	flag.Parse()

	listener, err := net.Listen(network, ":"+strconv.Itoa(*port))
	if err != nil {
		panic(err)
		return
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Some connection error: %s\n", err)
		}

		go tcp.HandleConnection(conn)
	}
}
