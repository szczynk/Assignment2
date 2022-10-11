package database

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"github.com/szczynk/Assignment2/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func StartDB() *gorm.DB {
	var (
		dbHost = viper.GetString("database.host")
		dbUser = viper.GetString("database.user")
		dbPass = viper.GetString("database.password")
		dbName = viper.GetString("database.dbname")
		dbPort = viper.GetString("database.port")
		dbTz   = viper.GetString("database.timezone")
		db     *gorm.DB
		err    error
	)

	pgConfig := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		dbHost, dbUser, dbPass, dbName, dbPort, dbTz,
	)

	db, err = gorm.Open(postgres.Open(pgConfig), &gorm.Config{FullSaveAssociations: true})
	if err != nil {
		log.Fatal("error connecting to database: ", err)
	}

	db.Debug().AutoMigrate(models.Order{}, models.Item{})

	return db
}
