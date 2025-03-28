package logger

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func ProvideLogger() (*zap.Logger, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}
	return logger, nil
}

var Module = fx.Module(
	"logger",
	fx.Provide(ProvideLogger),
)
