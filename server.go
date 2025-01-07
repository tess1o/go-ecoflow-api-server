package main

import (
	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog/v2"
	"go-ecoflow-api-server/constants"
	"go-ecoflow-api-server/handlers"
	"go-ecoflow-api-server/logger"
	"go-ecoflow-api-server/middleware"
	"log/slog"
	"net/http"
)

func main() {
	log := logger.GetLogger(slog.LevelDebug)
	baseHandler := handlers.NewBaseHandler(log)

	router := chi.NewRouter()
	setMiddleware(router, log, baseHandler)

	deviceHandler := handlers.NewDeviceHandler(baseHandler)

	router.HandleFunc("GET /api/devices", deviceHandler.GetDevicesList())
	router.HandleFunc("GET /api/devices/{serial_number}/parameters", deviceHandler.GetDeviceParametersAll())
	router.HandleFunc("POST /api/devices/{serial_number}/parameters/query", deviceHandler.GetDeviceParametersQuery())

	powerStationHandler := handlers.NewPowerStationHandler(baseHandler)

	router.HandleFunc("PUT /api/power_station/{serial_number}/out/ac", powerStationHandler.PowerStationEnableAc())
	router.HandleFunc("PUT /api/power_station/{serial_number}/out/dc", powerStationHandler.PowerStationEnableDc())
	router.HandleFunc("PUT /api/power_station/{serial_number}/out/car", powerStationHandler.PowerStationSetEnableCarCharging())
	router.HandleFunc("PUT /api/power_station/{serial_number}/input/speed", powerStationHandler.PowerStationSetChargingSpeed())
	router.HandleFunc("PUT /api/power_station/{serial_number}/input/car", powerStationHandler.PowerStationSetCarInput())
	router.HandleFunc("PUT /api/power_station/{serial_number}/standby", powerStationHandler.PowerStationSetStandBy())

	slog.Info("Starting Ecoflow API Server on :8080...")

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Error("Failed to start server", "error", err)
	}
}

func setMiddleware(router *chi.Mux, log *httplog.Logger, baseHandler *handlers.BaseHandler) {
	router.Use(chimiddleware.RequestID)
	router.Use(chimiddleware.RealIP)
	router.Use(httplog.RequestLogger(log))
	router.Use(chimiddleware.Recoverer)
	router.Use(chimiddleware.Timeout(constants.RequestTimeout))
	router.Use(middleware.NewAuthHeadersMiddleware(baseHandler).CheckAuthHeaders)
	router.Use(middleware.NewRateLimitMiddleware(baseHandler).RateLimit())
}
