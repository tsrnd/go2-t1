package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB {
	dbType := os.Getenv("DATABASE_TYPE")
	dbUsername := os.Getenv("DB_USER_NAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	fmt.Println(dbPassword)
	db, err := sql.Open(dbType, dbUsername+":"+dbPassword+"@/"+dbName)
	if err != nil {
		panic(err)
	}
	return db
}
