package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandlers(t *testing.T) {
	tests := []struct {
		name           string
		handler        http.HandlerFunc
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Home Handler",
			handler:        home,
			expectedStatus: http.StatusOK,
			expectedBody:   "Welcome to the Shapes API",
		},
		{
			name:           "Health Handler",
			handler:        health,
			expectedStatus: http.StatusOK,
			expectedBody:   "Server is running",
		},
		{
			name:           "About Handler",
			handler:        about,
			expectedStatus: http.StatusOK,
			expectedBody:   "Anya Andrews",
		},
		{
			name:           "Time Handler",
			handler:        getServerTime,
			expectedStatus: http.StatusOK,
			expectedBody:   "",
		},
		{
			name:           "Random Handler",
			handler:        random,
			expectedStatus: http.StatusOK,
			expectedBody:   "Random Number:",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/", nil)
			rr := httptest.NewRecorder()

			handler := http.HandlerFunc(tt.handler)
			handler.ServeHTTP(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					rr.Code, tt.expectedStatus)
			}

			if !strings.Contains(rr.Body.String(), tt.expectedBody) {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), tt.expectedBody)
			}
		})
	}
}
