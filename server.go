package main

import (
	"log"
	"net/http"
)

const (
	HeaderAuthorization = "Authorization"
	HeaderXSecretToken  = "X-Secret-Token"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("GET /api/devices", GetDevicesList())
	router.HandleFunc("GET /api/devices/{serial_number}/parameters", GetDeviceParametersAll())
	router.HandleFunc("POST /api/devices/{serial_number}/parameters/query", GetDeviceParametersQuery())
	router.HandleFunc("PUT /api/power_station/{serial_number}/out/ac", PowerStationEnableAc())
	router.HandleFunc("PUT /api/power_station/{serial_number}/out/dc", PowerStationEnableDc())
	router.HandleFunc("PUT /api/power_station/{serial_number}/out/car", PowerStationSetEnableCarCharging())
	router.HandleFunc("PUT /api/power_station/{serial_number}/input/speed", PowerStationSetChargingSpeed())
	router.HandleFunc("PUT /api/power_station/{serial_number}/input/car", PowerStationSetCarInput())
	router.HandleFunc("PUT /api/power_station/{serial_number}/standby", PowerStationSetStandBy())

	log.Println("Starting Ecoflow API Server on :8080...")
	err := http.ListenAndServe(":8080", LoggingMiddleware(AuthCheckMiddleware(router)))
	log.Fatal(err)
}
