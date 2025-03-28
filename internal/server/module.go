package server

import (
	"github.com/froppa/company-api/internal/repositories"
	"github.com/froppa/company-api/pkg/logger"
	"go.uber.org/fx"
)

var Module = fx.Options(
	repositories.Module,
	logger.Module,
	fx.Provide(NewHandler),
	fx.Invoke(NewServer),
)
