package model

type Link struct {
	Uuid string `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Title string
	Description string
	Url string
	Owner   User `gorm:"ForeignKey:OwnerId;AssociationForeignKey:Uuid"`
	OwnerId string `sql:"type:uuid REFERENCES users(uuid)"`
}