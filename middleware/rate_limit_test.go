package middleware

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"go-ecoflow-api-server/handlers"
)

func TestRateLimit(t *testing.T) {
	baseHandler := &handlers.BaseHandler{}
	limit := 2
	windowLength := time.Second
	rateLimitMiddleware := NewRateLimitMiddleware(baseHandler, limit, windowLength)
	rateLimit := rateLimitMiddleware.RateLimit()

	tests := []struct {
		name           string
		requests       int
		interval       time.Duration
		expectedStatus int
	}{
		{
			name:           "below_limit",
			requests:       2,
			interval:       0,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "exceed_limit",
			requests:       3,
			interval:       0,
			expectedStatus: http.StatusTooManyRequests,
		},
		{
			name:           "reset_rate_limit",
			requests:       2,
			interval:       windowLength,
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := chi.NewRouter()
			router.Use(rateLimit)
			router.Get("/", func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			})

			for i := 0; i < tt.requests; i++ {
				req := httptest.NewRequest(http.MethodGet, "/", nil)
				rec := httptest.NewRecorder()

				router.ServeHTTP(rec, req)
				status := rec.Code

				if i == tt.requests-1 {
					if status != tt.expectedStatus {
						t.Errorf("expected status %v, got %v", tt.expectedStatus, status)
					}
				}

				time.Sleep(tt.interval)
			}
		})
	}
}
