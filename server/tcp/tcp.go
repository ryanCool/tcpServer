package tcp

import (
	"github.com/honestbeeHomeTest/externalAPI"

	"bufio"
	"fmt"
	"net"
	"sync"
	"time"
)

const (
	quitCommand = "quit"
)

var (
	lock                sync.RWMutex
	readTimeoutDuration = 10 * time.Second
	ErrClientCloseConn  = fmt.Errorf("client close connection")
	ConnClients         = []string{}
	CurrReqRate         = 0
	CloseConnections    = 0
	MessageReceiveRate  = 0
	MessageSentRate     = 0
	DataReceiveRate     = 0
	DataSentRate        = 0
	ProcessedReq        = 0
	RemainingJobs       = 0

	SecReq         = 0
	SecMessageRcv  = 0
	SecMessageSent = 0
	SecDataRcv     = 0
	SecDataSent    = 0

	RateTicker = time.Tick(1 * time.Second)
)

func init() {
	go func() {
		for {
			<-RateTicker
			statis()
		}
	}()
}

func statis() {
	lock.Lock()
	CurrReqRate = SecReq
	SecReq = 0
	MessageSentRate = SecMessageSent
	SecMessageSent = 0
	MessageReceiveRate = SecMessageRcv
	SecMessageRcv = 0
	DataReceiveRate = SecDataRcv
	SecDataRcv = 0
	DataSentRate = SecDataSent
	SecDataSent = 0
	lock.Unlock()
}

func HandleConnection(conn net.Conn) {
	remoteAddr := conn.RemoteAddr().String()
	fmt.Println("Client connected from " + remoteAddr)
	lock.Lock()
	ConnClients = append(ConnClients, conn.RemoteAddr().String())
	lock.Unlock()

	defer func() {
		CloseConnections++
		err := conn.Close()
		if err != nil {
			fmt.Println("conn close fail:", err)
		}

		lock.Lock()
		for i, client := range ConnClients {
			if client == conn.RemoteAddr().String() {
				ConnClients = append(ConnClients[:i], ConnClients[i+1:]...)
			}
		}
		lock.Unlock()
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
	lock.Lock()
	SecMessageRcv++
	SecDataRcv = SecDataRcv + len([]byte(message))
	lock.Unlock()
	if message == "quit" {
		_, err := conn.Write([]byte(quitCommand + "\n"))
		if err != nil {
			return err
		}
		lock.Lock()
		SecMessageSent++
		SecDataSent = SecDataSent + len([]byte(quitCommand))
		lock.Unlock()
		return ErrClientCloseConn
	} else {
		lock.Lock()
		SecReq++
		RemainingJobs++
		lock.Unlock()
		<-externalAPI.RateLimit
		result, err := externalAPI.QueryCatByColor(message)
		if err != nil {
			return err
		}
		_, err = conn.Write([]byte(result + "\n"))
		if err != nil {
			return err
		}
		lock.Lock()
		ProcessedReq++
		SecMessageSent++
		RemainingJobs--
		SecDataSent = SecDataSent + len([]byte(result))
		lock.Unlock()
	}
	return nil
}
