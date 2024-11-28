package main

import (
	"context"
	"crm/internal/repository/handlers"
	"crm/internal/service"
	"crm/internal/srvenv"
	"log"
	"net/http"
	"os"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func main() {
	// fx application
	app := fx.New(
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
			service.NewAccountService,
			service.NewContactService,
			service.NewBidService,
			service.NewPartnerService,
		),
		fx.Invoke(func(*http.Server) {}),
	)

	// if failed to start application
	if err := app.Start(context.Background()); err != nil {
		log.Printf("Failed to start application: %v\n", err)
		os.Exit(1)
	}

	<-app.Done()

	// if failed to stop application gracefully
	if err := app.Stop(context.Background()); err != nil {
		log.Printf("Failed to stop application gracefully: %v\n", err)
		os.Exit(1)
	}
}
