package database

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Connect() *gorm.DB {
	dbType := os.Getenv("DATABASE_TYPE")
	dbUsername := os.Getenv("DB_USER_NAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	db, err := gorm.Open(dbType, dbUsername+":"+dbPassword+"@/"+dbName+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err.Error())
	}
	return db
}
