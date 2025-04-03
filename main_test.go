package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandleRPC(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		requestBody    string
		expectedStatus int
		mockResponse   string
	}{
		{
			name:           "valid eth_blockNumber",
			method:         "POST",
			requestBody:    `{"jsonrpc":"2.0","method":"eth_blockNumber","id":2}`,
			expectedStatus: http.StatusOK,
			mockResponse:   `{"jsonrpc":"2.0","id":2,"result":"0x1234"}`,
		},
		{
			name:           "invalid method",
			method:         "POST",
			requestBody:    `{"jsonrpc":"2.0","method":"invalid_method","id":2}`,
			expectedStatus: http.StatusForbidden,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte(tt.mockResponse))
			}))
			defer mockServer.Close()
			rpcURL = mockServer.URL

			req := httptest.NewRequest(tt.method, "/", strings.NewReader(tt.requestBody))
			req.Header.Set("Content-Type", "application/json")

			rr := httptest.NewRecorder()
			handleRPC(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, rr.Code)
			}
		})
	}
}