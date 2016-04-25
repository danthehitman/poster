package model

type UsersResources struct {
	UserId string `sql:"type:uuid"`
	ResourceId string `sql:"type:uuid"`
	ResourceType string
}
