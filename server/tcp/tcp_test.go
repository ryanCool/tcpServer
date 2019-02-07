package tcp

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"testing"
	"time"
)

const (
	testPort = "3000"
)

var (
	createDone = make(chan int)
)

func TestMain(m *testing.M) {
	go CreateTCPServer()
	<-createDone
	os.Exit(m.Run())
}

func CreateTCPServer() {
	listener, err := net.Listen("tcp", ":"+testPort)
	if err != nil {
		panic(err)
	}
	readTimeoutDuration = 2 * time.Second
	createDone <- 1
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Some connection error: %s\n", err)
		}

		go HandleConnection(conn)
	}

}

func TestHandleMessageQuit(t *testing.T) {
	conn, err := net.Dial("tcp", "localhost:"+testPort)
	if err != nil {
		t.Errorf("err=%v", err)
		return
	}
	defer conn.Close()
	_, err = conn.Write([]byte("quit\n"))
	if err != nil {
		t.Errorf("err=%v", err)
		return
	}

	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		if err != nil {
			t.Errorf("err=%v", err)
			return
		}

		if scanner.Text() != "quit" {
			t.Errorf("should receive quit from server, but receive %v", scanner.Text())
			return
		} else {
			break
		}
	}

}
