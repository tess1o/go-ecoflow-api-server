package middleware

import (
	"go-ecoflow-api-server/constants"
	"go-ecoflow-api-server/handlers"
	"net/http"
)

type AuthHeadersMiddleware struct {
	*handlers.BaseHandler
	headers []string
}

func NewAuthHeadersMiddleware(baseHandler *handlers.BaseHandler, headers []string) *AuthHeadersMiddleware {
	return &AuthHeadersMiddleware{
		BaseHandler: baseHandler,
		headers:     headers,
	}
}

func (a *AuthHeadersMiddleware) CheckAuthHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, h := range a.headers {
			if v, exists := r.Header[h]; !exists || (len(v) == 1 && v[0] == "") {
				a.RespondWithError(w, http.StatusUnauthorized, constants.ErrMandatoryHeaderMissing, "Mandatory header is missing or empty", map[string]string{
					"header": h,
				})
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}
