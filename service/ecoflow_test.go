package service

import (
	"go-ecoflow-api-server/constants"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetEcoflowClient(t *testing.T) {
	tests := []struct {
		name          string
		mockRequest   *http.Request
		expectedError bool
	}{
		{
			name: "valid tokens",
			mockRequest: func() *http.Request {
				req := httptest.NewRequest("GET", "/", nil)
				req.Header.Set(constants.HeaderAuthorization, "Bearer valid-token")
				req.Header.Set(constants.HeaderXSecretToken, "valid-secret-token")
				return req
			}(),
			expectedError: false,
		},
		{
			name: "missing access token",
			mockRequest: func() *http.Request {
				req := httptest.NewRequest("GET", "/", nil)
				req.Header.Set(constants.HeaderXSecretToken, "valid-secret-token")
				return req
			}(),
			expectedError: true,
		},
		{
			name: "invalid access token",
			mockRequest: func() *http.Request {
				req := httptest.NewRequest("GET", "/", nil)
				req.Header.Set(constants.HeaderAuthorization, "no bearer string")
				return req
			}(),
			expectedError: true,
		},
		{
			name: "invalid secret token",
			mockRequest: func() *http.Request {
				req := httptest.NewRequest("GET", "/", nil)
				req.Header.Set("Authorization", "Bearer invalid-token")
				return req
			}(),
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			client, err := GetEcoflowClient(tt.mockRequest)
			if tt.expectedError {
				if err == nil {
					t.Errorf("expected error but got nil")
				}
				if client != nil {
					t.Errorf("expected client to be nil, got: %v", client)
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
			}
		})
	}
}
