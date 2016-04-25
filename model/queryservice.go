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

func GetPostById(id string) (Post, error) {
	var post Post
	err := Db.Where("uuid = ?", id).First(&post).Error
	return post, err
}

func GetAuthenticatedUser(token string) *User {
	var session Session
	if err := Db.Where("uuid = ? and expiration_date > ?", token, time.Now()).Preload("User").Find(&session).Error; err != nil{
		return nil
	}
	return &session.User
}

func GetPostsForUser(userId string) ([]Post, error) {
	resources, err := GetAuthorizedResourcesForUser(userId, "post")
	if err != nil {
		return nil, err
	}
	var posts []Post
	Db.Where("uuid in (?)", resources).Find(&posts)
	return posts, Db.Error
}

func GetAuthorizedResourcesForUser(userId string, resourceType string) ([]string, error){
	var resources []UsersResources
	Db.Table("users_resources").Select([]string{"resource_id"}).Where("user_id = ? and resource_type = ?", userId, resourceType).Scan(&resources)
	resStrings := make([]string, len(resources))
	for i, v := range resources {
		resStrings[i] = v.ResourceId
	}
	return resStrings, Db.Error
}
