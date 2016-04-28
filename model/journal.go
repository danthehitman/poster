package model

type Journal struct {
	Uuid string `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Title string
	Description string
	Owner   User `gorm:"ForeignKey:OwnerId;AssociationForeignKey:Uuid"`
	OwnerId string `sql:"type:uuid REFERENCES users(uuid)"`
	Posts []Post `gorm:"many2many:journal_post;AssociationForeignKey:Uuid;ForeignKey:Uuid"`
	Images []Image `gorm:"many2many:journal_image;AssociationForeignKey:Uuid;ForeignKey:Uuid"`
	Links []Link `gorm:"many2many:journal_link;AssociationForeignKey:Uuid;ForeignKey:Uuid"`
}