package main

import (
	"net/http"
)

func AuthCheckMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, h := range sensitiveHeaders {
			if _, exists := r.Header[h]; !exists {
				respondWithError(w, http.StatusUnauthorized, ErrMandatoryHeaderMissing, "Mandatory header is missing", map[string]string{
					"header": h,
				})
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}
