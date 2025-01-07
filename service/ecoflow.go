package service

import (
	"errors"
	"github.com/tess1o/go-ecoflow"
	"go-ecoflow-api-server/constants"
	"net/http"
)

func GetEcoflowClient(r *http.Request) (*ecoflow.Client, error) {
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
	authHeader := r.Header.Get(constants.HeaderAuthorization)
	if authHeader == "" || len(authHeader) <= 7 || authHeader[:7] != "Bearer " {
		return "", errors.New("unauthorized: invalid or missing " + constants.HeaderAuthorization + " header")
	}
	return authHeader[7:], nil
}

// getSecretToken We have middleware to check existence of this header so no need to have additional checks
func getSecretToken(r *http.Request) string {
	return r.Header.Get(constants.HeaderXSecretToken)
}
