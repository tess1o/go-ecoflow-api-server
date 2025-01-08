package handlers

import (
	"encoding/json"
	"github.com/go-chi/httplog/v2"
	"github.com/tess1o/go-ecoflow"
	"go-ecoflow-api-server/constants"
	"net/http"
)

// ClientProvider is a function type that takes an HTTP request and returns an ecoflow client and an error.
type ClientProvider func(r *http.Request) (*ecoflow.Client, error)

// BaseHandler provides utility methods for HTTP response handling and client retrieval in API handlers.
type BaseHandler struct {
	Logger   *httplog.Logger
	Provider ClientProvider
}

// SuccessResponse represents a successful API response.
type SuccessResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

// ErrorResponse represents an error API response.
type ErrorResponse struct {
	Success bool       `json:"success"`
	Error   ErrorField `json:"error"`
}

// ErrorField contains details about the error.
type ErrorField struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"` // Optional
}

func NewBaseHandler(logger *httplog.Logger, provider ClientProvider) *BaseHandler {
	return &BaseHandler{
		Logger:   logger,
		Provider: provider,
	}
}

// RespondWithJSON sends a JSON response with the specified status code and payload to the HTTP response writer.
// It sets the Content-Type header to "application/json" and encodes the provided payload into the response body.
// Logs an error using the handler's logger if the response encoding fails.
func (b *BaseHandler) RespondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(payload)
	if err != nil {
		b.Logger.Error("failed to write resoonse", "error", err)
	}
}

// RespondWithSuccess sends a standardized success response with HTTP 200 status, including the provided data in JSON format.
func (b *BaseHandler) RespondWithSuccess(w http.ResponseWriter, data interface{}) {
	response := SuccessResponse{
		Success: true,
		Data:    data,
	}
	b.RespondWithJSON(w, http.StatusOK, response)
}

// RespondWithError sends a standardized error response as JSON, including the HTTP status code, error code, message, and details.
func (b *BaseHandler) RespondWithError(w http.ResponseWriter, statusCode int, code, message string, details interface{}) {
	response := ErrorResponse{
		Success: false,
		Error: ErrorField{
			Code:    code,
			Message: message,
			Details: details,
		},
	}
	b.RespondWithJSON(w, statusCode, response)
}

// GetEcoflowClientOrRespondWithError retrieves an Ecoflow client using the request context or sends an error response if unavailable.
// Returns the client and a boolean indicating success (true) or failure (false).
func (b *BaseHandler) GetEcoflowClientOrRespondWithError(r *http.Request, w http.ResponseWriter) (*ecoflow.Client, bool) {
	client, err := b.Provider(r)
	if err != nil {
		b.RespondWithError(w, http.StatusUnauthorized, constants.ErrInvalidAuthHeader, err.Error(), nil)
		return nil, false
	}
	return client, true
}
