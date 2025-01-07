package logger

import (
	"github.com/go-chi/httplog/v2"
	"go-ecoflow-api-server/constants"
	"log/slog"
)

func GetLogger(logLevel slog.Level) *httplog.Logger {
	return httplog.NewLogger("go-ecoflow-api-server", httplog.Options{
		JSON:           true,
		LogLevel:       logLevel,
		Concise:        true,
		RequestHeaders: true,
		HideRequestHeaders: []string{
			constants.HeaderAuthorization,
			constants.HeaderXSecretToken,
		},
		MessageFieldName: "message",
	})
}
