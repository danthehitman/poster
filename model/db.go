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

func InitiDb(connectionString string, recreate bool) {
	var err error
	Db, err = gorm.Open("postgres", connectionString)
	checkErr(err, "Failed to open db.")
	if recreate {
		err = Db.Exec("DROP SCHEMA api CASCADE").Error
		checkErr(err, "Failed to drop schema.")
	}
	err = Db.Exec("CREATE SCHEMA api AUTHORIZATION poster;").Error
	checkErr(err, "Failed to create schema api")
	err = Db.Exec(`CREATE EXTENSION "uuid-ossp";`).Error
	checkErr(err, "Failed to create extension uuid-ossp")
	err = Db.CreateTable(&User{}).Error
	checkErr(err, "Failed to create table  Users")
	err = Db.CreateTable(&Session{}).Error
	checkErr(err, "Failed to create table  Session")
	err = Db.CreateTable(&ResourceAuthorization{}).Error
	checkErr(err, "Failed to create table  ResourceAuthorization")
	err = Db.CreateTable(&ResourceGroup{}).Error
	checkErr(err, "Failed to create table  ResourceGroup")
	err = Db.CreateTable(&Link{}).Error
	checkErr(err, "Failed to create table  Journal")
	err = Db.CreateTable(&Image{}).Error
	checkErr(err, "Failed to create table  Journal")
	err = Db.CreateTable(&Post{}).Error
	checkErr(err, "Failed to create table  Post")
	err = Db.CreateTable(&Journal{}).Error
	checkErr(err, "Failed to create table  Journal")

	err = Db.Exec("CREATE OR REPLACE VIEW api.users_resources AS " +
	"SELECT ra.resource_id, " +
	"ra.user_id, ra.resource_type, ra.action " +
	"FROM api.resource_authorizations ra " +
	"UNION ALL " +
	"SELECT rg.resource_id, " +
	"ra.user_id, rg.resource_type, 1 " +
	"FROM api.resource_authorizations ra " +
	"RIGHT JOIN api.resource_groups rg ON rg.parent_resource_id = ra.resource_id; ").Error
	checkErr(err, "Failed to create view user_resources")
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Println(msg, err)
	}
}

func CreateSessionForUser(user User) (string, error) {
	var session = Session{ User:user }
	session.ExpirationDate = time.Now().Add(time.Hour)
	Db.Create(&session)
	return session.Uuid, Db.Error
}

func CreateSession(session Session) (Session, error) {
	Db.Create(&session)
	return session, Db.Error
}

func CreatePost(post Post) (Post, error) {
	Db.Create(&post)
	return post, Db.Error
}

func CreateJournal(journal Journal) (Journal, error) {
	Db.Create(&journal)
	return journal, Db.Error
}

func UpdateJournal(journal Journal) (Journal, error) {
	Db.Update(&journal)
	return journal, Db.Error
}

func CreateImage(image Image) (Image, error) {
	Db.Create(&image)
	return image, Db.Error
}

func CreateLink(link Link) (Link, error) {
	Db.Create(&link)
	return link, Db.Error
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

func CreateResourceGroup(resourceGroup ResourceGroup) error{
	Db.Create(&resourceGroup)
	return Db.Error
}