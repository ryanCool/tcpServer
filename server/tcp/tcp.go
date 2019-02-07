package tcp

import (
	"github.com/honestbeeHomeTest/externalAPI"

	"bufio"
	"fmt"
	"net"
	"time"
)

const (
	quitCommand = "quit"
)

var (
	readTimeoutDuration = 10 * time.Second
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
		//refresh read timeout
		err := conn.SetReadDeadline(time.Now().Add(readTimeoutDuration))
		if err != nil {
			fmt.Println("set dead line fail :", err)
			return
		}

		err = handleMessage(scanner.Text(), conn)
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
	} else {
		result, err := externalAPI.QueryCatColor(message)
		if err != nil {
			return err
		}
		_, err = conn.Write([]byte(result))
		if err != nil {
			return err
		}
	}
	return nil
}
