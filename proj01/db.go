package main

import (
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"os"
)

var db gorm.DB

func dbInit() error {
	var err error

	dsn := os.Getenv("DBDSN")

	if dsn == "" {
		log.Println("Missing DBDSN env variable")
		return errors.New("No DB connection set.")
	}

	db, err = gorm.Open("mysql", dsn)

	if err != nil {
		log.Println(err.Error())
		return errors.New("Unable to initialize db driver.")
	}

	db.LogMode(true)
	db.DB().SetMaxIdleConns(100)
	db.DB().SetMaxOpenConns(10)

	return nil
}

func dbMissing() bool {
	err := db.DB().Ping()

	if err != nil {
		log.Println(err.Error())
	}

	return err != nil
}
