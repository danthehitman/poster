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
	err = Db.Exec("drop extension \"uuid-ossp\";").Error
	checkErr(err, "Failed to create extension uuid-ossp")
	err = Db.Exec("CREATE EXTENSION \"uuid-ossp\";").Error
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
	return session.Uuid, Db.Error
}

func CreatePost(post Post) (Post, error) {
	Db.Create(&post)
	return post, Db.Error
}

func CreateUser(user User) (User, error){
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return user, err
	}
	user.Password = string(hashedPassword[:])
	Db.Create(&user)
	return user, Db.Error
}

func DeleteUserById(userId string) error {
	Db.Where("uuid = ?", userId).Delete(&User{})
	return Db.Error
}

func CreateResourceAuthorization(resourceAuth ResourceAuthorization) error{
	Db.Create(&resourceAuth)
	return Db.Error
}