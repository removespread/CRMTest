package main

import (
	"context"
	"crm/internal/middleware"
	"crm/internal/repository/postgres"
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
			postgres.NewAccountHandler,
			postgres.NewContactHandler,
			middleware.NewAuthHandler,
			postgres.NewBidHandler,
			postgres.NewPartnerHandler,
			srvenv.NewLogger,
			srvenv.NewConfig,
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
