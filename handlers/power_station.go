package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/tess1o/go-ecoflow"
	"go-ecoflow-api-server/constants"
	"net/http"
)

type PowerStationHandler struct {
	*BaseHandler
}

func NewPowerStationHandler(baseHandler *BaseHandler) *PowerStationHandler {
	return &PowerStationHandler{baseHandler}
}

func (h *PowerStationHandler) RegisterRoutes(router chi.Router) {
	router.Put("/api/power_station/{serial_number}/out/ac", h.PowerStationEnableAc())
	router.Put("/api/power_station/{serial_number}/out/dc", h.PowerStationEnableDc())
	router.Put("/api/power_station/{serial_number}/out/car", h.PowerStationSetEnableCarCharging())
	router.Put("/api/power_station/{serial_number}/input/speed", h.PowerStationSetChargingSpeed())
	router.Put("/api/power_station/{serial_number}/input/car", h.PowerStationSetCarInput())
	router.Put("/api/power_station/{serial_number}/standby", h.PowerStationSetStandBy())
}

type ChangeStateRequest struct {
	State string `json:"state"`
}

// PowerStationSetEnableCarCharging enables or disables the car charger switch.
//
// @Summary Enable/Disable Car Charger
// @Description Enables or disables the car charger for the power station based on the provided state (on/off).
// @Tags Power Station
// @Accept json
// @Produce json
// @Param serial_number path string true "Serial number of the power station"
// @Param requestBody body ChangeStateRequest true "Request body containing the state"
// @Success 200 {object} SuccessResponse "Successfully toggled car charger state"
// @Failure 400 {object} ErrorResponse "Invalid input or request"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /api/power_station/{serial_number}/out/car [put]
func (h *PowerStationHandler) PowerStationSetEnableCarCharging() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		sn := r.PathValue("serial_number")

		var requestBody ChangeStateRequest
		err := json.NewDecoder(r.Body).Decode(&requestBody)
		if err != nil {
			h.RespondWithError(w, http.StatusBadRequest, constants.ErrInvalidJsonBody, "Invalid JSON Body", map[string]string{
				"serial_number": sn,
				"error":         err.Error(),
			})
			return
		}

		if requestBody.State != "on" && requestBody.State != "off" {
			h.RespondWithError(w, http.StatusBadRequest, constants.ErrInvalidParameters, "Invalid request. State must be 'on' or 'off'", map[string]string{
				"serial_number": sn,
				"state":         requestBody.State,
			})
			return
		}

		client, ok := h.GetEcoflowClientOrRespondWithError(r, w)
		if !ok {
			return
		}

		var newState ecoflow.SettingSwitcher
		if requestBody.State == "on" {
			newState = ecoflow.SettingEnabled
		} else {
			newState = ecoflow.SettingDisabled
		}

		ecoflowResponse, err := client.GetPowerStation(sn).SetCarChargerSwitch(context.Background(), newState)
		if err != nil {
			h.RespondWithError(w, http.StatusInternalServerError, constants.ErrEnableCarOut, err.Error(), map[string]string{
				"serial_number": sn,
			})
		}
		h.RespondWithSuccess(w, ecoflowResponse)
	}
}

