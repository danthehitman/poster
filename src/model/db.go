package model

import (
	_ "github.com/lib/pq"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

var db *gorm.DB

func InitiDb(connectionString string) {
	db, err := gorm.Open("postgres", connectionString)
	checkErr(err, "Failed to open db.")
	db.CreateTable(&Session{})
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}