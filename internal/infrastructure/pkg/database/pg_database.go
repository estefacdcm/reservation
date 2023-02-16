package database

import (
	"fmt"
	"os"

	"github.com/estefacdcm/reservation/internal/domain/constants"
	"github.com/sirupsen/logrus"
	gormpostgres "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	connectionString = "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable"
)

func NewPostgresConnection() (db *gorm.DB) {
	//connection := fmt.Sprintf(connectionString, os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	connection := fmt.Sprintf(connectionString, "localhost", "5432", "postgres", "postgres", "reservation")
	db, err := gorm.Open(gormpostgres.Open(connection), &gorm.Config{
		SkipDefaultTransaction: false,
		Logger:                 logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		logrus.Errorf("Error doing DB connection: %v", err.Error())
		panic(err)
	}
	logrus.Printf("Successful DB connection env %v", os.Getenv("ENV"))
	return db
}

func GetDatabaseNameConnection() (string, error) {
	return constants.DBConnectionName, nil
}
