package tcp

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
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
	readCount := 0
	for scanner.Scan() {
		if err != nil {
			t.Errorf("err=%v", err)
			return
		}

		if scanner.Text() != "quit" {
			t.Errorf("should receive quit from server, but receive %v", scanner.Text())
			return
		} else {
			readCount++
			break
		}
	}

	if readCount == 0 {
		t.Errorf("server should send quit")
	}
}

func TestHandleMessageQuitTimeout(t *testing.T) {
	conn, err := net.Dial("tcp", "localhost:"+testPort)
	if err != nil {
		t.Errorf("err=%v", err)
		return
	}
	defer conn.Close()

	time.Sleep(3 * time.Second)

	_, err = conn.Write([]byte("quit\n"))
	if err != nil {
		t.Errorf("err=%v", err)
		return
	}
	scanner := bufio.NewScanner(conn)
	readCount := 0
	for scanner.Scan() {
		if err != nil {
			t.Errorf("err=%v", err)
			return
		}

		if scanner.Text() != "quit" {
			t.Errorf("should receive quit from server, but receive %v", scanner.Text())
			return
		} else {
			readCount++
			break
		}
	}

	if readCount != 0 {
		t.Errorf("server should timeout and receive nothing")
	}
}

func TestHandleMessageExternalAPI(t *testing.T) {
	conn, err := net.Dial("tcp", "localhost:"+testPort)
	if err != nil {
		t.Errorf("err=%v", err)
		return
	}
	defer conn.Close()
	_, err = conn.Write([]byte("黑色\n"))
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

		//We expect at least contain one result of black cat in query list
		//TODO : mock QueryCatByColor response
		if !strings.Contains(scanner.Text(), `"animal_colour":"黑色"`) {
			t.Errorf("external api result is not as expected %v", scanner.Text())
			return
		} else {
			break
		}
	}

}
