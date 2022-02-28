package database

import (
	"fmt"

	"github.com/cmparrela/go-clean-architecture/domain/entity"
	"github.com/cmparrela/go-clean-architecture/infrastructure/adapters/env"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupDatabase() *gorm.DB {
	dsn := getDataSourceName()
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	database.AutoMigrate(&entity.User{})

	return database
}

func getDataSourceName() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		env.GetEnv("MYSQL_USER"),
		env.GetEnv("MYSQL_PASSWORD"),
		env.GetEnv("MYSQL_HOST"),
		env.GetEnv("MYSQL_PORT"),
		env.GetEnv("MYSQL_DATABASE"),
	)
}
