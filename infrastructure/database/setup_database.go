package database

import (
	"fmt"
	"github.com/cmparrela/go-clean-architecture/infrastructure/config"

	"github.com/cmparrela/go-clean-architecture/domain/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupDatabase(config *config.Config) *gorm.DB {
	dsn := getDataSourceName(config)
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

func getDataSourceName(config *config.Config) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		config.DatabaseUser,
		config.DatabasePassword,
		config.DatabaseHost,
		config.DatabasePort,
		config.DatabaseName,
	)
}
