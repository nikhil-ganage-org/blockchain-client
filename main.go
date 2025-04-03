package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

var rpcURL = "https://polygon-rpc.com/"

type JsonRPCRequest struct {
	JsonRPC string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	ID      int           `json:"id"`
}

func handleRPC(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var reqBody JsonRPCRequest
	if err := json.Unmarshal(bodyBytes, &reqBody); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	allowedMethods := map[string]bool{
		"eth_blockNumber":        true,
		"eth_getBlockByNumber":   true,
	}

	if !allowedMethods[reqBody.Method] {
		http.Error(w, "Method not allowed", http.StatusForbidden)
		return
	}

	proxyReq, err := http.NewRequest(http.MethodPost, rpcURL, bytes.NewReader(bodyBytes))
	if err != nil {
		http.Error(w, "Error creating request", http.StatusInternalServerError)
		return
	}
	proxyReq.Header = r.Header.Clone()
	proxyReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(proxyReq)
	if err != nil {
		http.Error(w, "Error forwarding request", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}

func main() {
	http.HandleFunc("/", handleRPC)
	log.Println("Server starting on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}