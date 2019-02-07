package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"strconv"
	"sync"
)

const (
	quitCommand = "quit"
	network     = "tcp"
)

var (
	destPort     = flag.Int("port", 3000, "tcp dial port")
	destHost     = flag.String("host", "localhost", "tcp dial host")
	workerNumber = flag.Int("wn", 50, "Concurrent client")
	queryParams  = []string{"黑色", "灰色", "白色", "quit"}
)

func main() {
	flag.Parse()
	var wg sync.WaitGroup
	for w := 1; w <= *workerNumber; w++ {
		wg.Add(1)
		go requestWorker(&wg, w)
	}

	wg.Wait()
	fmt.Println("test finish")
}

func requestWorker(wg *sync.WaitGroup, n int) {
	fmt.Printf("worker : %d prepare connect\n", n)
	conn, err := net.Dial(network, *destHost+":"+strconv.Itoa(*destPort))
	if err != nil {
		fmt.Printf("work: %d dial fail :%v \n", n, err)
		return
	}

	defer conn.Close()
	for _, query := range queryParams {
		fmt.Printf("worker : %d query %s\n", n, query)

		_, err = conn.Write([]byte(query + "\n"))
		if err != nil {
			fmt.Printf("work: %d write fail :%v ", n, err)
			continue
		}
		scanner := bufio.NewScanner(conn)
		if scanner.Scan() {
			fmt.Printf("worker : %d query done %s\n", n, scanner.Text())
			if scanner.Text() == "quit" {
				break
			}
		}

	}
	wg.Done()
}
