package main

import (
	"flag"
)

const (
	quitCommand = "quit"
	network     = "tcp"
)

var (
	destPort = flag.Int("port", 8000, "tcp dial port")
	destHost = flag.String("host", "localhost", "tcp dial host")
)

func main() {
}