// PowerStationEnableDc enables or disables DC
// @Summary Enable/Disable DC Output
// @Description Enables or disables the DC output switch for the power station based on the provided state (on/off).
// @Tags Power Station
// @Accept json
// @Produce json
// @Param serial_number path string true "Serial number of the power station"
// @Param requestBody body ChangeStateRequest true "Request body containing the state"
// @Success 200 {object} SuccessResponse "Successfully toggled DC output state"
// @Failure 400 {object} ErrorResponse "Invalid input or request"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /api/power_station/{serial_number}/out/dc [put]
func (h *PowerStationHandler) PowerStationEnableDc() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		sn := r.PathValue("serial_number")

		var requestBody ChangeStateRequest

		err := json.NewDecoder(r.Body).Decode(&requestBody)
		if err != nil {
			h.RespondWithError(w, http.StatusBadRequest, constants.ErrInvalidJsonBody, "Invalid JSON Body", map[string]string{
				"serial_number": sn,
				"error":         err.Error(),
			})
			return
		}

		if requestBody.State != "on" && requestBody.State != "off" {
			h.RespondWithError(w, http.StatusBadRequest, constants.ErrInvalidParameters, "Invalid request. State must be 'on' or 'off'", map[string]string{
				"serial_number": sn,
				"state":         requestBody.State,
			})
			return
		}

		client, ok := h.GetEcoflowClientOrRespondWithError(r, w)
		if !ok {
			return
		}

		var newState ecoflow.SettingSwitcher
		if requestBody.State == "on" {
			newState = ecoflow.SettingEnabled
		} else {
			newState = ecoflow.SettingDisabled
		}

		ecoflowResponse, err := client.GetPowerStation(sn).SetDcSwitch(context.Background(), newState)
		if err != nil {
			h.RespondWithError(w, http.StatusInternalServerError, constants.ErrEnableDcOut, err.Error(), map[string]string{
				"serial_number": sn,
			})
		}
		h.RespondWithSuccess(w, ecoflowResponse)
	}
}

type EnableAcRequest struct {
	AcState     string `json:"ac_state"`
	XBoostState string `json:"xboost_state"`
	OutFreq     int    `json:"out_freq"`
	OutVoltage  int    `json:"out_voltage"`
}

// PowerStationEnableAc enables or disables AC with additional settings
// @Summary Enable/Disable AC Output with settings
// @Description Enables or disables the AC output switch for the power station with additional settings, like XBoost state, output frequency, and voltage.
// @Tags Power Station
// @Accept json
// @Produce json
// @Param serial_number path string true "Serial number of the power station"
// @Param requestBody body EnableAcRequest true "Request body containing AC state, XBoost state, output frequency, and output voltage"
// @Success 200 {object} SuccessResponse "Successfully toggled AC output state with defined settings"
// @Failure 400 {object} ErrorResponse "Invalid input or request"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /api/power_station/{serial_number}/out/ac [put]
func (h *PowerStationHandler) PowerStationEnableAc() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		sn := r.PathValue("serial_number")

		var requestBody EnableAcRequest
		err := json.NewDecoder(r.Body).Decode(&requestBody)
		if err != nil {
			h.RespondWithError(w, http.StatusBadRequest, constants.ErrInvalidJsonBody, "Invalid JSON Body", map[string]string{
				"serial_number": sn,
				"error":         err.Error(),
			})
			return
		}

		if requestBody.AcState != "on" && requestBody.AcState != "off" {
			h.RespondWithError(w, http.StatusBadRequest, constants.ErrInvalidParameters, "Invalid request. ac_state must be 'on' or 'off'", map[string]string{
				"serial_number": sn,
				"ac_state":      requestBody.AcState,
			})
			return
		}

		if requestBody.XBoostState != "on" && requestBody.XBoostState != "off" {
			h.RespondWithError(w, http.StatusBadRequest, constants.ErrInvalidParameters, "Invalid request. xboost_state must be 'on' or 'off'", map[string]string{
				"serial_number": sn,
				"xboost_state":  requestBody.XBoostState,
			})
			return
		}

		if requestBody.OutFreq != 50 && requestBody.OutFreq != 60 {
			h.RespondWithError(w, http.StatusBadRequest, constants.ErrInvalidParameters, "Invalid request. out_freq must be 50 or 60", map[string]string{
				"serial_number": sn,
				"out_freq":      fmt.Sprintf("%d", requestBody.OutFreq),
			})
			return
		}

		if requestBody.OutVoltage == 0 {
			h.RespondWithError(w, http.StatusBadRequest, constants.ErrInvalidParameters, "Invalid request. out_voltage must not be 0", map[string]string{
				"serial_number": sn,
				"out_voltage":   fmt.Sprintf("%d", requestBody.OutVoltage),
			})
			return
		}

		client, ok := h.GetEcoflowClientOrRespondWithError(r, w)
		if !ok {
			return
		}

		var newAcState ecoflow.SettingSwitcher
		if requestBody.AcState == "on" {
			newAcState = ecoflow.SettingEnabled
		} else {
			newAcState = ecoflow.SettingDisabled
		}

		var newXBoostState ecoflow.SettingSwitcher
		if requestBody.XBoostState == "on" {
			newXBoostState = ecoflow.SettingEnabled
		} else {
			newXBoostState = ecoflow.SettingDisabled
		}

		var newOutFreq ecoflow.GridFrequency
		if requestBody.OutFreq == 50 {
			newOutFreq = ecoflow.GridFrequency50Hz
		} else {
			newOutFreq = ecoflow.GridFrequency60Hz
		}

		ecoflowResponse, err := client.GetPowerStation(sn).SetAcEnabled(context.Background(), newAcState, newXBoostState, newOutFreq, requestBody.OutVoltage)

		if err != nil {
			h.RespondWithError(w, http.StatusInternalServerError, constants.ErrEnableAcOut, err.Error(), map[string]string{
				"serial_number": sn,
			})
		}
		h.RespondWithSuccess(w, ecoflowResponse)
	}
}

