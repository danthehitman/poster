package model

import (
	_ "github.com/lib/pq"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"golang.org/x/crypto/bcrypt"
)

var Db *gorm.DB

func InitiDb(connectionString string) {
	var err error
	Db, err = gorm.Open("postgres", connectionString)
	checkErr(err, "Failed to open db.")
	err = Db.Exec("CREATE SCHEMA api AUTHORIZATION poster;").Error
	checkErr(err, "Failed to create schema api")
	err = Db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";").Error
	checkErr(err, "Failed to create extension uuid-ossp")
	Db.CreateTable(&User{})
	Db.CreateTable(&Session{})
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Println(msg, err)
	}
}

func CreateSession(email string, pass string) (string, error) {
	var user User
	Db.Where("email = ?", email).First(&user)
	if user.Email == "" {

	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass))
	if err != nil {
		return "", err
	}
	var session = Session{ User:user }
	Db.Create(&session)
	return session.SessionId, nil
}

//func SessionForUser(sessionId string) (User, error) {
//	var session Session
//	Db.Where("session_id = ?", sessionId).First(&session)
//	Db.Model(&session).Related(&session.User)
//	return session.User, nil
//}

func CreateUser(user User) (User, error){
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return user, err
	}
	user.Password = string(hashedPassword[:])
	Db.Create(&user)
	return user, nil
}