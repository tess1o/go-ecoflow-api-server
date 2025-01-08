package handlers

import (
	"errors"
	"github.com/go-chi/httplog/v2"
	"github.com/tess1o/go-ecoflow"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"go-ecoflow-api-server/constants"
)

func TestBaseHandler_GetEcoflowClientOrRespondWithError(t *testing.T) {
	tests := []struct {
		name           string
		providerFunc   func(r *http.Request) (*ecoflow.Client, error)
		expectedStatus int
		expectSuccess  bool
	}{
		{
			name: "valid client",
			providerFunc: func(r *http.Request) (*ecoflow.Client, error) {
				return &ecoflow.Client{}, nil
			},
			expectedStatus: http.StatusOK,
			expectSuccess:  true,
		},
		{
			name: "invalid client error",
			providerFunc: func(r *http.Request) (*ecoflow.Client, error) {
				return nil, errors.New("invalid client")
			},
			expectedStatus: http.StatusUnauthorized,
			expectSuccess:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger := httplog.NewLogger("test-logger", httplog.Options{}) // logger instance
			handler := NewBaseHandler(logger, tt.providerFunc)

			req, err := http.NewRequest(http.MethodGet, "/", nil)
			assert.NoError(t, err)

			// Mock the response writer
			rec := httptest.NewRecorder()

			client, success := handler.GetEcoflowClientOrRespondWithError(req, rec)

			if tt.expectSuccess {
				assert.NotNil(t, client)
				assert.True(t, success)
				assert.Equal(t, tt.expectedStatus, rec.Result().StatusCode)
			} else {
				assert.Nil(t, client)
				assert.False(t, success)
				assert.Equal(t, tt.expectedStatus, rec.Result().StatusCode)

				// Verify response body for error case
				respBody := rec.Body.String()
				assert.Contains(t, respBody, constants.ErrInvalidAuthHeader)
			}
		})
	}
}
