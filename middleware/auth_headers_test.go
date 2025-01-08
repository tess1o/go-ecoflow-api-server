package middleware

import (
	"go-ecoflow-api-server/constants"
	"go-ecoflow-api-server/handlers"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCheckAuthHeaders(t *testing.T) {
	tests := []struct {
		name           string
		headers        map[string]string
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "all headers present",
			headers: map[string]string{
				"Header1": "Value1",
				"Header2": "Value2",
			},
			expectedStatus: http.StatusOK,
			expectedBody:   "",
		},
		{
			name:           "no headers present",
			headers:        map[string]string{},
			expectedStatus: http.StatusUnauthorized,
			expectedBody:   constants.ErrMandatoryHeaderMissing,
		},
		{
			name: "one mandatory header missing",
			headers: map[string]string{
				"Header1": "Value1",
			},
			expectedStatus: http.StatusUnauthorized,
			expectedBody:   constants.ErrMandatoryHeaderMissing,
		},
		{
			name: "headers with empty values",
			headers: map[string]string{
				"Header1": "",
				"Header2": "",
			},
			expectedStatus: http.StatusUnauthorized,
			expectedBody:   constants.ErrMandatoryHeaderMissing,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			baseHandler := &handlers.BaseHandler{}
			authHeaders := []string{"Header1", "Header2"}
			middleware := NewAuthHeadersMiddleware(baseHandler, authHeaders)

			nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			})

			testHandler := middleware.CheckAuthHeaders(nextHandler)

			req := httptest.NewRequest(http.MethodGet, "/", nil)
			for key, value := range tt.headers {
				req.Header.Set(key, value)
			}

			recorder := httptest.NewRecorder()

			testHandler.ServeHTTP(recorder, req)

			if recorder.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, recorder.Code)
			}

			if tt.expectedBody != "" && !strings.Contains(recorder.Body.String(), tt.expectedBody) {
				t.Errorf("expected body to contain %q, got %q", tt.expectedBody, recorder.Body.String())
			}
		})
	}
}
