package database

import (
	"fmt"

	"github.com/cmparrela/go-clean-architecture/domain/entity"
	"github.com/cmparrela/go-clean-architecture/infrastructure/adapters/env"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupDatabase(env *env.Config) *gorm.DB {
	dsn := getDataSourceName(env)
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = database.AutoMigrate(&entity.User{})
	if err != nil {
		panic("failed to auto migrate)")
	}

	return database
}

func getDataSourceName(env *env.Config) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		env.DatabaseUser,
		env.DatabasePassword,
		env.DatabaseHost,
		env.DatabasePort,
		env.DatabaseName,
	)
}
