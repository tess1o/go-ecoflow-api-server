package main

import (
	"encoding/json"
	"errors"
	"github.com/tess1o/go-ecoflow"
	"log"
	"net/http"
)

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

func getEcoflowClientOrRespondWithError(r *http.Request, w http.ResponseWriter) (*ecoflow.Client, bool) {
	client, err := getEcoflowClient(r)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, ErrInvalidAuthHeader, err.Error(), nil)
		return nil, false
	}
	return client, true
}

func getEcoflowClient(r *http.Request) (*ecoflow.Client, error) {
	accessToken, secretToken, err := getTokens(r)
	if err != nil {
		return nil, err
	}
	return ecoflow.NewEcoflowClient(accessToken, secretToken), nil
}

func getTokens(r *http.Request) (string, string, error) {
	accessToken, err := getAuthorizationHeader(r)
	if err != nil {
		return "", "", err
	}
	secretToken := getSecretToken(r)
	return accessToken, secretToken, nil
}

func getAuthorizationHeader(r *http.Request) (string, error) {
	authHeader := r.Header.Get(HeaderAuthorization)
	if authHeader == "" || len(authHeader) <= 7 || authHeader[:7] != "Bearer " {
		return "", errors.New("unauthorized: invalid or missing " + HeaderAuthorization + " header")
	}
	return authHeader[7:], nil
}

// getSecretToken We have middleware to check existence of this header so no need to have additional checks
func getSecretToken(r *http.Request) string {
	return r.Header.Get(HeaderXSecretToken)
}

func writeResponse(w http.ResponseWriter, statusCode int, msg string) {
	w.WriteHeader(statusCode)
	if _, err := w.Write([]byte(msg)); err != nil {
		log.Printf("failed to write response: %v", err)
	}
}

func writeJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		writeResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeResponse(w, statusCode, string(jsonData))
}

func respondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(payload)
	if err != nil {
		log.Printf("failed to write response: %v", err)
	}
}

func respondWithSuccess(w http.ResponseWriter, data interface{}) {
	response := SuccessResponse{
		Success: true,
		Data:    data,
	}
	respondWithJSON(w, http.StatusOK, response)
}

func respondWithError(w http.ResponseWriter, statusCode int, code, message string, details interface{}) {
	response := ErrorResponse{
		Success: false,
		Error: ErrorField{
			Code:    code,
			Message: message,
			Details: details,
		},
	}
	respondWithJSON(w, statusCode, response)
}
