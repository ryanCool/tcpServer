package tcp

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"time"
)

const (
	quitCommand = "quit"
	network     = "tcp"
)

var (
	readTimeoutDuration = 10 * time.Second
	port                = flag.Int("port", 8000, "tcp server listen port")
	ErrClientCloseConn  = fmt.Errorf("client close connection")
)

func HandleConnection(conn net.Conn) {

	remoteAddr := conn.RemoteAddr().String()
	fmt.Println("Client connected from " + remoteAddr)

	defer func() {
		err := conn.Close()
		if err != nil {
			fmt.Println("conn close fail:", err)
		}
		fmt.Println("Client at " + remoteAddr + " disconnected.")
	}()

	err := conn.SetReadDeadline(time.Now().Add(readTimeoutDuration))
	if err != nil {
		fmt.Println("set dead line fail :", err)
		return
	}

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		err := handleMessage(scanner.Text(), conn)
		if err != nil && err != ErrClientCloseConn {
			fmt.Println("handleMessage fail", err)
			return
		} else if err == ErrClientCloseConn {
			fmt.Println("client close connection")
			return
		}

	}

	if err := scanner.Err(); err != nil {
		fmt.Println("scanner err:", err)
	}
	return
}

func handleMessage(message string, conn net.Conn) error {
	if message == "quit" {
		_, err := conn.Write([]byte(quitCommand + "\n"))
		if err != nil {
			return err
		}
		return ErrClientCloseConn
	}
	return nil
}
