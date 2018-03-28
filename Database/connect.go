package database

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func Connect() *gorm.DB {
	dbType := os.Getenv("DATABASE_TYPE")
	dbUsername := os.Getenv("DB_USER_NAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHostName := os.Getenv("DB_HOSTNAME")
	dbPort := os.Getenv("DB_PORT")
	db, err := gorm.Open(dbType, "host="+dbHostName+" port="+dbPort+" user="+dbUsername+" dbname="+dbName+" password="+dbPassword)
	if err != nil {
		log.Fatal(err.Error())
	}
	db.LogMode(true)
	return db
}
