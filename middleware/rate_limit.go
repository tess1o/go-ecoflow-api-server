package middleware

import (
	"github.com/go-chi/httprate"
	"go-ecoflow-api-server/constants"
	"go-ecoflow-api-server/handlers"
	"net/http"
	"time"
)

type RateLimitMiddleware struct {
	*handlers.BaseHandler
	limit        int
	windowLength time.Duration
}

func NewRateLimitMiddleware(b *handlers.BaseHandler, limit int, windowLength time.Duration) *RateLimitMiddleware {
	return &RateLimitMiddleware{BaseHandler: b, limit: limit, windowLength: windowLength}
}

func (rl *RateLimitMiddleware) RateLimit() func(next http.Handler) http.Handler {
	rateLimiter := httprate.Limit(rl.limit, rl.windowLength,
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
