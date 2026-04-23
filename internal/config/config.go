package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Database_Url string
	Server_Port  string
}

func Load() (Config, error) {

	//godotenv.Load() reads .env and sets them into the process env
	//os.getenv -> reads those values

	if err := godotenv.Load(); err != nil {
		return Config{}, fmt.Errorf("error loading .env file")
	}

	Database_Url, err := extractEnv("DATABASE_URL")
	if err != nil {
		return Config{}, err
	}

	port, err := extractEnv("PORT")
	if err != nil {
		return Config{}, err
	}

	return Config{
		Database_Url: Database_Url,
		Server_Port:  port,
	}, nil
}

func extractEnv(key string) (string, error) {

	val := os.Getenv(key)

	if val == "" {
		return "", fmt.Errorf("missing environment variable: %s", key)
	}

	return val, nil
}
