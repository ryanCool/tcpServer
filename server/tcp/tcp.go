package tcp

import (
	"flag"
	"net"
)

const (
	quitCommand = "quit"
	network     = "tcp"
)

var (
	port = flag.Int("port", 8000, "tcp server listen port")
)

func HandleConnection(conn net.Conn) {

}

func handleMessage(message string, conn net.Conn) {

}
