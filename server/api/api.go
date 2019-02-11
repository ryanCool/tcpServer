package api

import (
	"encoding/json"
	"fmt"
	"github.com/ryanCool/tcpServer/server/tcp"
	"net/http"
)

type Stat struct {
	CurrConnCount      int      `json:"current_connection_count"`
	CurrReqRate        int      `json:"current_request_rate"`
	CloseConnections   int      `json:"close_connections"`
	MessageReceiveRate int      `json:"message_receive_rate"`
	MessageSentRate    int      `json:"message_sent_rate"`
	DataReceiveRate    int      `json:"data_receive_rate"`
	DataSentRate       int      `json:"data_sent_rate"`
	ProcessedReq       int      `json:"processed_request_count"`
	ConnClients        []string `json:"connected_client"`
	RemainingJobs      int      `json:"remaining_jobs"`
}

const (
	apiServePort = "8000"
)

func StartServer() {
	http.HandleFunc("/healthy", HealthHandler)
	http.HandleFunc("/stat", StatHandler)

	fmt.Printf("HTTP Server listening on port:%s\n", apiServePort)
	err := http.ListenAndServe(":"+apiServePort, nil)
	if err != nil {
		fmt.Printf("http server listen err%v\n", err)
	}
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	response := struct {
		Message string `json:"message"`
	}{Message: "healthy"}
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		fmt.Println("json new encoder fail")
	}
}

func StatHandler(w http.ResponseWriter, r *http.Request) {
	stat := &Stat{
		CurrConnCount:      len(tcp.ConnClients),
		CurrReqRate:        tcp.CurrReqRate,
		CloseConnections:   tcp.CloseConnections,
		MessageReceiveRate: tcp.MessageReceiveRate,
		MessageSentRate:    tcp.MessageSentRate,
		DataReceiveRate:    tcp.DataReceiveRate,
		DataSentRate:       tcp.DataSentRate,
		ProcessedReq:       tcp.ProcessedReq,
		ConnClients:        tcp.ConnClients,
		RemainingJobs:      tcp.RemainingJobs,
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(stat)
	if err != nil {
		fmt.Println("json new encoder fail")
	}
}
