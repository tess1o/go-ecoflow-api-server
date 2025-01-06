package main

import (
	"context"
	"encoding/json"
	"net/http"
)

func GetDevicesList() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		client, ok := getEcoflowClientOrRespondWithError(r, w)
		if !ok {
			return
		}

		ecoflowResponse, err := client.GetDeviceList(context.Background())
		if err != nil {
			respondWithError(w, http.StatusBadRequest, ErrGetDevicesList, err.Error(), nil)
			return
		}
		respondWithSuccess(w, ecoflowResponse)
	}
}

func GetDeviceParametersAll() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		client, ok := getEcoflowClientOrRespondWithError(r, w)
		if !ok {
			return
		}

		sn := r.PathValue("serial_number")
		ecoflowResponse, err := client.GetDeviceAllParameters(context.Background(), sn)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, ErrGetAllDeviceParameters, err.Error(), map[string]string{
				"serial_number": sn,
			})
			return
		}
		respondWithSuccess(w, ecoflowResponse)
	}
}

func GetDeviceParametersQuery() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		sn := r.PathValue("serial_number")

		var requestBody struct {
			Parameters []string `json:"parameters"`
		}

		err := json.NewDecoder(r.Body).Decode(&requestBody)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, ErrInvalidJsonBody, "Invalid JSON Body", map[string]string{
				"serial_number": sn,
				"error":         err.Error(),
			})
			return
		}

		if len(requestBody.Parameters) == 0 {
			respondWithError(w, http.StatusBadRequest, ErrInvalidJsonBody, "No parameters provided", map[string]string{
				"serial_number": sn,
			})
			return
		}

		client, ok := getEcoflowClientOrRespondWithError(r, w)
		if !ok {
			return
		}

		ecoflowResponse, err := client.GetDeviceParameters(context.Background(), sn, requestBody.Parameters)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, ErrGetDeviceParameters, err.Error(), map[string]string{
				"serial_number": sn,
			})
			return
		}
		respondWithSuccess(w, ecoflowResponse)
	}
}
