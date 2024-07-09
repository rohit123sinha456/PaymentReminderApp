package main

import (
	"log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	. "github.com/rohit123sinha456/payredapp/router"
)

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:Admin@123@tcp(127.0.0.1:3306)/custengage?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil{
		log.Fatal(err)
	}

	apiserver := NewAPIServer(":8080",db)
	apiserver.Run()

	// Migrate(db)
  }