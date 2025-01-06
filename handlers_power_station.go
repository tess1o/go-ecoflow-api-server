package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/tess1o/go-ecoflow"
	"net/http"
)

func PowerStationSetEnableCarCharging() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		sn := r.PathValue("serial_number")

		var requestBody struct {
			State string `json:"state"`
		}

		err := json.NewDecoder(r.Body).Decode(&requestBody)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, ErrInvalidJsonBody, "Invalid JSON Body", map[string]string{
				"serial_number": sn,
				"error":         err.Error(),
			})
			return
		}

		if requestBody.State != "on" && requestBody.State != "off" {
			respondWithError(w, http.StatusBadRequest, ErrInvalidParameters, "Invalid request. State must be 'on' or 'off'", map[string]string{
				"serial_number": sn,
				"state":         requestBody.State,
			})
			return
		}

		client, ok := getEcoflowClientOrRespondWithError(r, w)
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
			respondWithError(w, http.StatusInternalServerError, ErrEnableCarOut, err.Error(), map[string]string{
				"serial_number": sn,
			})
		}
		respondWithSuccess(w, ecoflowResponse)
	}
}

func PowerStationEnableDc() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		sn := r.PathValue("serial_number")

		var requestBody struct {
			State string `json:"state"`
		}

		err := json.NewDecoder(r.Body).Decode(&requestBody)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, ErrInvalidJsonBody, "Invalid JSON Body", map[string]string{
				"serial_number": sn,
				"error":         err.Error(),
			})
			return
		}

		if requestBody.State != "on" && requestBody.State != "off" {
			respondWithError(w, http.StatusBadRequest, ErrInvalidParameters, "Invalid request. State must be 'on' or 'off'", map[string]string{
				"serial_number": sn,
				"state":         requestBody.State,
			})
			return
		}

		client, ok := getEcoflowClientOrRespondWithError(r, w)
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
			respondWithError(w, http.StatusInternalServerError, ErrEnableDcOut, err.Error(), map[string]string{
				"serial_number": sn,
			})
		}
		respondWithSuccess(w, ecoflowResponse)
	}
}

func PowerStationEnableAc() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		sn := r.PathValue("serial_number")

		var requestBody struct {
			AcState     string `json:"ac_state"`
			XBoostState string `json:"xboost_state"`
			OutFreq     int    `json:"out_freq"`
			OutVoltage  int    `json:"out_voltage"`
		}

		err := json.NewDecoder(r.Body).Decode(&requestBody)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, ErrInvalidJsonBody, "Invalid JSON Body", map[string]string{
				"serial_number": sn,
				"error":         err.Error(),
			})
			return
		}

		if requestBody.AcState != "on" && requestBody.AcState != "off" {
			respondWithError(w, http.StatusBadRequest, ErrInvalidParameters, "Invalid request. ac_state must be 'on' or 'off'", map[string]string{
				"serial_number": sn,
				"ac_state":      requestBody.AcState,
			})
			return
		}

		if requestBody.XBoostState != "on" && requestBody.XBoostState != "off" {
			respondWithError(w, http.StatusBadRequest, ErrInvalidParameters, "Invalid request. xboost_state must be 'on' or 'off'", map[string]string{
				"serial_number": sn,
				"xboost_state":  requestBody.XBoostState,
			})
			return
		}

		if requestBody.OutFreq != 50 && requestBody.OutFreq != 60 {
			respondWithError(w, http.StatusBadRequest, ErrInvalidParameters, "Invalid request. out_freq must be 50 or 60", map[string]string{
				"serial_number": sn,
				"out_freq":      fmt.Sprintf("%d", requestBody.OutFreq),
			})
			return
		}

		if requestBody.OutVoltage == 0 {
			respondWithError(w, http.StatusBadRequest, ErrInvalidParameters, "Invalid request. out_voltage must not be 0", map[string]string{
				"serial_number": sn,
				"out_voltage":   fmt.Sprintf("%d", requestBody.OutVoltage),
			})
			return
		}

		client, ok := getEcoflowClientOrRespondWithError(r, w)
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
			respondWithError(w, http.StatusInternalServerError, ErrEnableAcOut, err.Error(), map[string]string{
				"serial_number": sn,
			})
		}
		respondWithSuccess(w, ecoflowResponse)
	}
}

