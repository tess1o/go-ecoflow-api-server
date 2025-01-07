package middleware

import (
	"github.com/go-chi/httprate"
	"go-ecoflow-api-server/constants"
	"go-ecoflow-api-server/handlers"
	"net/http"
	"time"
)

const (
	limit        = 60
	windowLength = time.Minute
)

type RateLimitMiddleware struct {
	*handlers.BaseHandler
}

func NewRateLimitMiddleware(b *handlers.BaseHandler) *RateLimitMiddleware {
	return &RateLimitMiddleware{b}
}

func (rl *RateLimitMiddleware) RateLimit() func(next http.Handler) http.Handler {
	rateLimiter := httprate.Limit(limit, windowLength,
		httprate.WithLimitHandler(func(w http.ResponseWriter, r *http.Request) {
			rl.RespondWithError(w, http.StatusTooManyRequests, constants.ErrRateLimitExceeded, "Rate limit exceeded", map[string]string{
				"url":         r.URL.String(),
				"method":      r.Method,
				"retry_after": w.Header().Get("Retry-After"),
			})
		}),
	)
	return rateLimiter
}
