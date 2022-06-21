package main

import (
	"os"

	"github.com/gopad/gopad-api/pkg/command"
	"github.com/joho/godotenv"
)

func main() {
	if env := os.Getenv("GOPAD_API_ENV_FILE"); env != "" {
		godotenv.Load(env)
	}

	if err := command.Run(); err != nil {
		os.Exit(1)
	}
}
