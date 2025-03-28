package repositories

import (
	"github.com/froppa/company-api/internal/db"
	"go.uber.org/fx"
)

var Module = fx.Options(
	db.Module,                 // Ensure the database session is provided
	fx.Provide(NewRepository), // Provide the repository
)