type SetChargingSpeedRequest struct {
	Watts int `json:"watts"`
}

// PowerStationSetChargingSpeed godoc
// @Summary Set the charging speed (in watts) for a power station.
// @Description Allows setting the charging speed in watts for a specific power station identified by its serial number.
// @Tags Power Station
// @Accept json
// @Produce json
// @Param serial_number path string true "Serial Number of the Power Station"
// @Param requestBody body SetChargingSpeedRequest true "Charging Speed Request Body"
// @Success 200 {object} SuccessResponse "Successfully set the charging speed"
// @Failure 400 {object} ErrorResponse "Invalid request parameters or JSON body"
// @Failure 500 {object} ErrorResponse "Internal server error occurred while processing the request"
// @Router /api/power_station/{serial_number}/charging-speed [put]
func (h *PowerStationHandler) PowerStationSetChargingSpeed() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		sn := r.PathValue("serial_number")

		var requestBody SetChargingSpeedRequest
		err := json.NewDecoder(r.Body).Decode(&requestBody)
		if err != nil {
			h.RespondWithError(w, http.StatusBadRequest, constants.ErrInvalidJsonBody, "Invalid JSON Body", map[string]string{
				"serial_number": sn,
				"error":         err.Error(),
			})
			return
		}

		if requestBody.Watts <= 0 {
			h.RespondWithError(w, http.StatusBadRequest, constants.ErrInvalidParameters, "Invalid request. watts must be greater than 0", map[string]string{
				"serial_number": sn,
				"watts":         fmt.Sprintf("%d", requestBody.Watts),
			})
			return
		}

		client, ok := h.GetEcoflowClientOrRespondWithError(r, w)
		if !ok {
			return
		}

		ecoflowResponse, err := client.GetPowerStation(sn).SetAcChargingSettings(context.Background(), requestBody.Watts, 0)
		if err != nil {
			h.RespondWithError(w, http.StatusInternalServerError, constants.ErrPowerStationSetChargingSpeed, err.Error(), map[string]string{
				"serial_number": sn,
			})
		}
		h.RespondWithSuccess(w, ecoflowResponse)
	}
}

type InputAmpsRequest struct {
	InputAmps int `json:"amps"`
}

// PowerStationSetCarInput set the input for car charger
// @Summary Set the car input charging current for a power station.
// @Description Allows setting the car input charging current (in amps) for a specific power station identified by its serial number.
// @Tags Power Station
// @Accept json
// @Produce json
// @Param serial_number path string true "Serial Number of the Power Station"
// @Param requestBody body InputAmpsRequest true "Car Input Charging Request Body"
// @Success 200 {object} SuccessResponse "Successfully set the car input current"
// @Failure 400 {object} ErrorResponse "Invalid request parameters or JSON body"
// @Failure 500 {object} ErrorResponse "Internal server error occurred while processing the request"
// @Router /api/power_station/{serial_number}/car-input [put]
func (h *PowerStationHandler) PowerStationSetCarInput() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		sn := r.PathValue("serial_number")

		var requestBody InputAmpsRequest
		err := json.NewDecoder(r.Body).Decode(&requestBody)
		if err != nil {
			h.RespondWithError(w, http.StatusBadRequest, constants.ErrInvalidJsonBody, "Invalid JSON Body", map[string]string{
				"error":         err.Error(),
				"serial_number": sn,
			})
			return
		}
		if requestBody.InputAmps < 4 || requestBody.InputAmps > 10 {
			h.RespondWithError(w, http.StatusBadRequest, constants.ErrInvalidParameters, "Invalid request. amps must be between 4 and 10", map[string]string{
				"serial_number": sn,
				"amps":          fmt.Sprintf("%d", requestBody.InputAmps),
			})
			return
		}

		client, ok := h.GetEcoflowClientOrRespondWithError(r, w)
		if !ok {
			return
		}

		ecoflowResponse, err := client.GetPowerStation(sn).Set12VDcChargingCurrent(context.Background(), requestBody.InputAmps*1000)
		if err != nil {
			h.RespondWithError(w, http.StatusInternalServerError, constants.ErrPowerStationSetCarInput, err.Error(), map[string]string{
				"serial_number": sn,
			})
		}
		h.RespondWithSuccess(w, ecoflowResponse)
	}
}

