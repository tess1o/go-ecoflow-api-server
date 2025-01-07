package middleware

import (
	"go-ecoflow-api-server/constants"
	"go-ecoflow-api-server/handlers"
	"net/http"
)

var authHeaders = []string{
	constants.HeaderAuthorization,
	constants.HeaderXSecretToken,
}

type AuthHeadersMiddleware struct {
	*handlers.BaseHandler
}

func NewAuthHeadersMiddleware(baseHandler *handlers.BaseHandler) *AuthHeadersMiddleware {
	return &AuthHeadersMiddleware{baseHandler}
}

func (a *AuthHeadersMiddleware) CheckAuthHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, h := range authHeaders {
			if _, exists := r.Header[h]; !exists {
				a.RespondWithError(w, http.StatusUnauthorized, constants.ErrMandatoryHeaderMissing, "Mandatory header is missing", map[string]string{
					"header": h,
				})
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}
