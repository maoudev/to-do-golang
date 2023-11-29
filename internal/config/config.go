package config

import (
	"log/slog"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// API_PORT is the port in which the server runs.
	API_PORT string

	// DSN is the database connection string.
	DSN string

	HASH_COST int

	SECRET_KEY string
)

// GetEnv Load the environment variables and keep them in exportable variables.
func GetEnv() {
	if err := godotenv.Load(); err != nil {
		slog.Error(err.Error())
		return
	}

	API_PORT = os.Getenv("PORT")
	DSN = os.Getenv("DSN")
	HASH_COST, _ = strconv.Atoi(os.Getenv("HASH_COST"))
	SECRET_KEY = os.Getenv("SECRET_KEY")

	slog.Info("Satisfactorily loaded environment variables")
}