type StandByRequest struct {
	Type    string `json:"type"`
	StandBy int    `json:"stand_by"`
}

// PowerStationSetStandBy set stand by parameters for Device, AC, Car LCD screen.
// @Summary Set standby settings for a power station.
// @Description Allows setting standby time and standby type for a specific power station identified by its serial number.
// @Tags Power Station
// @Accept json
// @Produce json
// @Param serial_number path string true "Serial Number of the Power Station"
// @Param requestBody body StandByRequest true "Standby Request Body"
// @Success 200 {object} SuccessResponse "Successfully set the standby settings"
// @Failure 400 {object} ErrorResponse "Invalid request parameters or JSON body"
// @Failure 500 {object} ErrorResponse "Internal server error occurred while processing the request"
// @Router /api/power_station/{serial_number}/standby [put]
func (h *PowerStationHandler) PowerStationSetStandBy() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		sn := r.PathValue("serial_number")

		var requestBody StandByRequest

		err := json.NewDecoder(r.Body).Decode(&requestBody)
		if err != nil {
			h.RespondWithError(w, http.StatusBadRequest, constants.ErrInvalidJsonBody, "Invalid JSON Body", map[string]string{
				"serial_number": sn,
				"error":         err.Error(),
			})
			return
		}

		if requestBody.StandBy < 0 {
			h.RespondWithError(w, http.StatusBadRequest, constants.ErrInvalidParameters, "Invalid request. stand_by must be greater than 0", map[string]string{
				"serial_number": sn,
				"stand_by":      fmt.Sprintf("%d", requestBody.StandBy),
			})
			return
		}

		if requestBody.Type != "device" && requestBody.Type != "ac" && requestBody.Type != "car" && requestBody.Type != "lcd" {
			h.RespondWithError(w, http.StatusBadRequest, constants.ErrInvalidParameters, "Invalid request. type must be 'device', 'ac', 'car' or 'lcd'", map[string]string{
				"serial_number": sn,
				"type":          requestBody.Type,
			})
			return
		}

		client, ok := h.GetEcoflowClientOrRespondWithError(r, w)
		if !ok {
			return
		}
		ecoflowResponse, err := handleStandbyType(context.Background(), client.GetPowerStation(sn), requestBody.Type, requestBody.StandBy)
		if err != nil {
			h.RespondWithError(w, http.StatusInternalServerError, constants.ErrPowerStationSetStandBy, err.Error(), map[string]string{
				"serial_number": sn,
			})
		}
		h.RespondWithSuccess(w, ecoflowResponse)
	}
}

func handleStandbyType(ctx context.Context, ps *ecoflow.PowerStation, standbyType string, standbyTime int) (*ecoflow.CmdSetResponse, error) {
	switch standbyType {
	case "device":
		return ps.SetStandByTime(ctx, standbyTime)
	case "ac":
		return ps.SetAcStandByTime(ctx, standbyTime)
	case "car":
		return ps.SetCarStandByTime(ctx, standbyTime)
	case "lcd":
		return ps.SetLcdScreenTimeout(ctx, standbyTime)
	default:
		return nil, fmt.Errorf("invalid standby type: %s", standbyType)
	}
}
