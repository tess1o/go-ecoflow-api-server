package handlers

import (
	"context"
	"encoding/json"
	"go-ecoflow-api-server/constants"
	"net/http"
)

type DeviceHandler struct {
	*BaseHandler
}

func NewDeviceHandler(baseHandler *BaseHandler) *DeviceHandler {
	return &DeviceHandler{baseHandler}
}

func (h *DeviceHandler) GetDevicesList() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		client, ok := h.GetEcoflowClientOrRespondWithError(r, w)
		if !ok {
			return
		}

		ecoflowResponse, err := client.GetDeviceList(context.Background())
		if err != nil {
			h.RespondWithError(w, http.StatusBadRequest, constants.ErrGetDevicesList, err.Error(), nil)
			return
		}
		h.RespondWithSuccess(w, ecoflowResponse)
	}
}

func (h *DeviceHandler) GetDeviceParametersAll() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		client, ok := h.GetEcoflowClientOrRespondWithError(r, w)
		if !ok {
			return
		}

		sn := r.PathValue("serial_number")
		ecoflowResponse, err := client.GetDeviceAllParameters(context.Background(), sn)
		if err != nil {
			h.RespondWithError(w, http.StatusBadRequest, constants.ErrGetAllDeviceParameters, err.Error(), map[string]string{
				"serial_number": sn,
			})
			return
		}
		h.RespondWithSuccess(w, ecoflowResponse)
	}
}

func (h *DeviceHandler) GetDeviceParametersQuery() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		sn := r.PathValue("serial_number")

		var requestBody struct {
			Parameters []string `json:"parameters"`
		}

		err := json.NewDecoder(r.Body).Decode(&requestBody)
		if err != nil {
			h.RespondWithError(w, http.StatusBadRequest, constants.ErrInvalidJsonBody, "Invalid JSON Body", map[string]string{
				"serial_number": sn,
				"error":         err.Error(),
			})
			return
		}

		if len(requestBody.Parameters) == 0 {
			h.RespondWithError(w, http.StatusBadRequest, constants.ErrInvalidJsonBody, "No parameters provided", map[string]string{
				"serial_number": sn,
			})
			return
		}

		client, ok := h.GetEcoflowClientOrRespondWithError(r, w)
		if !ok {
			return
		}

		ecoflowResponse, err := client.GetDeviceParameters(context.Background(), sn, requestBody.Parameters)
		if err != nil {
			h.RespondWithError(w, http.StatusBadRequest, constants.ErrGetDeviceParameters, err.Error(), map[string]string{
				"serial_number": sn,
			})
			return
		}
		h.RespondWithSuccess(w, ecoflowResponse)
	}
}
