package main

import (
	"log"

	"github.com/froppa/company-api/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatalf("Command execution failed: %v", err)
	}
}
