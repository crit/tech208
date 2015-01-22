package main

import (
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"os"
)

var db gorm.DB

var storer DbManager = DbManager{
	DSN:     os.Getenv("DB_DSN"),
	Logging: os.Getenv("DB_LOGGING") == "on",
}

type DbManager struct {
	DSN     string
	Logging bool
	Start   bool
	Error   error
}

func (m *DbManager) InitDb() (err error) {
	if m.DSN == "" {
		log.Println("Missing DB_DSN env variable")
		m.Error = errors.New("No DB connection set.")
		return m.Error
	}

	db, err = gorm.Open("mysql", m.DSN)

	if err != nil {
		log.Println(err.Error())
		m.Error = errors.New("Unable to initialize db driver.")
		return m.Error
	}

	db.LogMode(m.Logging)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	m.Start = true

	return nil
}

func (m *DbManager) MissingDb() bool {
	if !m.Start {
		return true
	}

	if err := db.DB().Ping(); err != nil {
		log.Println(err.Error())
		return true
	}

	return false
}
