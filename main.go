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

	router := chi.NewRouter()
	baseHandler := handlers.NewBaseHandler(log, service.GetEcoflowClient)
	deviceHandler := handlers.NewDeviceHandler(baseHandler)
	powerStationHandler := handlers.NewPowerStationHandler(baseHandler)

	// create api routes
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
	router.Use(chimiddleware.RequestID)                         //add request id to each request
	router.Use(chimiddleware.RealIP)                            //get real ip address for headers
	router.Use(httplog.RequestLogger(log))                      //log all requests without sensitive headers
	router.Use(chimiddleware.Recoverer)                         //recover in case of panic
	router.Use(chimiddleware.Timeout(constants.RequestTimeout)) //max request duration

	authheaders := []string{constants.HeaderAuthorization, constants.HeaderXSecretToken}
	router.Use(middleware.NewAuthHeadersMiddleware(baseHandler, authheaders).CheckAuthHeaders)                                   // check mandatory auth headers
	router.Use(middleware.NewRateLimitMiddleware(baseHandler, constants.RateLimit, constants.RateLimitWindowLength).RateLimit()) // rate limit (60 requests per minute)
}
