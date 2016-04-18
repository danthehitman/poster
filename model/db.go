package model

import (
	_ "github.com/lib/pq"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"golang.org/x/crypto/bcrypt"
	"time"
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
	Db.CreateTable(&ResourceAuthorization{})
	Db.CreateTable(&ResourceGroup{})
	Db.CreateTable(&Post{})
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Println(msg, err)
	}
}

func CreateSession(user User) (string, error) {
	var session = Session{ User:user }
	session.ExpirationDate = time.Now().Add(time.Hour)
	Db.Create(&session)
	return session.Uuid, nil
}

func CreatePost(post Post) (Post, error) {
	Db.Create(&post)
	return post, nil
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