func PowerStationSetChargingSpeed() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		sn := r.PathValue("serial_number")

		var requestBody struct {
			Watts int `json:"watts"`
		}

		err := json.NewDecoder(r.Body).Decode(&requestBody)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, ErrInvalidJsonBody, "Invalid JSON Body", map[string]string{
				"serial_number": sn,
				"error":         err.Error(),
			})
			return
		}

		if requestBody.Watts <= 0 {
			respondWithError(w, http.StatusBadRequest, ErrInvalidParameters, "Invalid request. watts must be greater than 0", map[string]string{
				"serial_number": sn,
				"watts":         fmt.Sprintf("%d", requestBody.Watts),
			})
			return
		}

		client, ok := getEcoflowClientOrRespondWithError(r, w)
		if !ok {
			return
		}

		ecoflowResponse, err := client.GetPowerStation(sn).SetAcChargingSettings(context.Background(), requestBody.Watts, 0)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, ErrPowerStationSetChargingSpeed, err.Error(), map[string]string{
				"serial_number": sn,
			})
		}
		respondWithSuccess(w, ecoflowResponse)
	}
}

func PowerStationSetCarInput() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		sn := r.PathValue("serial_number")
		var requestBody struct {
			InputAmps int `json:"amps"`
		}

		err := json.NewDecoder(r.Body).Decode(&requestBody)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, ErrInvalidJsonBody, "Invalid JSON Body", map[string]string{
				"error":         err.Error(),
				"serial_number": sn,
			})
			return
		}
		if requestBody.InputAmps < 4 || requestBody.InputAmps > 10 {
			respondWithError(w, http.StatusBadRequest, ErrInvalidParameters, "Invalid request. amps must be between 4 and 10", map[string]string{
				"serial_number": sn,
				"amps":          fmt.Sprintf("%d", requestBody.InputAmps),
			})
			return
		}

		client, ok := getEcoflowClientOrRespondWithError(r, w)
		if !ok {
			return
		}

		ecoflowResponse, err := client.GetPowerStation(sn).Set12VDcChargingCurrent(context.Background(), requestBody.InputAmps*1000)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, ErrPowerStationSetCarInput, err.Error(), map[string]string{
				"serial_number": sn,
			})
		}
		respondWithSuccess(w, ecoflowResponse)
	}
}

func PowerStationSetStandBy() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		sn := r.PathValue("serial_number")
		var requestBody struct {
			Type    string `json:"type"`
			StandBy int    `json:"stand_by"`
		}

		err := json.NewDecoder(r.Body).Decode(&requestBody)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, ErrInvalidJsonBody, "Invalid JSON Body", map[string]string{
				"serial_number": sn,
				"error":         err.Error(),
			})
			return
		}

		if requestBody.StandBy < 0 {
			respondWithError(w, http.StatusBadRequest, ErrInvalidParameters, "Invalid request. stand_by must be greater than 0", map[string]string{
				"serial_number": sn,
				"stand_by":      fmt.Sprintf("%d", requestBody.StandBy),
			})
			return
		}

		if requestBody.Type != "device" && requestBody.Type != "ac" && requestBody.Type != "car" && requestBody.Type != "lcd" {
			respondWithError(w, http.StatusBadRequest, ErrInvalidParameters, "Invalid request. type must be 'device', 'ac', 'car' or 'lcd'", map[string]string{
				"serial_number": sn,
				"type":          requestBody.Type,
			})
			return
		}

		client, ok := getEcoflowClientOrRespondWithError(r, w)
		if !ok {
			return
		}
		ecoflowResponse, err := handleStandbyType(context.Background(), client.GetPowerStation(sn), requestBody.Type, requestBody.StandBy)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, ErrPowerStationSetStandBy, err.Error(), map[string]string{
				"serial_number": sn,
			})
		}
		respondWithSuccess(w, ecoflowResponse)
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
