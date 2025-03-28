package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/froppa/company-api/config"
	"github.com/gorilla/mux"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewServer(
	lc fx.Lifecycle,
	cfg *config.Config,
	logger *zap.Logger,
	handler *Handler,
) {
	router := mux.NewRouter()
	handler.RegisterRoutes(router)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Server.Port),
		Handler: router,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Info("Starting HTTP server", zap.String("addr", server.Addr))
			go func() {
				if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					logger.Fatal("HTTP server error", zap.Error(err))
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("Shutting down HTTP server")
			return server.Shutdown(ctx)
		},
	})
}
