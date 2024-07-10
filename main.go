package main

import (
	"log"
	"os"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/joho/godotenv"
	. "github.com/rohit123sinha456/payredapp/router"
	// . "github.com/rohit123sinha456/payredapp/migrations"
	"time"
	"fmt"
)

func main() {
	var db *gorm.DB
	err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }
	// Get the values from the .env file
    user := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    host := os.Getenv("DB_HOST")
    port := os.Getenv("DB_PORT")
    dbname := os.Getenv("DB_NAME")
    charset := os.Getenv("DB_CHARSET")
    parseTime := os.Getenv("DB_PARSE_TIME")
    loc := os.Getenv("DB_LOC")

	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	// dsn := "root:Admin@123@tcp(127.0.0.1:3306)/custengage?charset=utf8mb4&parseTime=True&loc=Local"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s",user, password, host, port, dbname, charset, parseTime, loc)
	log.Printf(dsn)
	timeout:= 20
	for i:= 1; i<= timeout; i++ {
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		time.Sleep(time.Second * 5)
	}

	if db == nil {
		log.Fatal("Couldn't connect DB")
	}
	apiserver := NewAPIServer(":8080",db)
	apiserver.Run()

	// Migrate(db)
  }