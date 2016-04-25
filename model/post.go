package model

type Post struct {
	Uuid string `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Title string
	Description string
	Body string
	Owner   User `gorm:"ForeignKey:OwnerId;AssociationForeignKey:Uuid"`
	OwnerId string `sql:"type:uuid REFERENCES users(uuid)"`
}