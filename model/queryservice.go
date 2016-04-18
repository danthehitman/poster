package model

import (
	"errors"
	"time"
)

func GetUserByEmail(email string) (User, error) {
	var user User
	var err error
	Db.Where("email = ?", email).First(&user)
	if user.Email == "" {
		err = errors.New("User not found.")
	}
	return user, err
}

func GetUserById(id string) (User, error) {
	var user User
	var err error
	Db.Where("uuid = ?", id).First(&user)
	if user.Email == "" {
		err = errors.New("User not found.")
	}
	return user, err
}

func GetAuthorizedUser(token string) *User {
	var session Session
	if err := Db.Where("uuid = ? and expiration_date > ?", token, time.Now()).Preload("User").Find(&session).Error; err != nil{
		return nil
	}
	return &session.User
}
