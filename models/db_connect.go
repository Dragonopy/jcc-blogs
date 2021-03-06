// models/setup.go
package models

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

// SetupDB : initializing mysql database
func ConnectDB() {
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASS")
	HOST := os.Getenv("DB_HOST")
	PORT := os.Getenv("DB_PORT")
	DBNAME := os.Getenv("DB_NAME")

	if USER == "" {
		// use default
		USER = "jcc_blogs"
		PASS = "ZAqJ_aFa6EZi26-r"
		HOST = "localhost"
		PORT = "3306"
		DBNAME = "jcc_blogs"
	}

	// GORM using DSN Format
	DSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)
	db, err := gorm.Open("mysql", DSN)
	if err != nil {
		panic(err.Error())
	}
	DB = db
	log.Println(DB)
}
