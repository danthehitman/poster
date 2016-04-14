package model

type Session struct {
	SessionId string `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	User User `gorm:"ForeignKey:UserId;AssociationForeignKey:Uuid"`
	UserId string `sql:"type:uuid REFERENCES users(uuid)"`
}
