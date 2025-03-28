package cmd

import (
	"github.com/froppa/company-api/config"
	"github.com/froppa/company-api/internal/server"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

var serveCmd = &cobra.Command{
	Use:   "serve_http",
	Short: "Start HTTP server",
	Run: func(cmd *cobra.Command, args []string) {
		app := fx.New(
			config.Module, // Provide configuration
			server.Module, // Provide server
		)

		app.Run()
	},
}

func Execute() error {
	return serveCmd.Execute()
}
