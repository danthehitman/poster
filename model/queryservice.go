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

func GetJournalById(id string) (Journal, error) {
	var entity Journal
	err := Db.Where("uuid = ?", id).First(&entity).Error
	return entity, err
}

func GetJournalsForUser(userId string) ([]Journal, error) {
	resources, err := GetAuthorizedResourcesForUser(userId, "journal")
	if err != nil {
		return nil, err
	}
	var entities []Journal
	Db.Where("uuid in (?)", resources).Find(&entities)
	return entities, Db.Error
}

func GetPostById(id string) (Post, error) {
	var post Post
	err := Db.Where("uuid = ?", id).First(&post).Error
	return post, err
}

func GetPostsByJournalId(journalId string) ([]Post, error) {
	var journal Journal
	err := Db.Where("uuid = ?", journalId).Preload("Posts").First(&journal).Error
	return journal.Posts, err
}

func IsPostInJournal(journalId string, postId string) (bool, error) {
	var count int = 0
	err := Db.Table(JournalPostJoinTable).Where("journal_uuid = ? and post_uuid = ?", journalId, postId).Count(&count).Error
	return count > 0, err
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
