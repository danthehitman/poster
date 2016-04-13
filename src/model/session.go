package model

type Session struct {
	SessionId string `sql:"type:uuid;default:uuid_generate_v4();primary key"`
	User string
}
