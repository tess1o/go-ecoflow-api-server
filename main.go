package main

import (
	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog/v2"
	httpSwagger "github.com/swaggo/http-swagger"
	"go-ecoflow-api-server/constants"
	_ "go-ecoflow-api-server/docs" // Import generated docs package
	"go-ecoflow-api-server/handlers"
	"go-ecoflow-api-server/logger"
	"go-ecoflow-api-server/middleware"
	"go-ecoflow-api-server/service"
	"log/slog"
	"net/http"
)

// @title Ecoflow API Server
// @version 1.0
// @description API for managing Ecoflow devices.
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey Authorization
// @type apiKey
// @name Authorization
// @in header
// @description Ecoflow Access Token. Use the format: Bearer <access_token>
//
// @securityDefinitions.apikey X-Secret-Token
// @type apiKey
// @name X-Secret-Token
// @in header
// @description Ecoflow Secret Token

// @security Authorization
// @security X-Secret-Token
func main() {
	log := logger.GetLogger(slog.LevelDebug)
	baseHandler := handlers.NewBaseHandler(log, service.GetEcoflowClient)

	router := chi.NewRouter()

	deviceHandler := handlers.NewDeviceHandler(baseHandler)
	powerStationHandler := handlers.NewPowerStationHandler(baseHandler)

	router.Group(func(apiRouter chi.Router) {
		setMiddleware(apiRouter, log, baseHandler)
		deviceHandler.RegisterRoutes(apiRouter)
		powerStationHandler.RegisterRoutes(apiRouter)
	})

	router.Get("/swagger/*", httpSwagger.WrapHandler)

	slog.Info("Starting Ecoflow API Server on :8080... Swagger is available at http://localhost:8080/swagger/index.html")

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Error("Failed to start server", "error", err)
	}
}

func setMiddleware(router chi.Router, log *httplog.Logger, baseHandler *handlers.BaseHandler) {
	router.Use(chimiddleware.RequestID)
	router.Use(chimiddleware.RealIP)
	router.Use(httplog.RequestLogger(log))
	router.Use(chimiddleware.Recoverer)
	router.Use(chimiddleware.Timeout(constants.RequestTimeout))
	router.Use(middleware.NewAuthHeadersMiddleware(baseHandler).CheckAuthHeaders)
	router.Use(middleware.NewRateLimitMiddleware(baseHandler).RateLimit())
}
