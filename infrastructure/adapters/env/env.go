package env

import (
	"os"

	"github.com/joho/godotenv"
)

var Env map[string]string

func GetEnv(key string) string {
	return os.Getenv(key)
}

func SetupEnvFile() {
	if err := godotenv.Load(); err != nil {
		panic("error loading .env file")
	}
}
