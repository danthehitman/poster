package model

import "time"

type Session struct {
	Uuid   string `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	ExpirationDate time.Time `gorm:"not null;"`
	User   User `gorm:"ForeignKey:UserId;AssociationForeignKey:Uuid"`
	UserId string `sql:"type:uuid REFERENCES users(uuid)"`
}
