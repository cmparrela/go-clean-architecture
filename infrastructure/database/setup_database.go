package database

import (
	"github.com/cmparrela/go-clean-architecture/config"
	"github.com/cmparrela/go-clean-architecture/domain/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupDatabase() *gorm.DB {
	dbConfig := config.NewDatabaseConfig()
	dsn := dbConfig.GetDataSourceName()

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	database.AutoMigrate(&entity.User{})

	return database
}
