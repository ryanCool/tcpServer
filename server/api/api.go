package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Stat struct {
	CurrConnCount      int      `json:"current_connection_count"`
	CurrReqRate        float64  `json:"current_request_rate"`
	CurrReqCount       int      `json:"current_request_count"`
	CloseConnections   int      `json:"close_connections"`
	MessageReceiveRate float64  `json:"message_receive_rate"`
	MessageSentRate    float64  `json:"message_sent_rate"`
	DataReceiveRate    float64  `json:"data_receive_rate"`
	DataSentRate       float64  `json:"data_sent_rate"`
	ProcessedReq       int      `json:"processed_request_count"`
	ConnClient         []string `json:"connected_client"`
	CurrGoRoutine      int      `json:"current_goroutine_count"`
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
	stat := &Stat{}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(stat)
	if err != nil {
		fmt.Println("json new encoder fail")
	}
}
