package handlers

import (
	"encoding/json"
	"github.com/go-chi/httplog/v2"
	"github.com/tess1o/go-ecoflow"
	"go-ecoflow-api-server/constants"
	"go-ecoflow-api-server/service"
	"net/http"
)

type BaseHandler struct {
	Logger *httplog.Logger
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

func NewBaseHandler(logger *httplog.Logger) *BaseHandler {
	return &BaseHandler{
		Logger: logger,
	}
}

func (b *BaseHandler) RespondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(payload)
	if err != nil {
		b.Logger.Error("failed to write resoonse", "error", err)
	}
}

func (b *BaseHandler) RespondWithSuccess(w http.ResponseWriter, data interface{}) {
	response := SuccessResponse{
		Success: true,
		Data:    data,
	}
	b.RespondWithJSON(w, http.StatusOK, response)
}

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

func (b *BaseHandler) GetEcoflowClientOrRespondWithError(r *http.Request, w http.ResponseWriter) (*ecoflow.Client, bool) {
	client, err := service.GetEcoflowClient(r)
	if err != nil {
		b.RespondWithError(w, http.StatusUnauthorized, constants.ErrInvalidAuthHeader, err.Error(), nil)
		return nil, false
	}
	return client, true
}
