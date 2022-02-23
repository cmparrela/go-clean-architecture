package env

import (
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(key string) string {
	return os.Getenv(key)
}

func SetupEnvFile() {
	if err := godotenv.Load(); err != nil {
		panic("error loading .env file")
	}
}
