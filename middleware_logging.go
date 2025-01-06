package main

import (
	"log"
	"net/http"
	"time"
)

type ResponseWriterWrapper struct {
	http.ResponseWriter
	StatusCode int
}

// WriteHeader overrides the default WriteHeader to capture the status code
func (rw *ResponseWriterWrapper) WriteHeader(code int) {
	rw.StatusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

var sensitiveHeaders = []string{HeaderAuthorization, HeaderXSecretToken}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		headers := make(http.Header)
		for k, v := range r.Header {
			headers[k] = v
		}

		for _, h := range sensitiveHeaders {
			if _, exists := headers[h]; exists {
				headers[h] = []string{"<redacted>"}
			}
		}

		wrappedWriter := &ResponseWriterWrapper{ResponseWriter: w, StatusCode: http.StatusOK} // default status is 200

		log.Printf("Method: %s, URL: %s, Headers: %v, RemoteAddr: %s started",
			r.Method, r.URL.String(), headers, r.RemoteAddr)

		next.ServeHTTP(wrappedWriter, r)

		log.Printf("Method: %s, URL: %s, Headers: %v, RemoteAddr: %s completed in %v with status %d",
			r.Method, r.URL.String(), headers, r.RemoteAddr, time.Since(start), wrappedWriter.StatusCode)
	})
}
