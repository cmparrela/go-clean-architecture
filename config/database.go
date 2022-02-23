package config

import (
	"fmt"

	"github.com/cmparrela/go-clean-architecture/infrastructure/adapters/env"
)

type DatabaseConfig struct {
	Database string
	User     string
	Password string
	Host     string
	Port     string
}

func NewDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		env.GetEnv("MYSQL_DATABASE"),
		env.GetEnv("MYSQL_USER"),
		env.GetEnv("MYSQL_PASSWORD"),
		env.GetEnv("MYSQL_HOST"),
		env.GetEnv("MYSQL_PORT"),
	}
}

func (dbConfig *DatabaseConfig) GetDataSourceName() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Database,
	)
}
