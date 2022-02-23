package database

import (
	"app/config"
	"app/entities"

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

	database.AutoMigrate(&entities.User{})

	return database
}
