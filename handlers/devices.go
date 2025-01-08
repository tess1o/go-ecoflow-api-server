package handlers

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"go-ecoflow-api-server/constants"
	"net/http"
)

type DeviceHandler struct {
	*BaseHandler
}

func NewDeviceHandler(baseHandler *BaseHandler) *DeviceHandler {
	return &DeviceHandler{baseHandler}
}

func (h *DeviceHandler) RegisterRoutes(router chi.Router) {
	router.Get("/api/devices", h.GetDevicesList())
	router.Get("/api/devices/{serial_number}/parameters", h.GetDeviceParametersAll())
	router.Post("/api/devices/{serial_number}/parameters/query", h.GetDeviceParametersQuery())
}

// GetDevicesList handles retrieving a list of devices
// @Summary Get a list of devices
// @Description Returns a list of all devices associated with the user
// @Tags Devices
// @Produce json
// @Success 200 {object} SuccessResponse "List of devices retrieved successfully"
// @Failure 500 {object} ErrorResponse "Error retrieving device list"
// @Router /api/devices [get]
func (h *DeviceHandler) GetDevicesList() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		client, ok := h.GetEcoflowClientOrRespondWithError(r, w)
		if !ok {
			return
		}

		ecoflowResponse, err := client.GetDeviceList(context.Background())
		if err != nil {
			h.RespondWithError(w, http.StatusInternalServerError, constants.ErrGetDevicesList, err.Error(), nil)
			return
		}
		h.RespondWithSuccess(w, ecoflowResponse)
	}
}

// GetDeviceParametersAll handles retrieving all parameters for a specific device
// @Summary Get all parameters for a device
// @Description Retrieves all available parameters for a device using its serial number
// @Tags Devices
// @Produce json
// @Param serial_number path string true "Device Serial Number"
// @Success 200 {object} SuccessResponse "Parameters retrieved successfully"
// @Failure 500 {object} ErrorResponse "Error retrieving device parameters"
// @Router /api/devices/{serial_number}/parameters [get]
func (h *DeviceHandler) GetDeviceParametersAll() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		client, ok := h.GetEcoflowClientOrRespondWithError(r, w)
		if !ok {
			return
		}

		sn := r.PathValue("serial_number")
		ecoflowResponse, err := client.GetDeviceAllParameters(context.Background(), sn)
		if err != nil {
			h.RespondWithError(w, http.StatusInternalServerError, constants.ErrGetAllDeviceParameters, err.Error(), map[string]string{
				"serial_number": sn,
			})
			return
		}
		h.RespondWithSuccess(w, ecoflowResponse)
	}
}

// QueryParametersRequest represents the request body for querying specific parameters of a device
type QueryParametersRequest struct {
	Parameters []string `json:"parameters"`
}

// GetDeviceParametersQuery handles querying specific parameters for a device
// @Summary Query specific parameters for a device
// @Description Queries specific parameters for a device using its serial number
// @Tags Devices
// @Accept json
// @Produce json
// @Param serial_number path string true "Device Serial Number"
// @Param parameters body QueryParametersRequest true "List of parameters to query"
// @Success 200 {object} SuccessResponse "Requested parameters retrieved successfully"
// @Failure 400 {object} ErrorResponse "Error Invalid JSON Body"
// @Failure 500 {object} ErrorResponse "Error retrieving device parameters"
// @Router /api/devices/{serial_number}/parameters/query [post]
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
			h.RespondWithError(w, http.StatusInternalServerError, constants.ErrGetDeviceParameters, err.Error(), map[string]string{
				"serial_number": sn,
			})
			return
		}
		h.RespondWithSuccess(w, ecoflowResponse)
	}
}
