package db

import (
	"context"
	"fmt"

	"github.com/froppa/company-api/config"
	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/cockroachdb"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func ProvideDB(lifecycle fx.Lifecycle, cfg *config.Config, logger *zap.Logger) (db.Session, error) {
	sess, err := cockroachdb.Open(cockroachdb.ConnectionURL{
		Database: cfg.Database.Name,
		Host:     fmt.Sprintf("%s:%d", cfg.Database.Host, cfg.Database.Port),
		User:     cfg.Database.User,
		Options:  map[string]string{"sslmode": cfg.Database.SSLMode},
	})
	if err != nil {
		logger.Error("DB connection failed", zap.Error(err))
		return nil, fmt.Errorf("DB connection failed: %w", err)
	}

	lifecycle.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			logger.Info("Closing database connection")
			return sess.Close()
		},
	})

	logger.Info("Database connected successfully")
	return sess, nil
}

var Module = fx.Options(
	fx.Provide(ProvideDB),
)
