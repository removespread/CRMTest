package main

import (
	"crm/internal/repository/handlers"
	"crm/internal/srvenv"
	"net/http"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func main() {
	fx.New(
		fx.WithLogger(func(logger *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: logger}
		}),
		fx.Provide(
			srvenv.NewHTTPServer,
			handlers.NewAccountHandler,
			handlers.NewContactHandler,
			handlers.NewAuthHandler,
			handlers.NewBidHandler,
			handlers.NewPartnerHandler,
			srvenv.NewLogger,
		),
		fx.Invoke(func(*http.Server) {}),
	).Run()
}
