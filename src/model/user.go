package model


type User struct {
	Uuid string `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Email string `gorm:"unique_index"`
	FirstName string
	LastName string
	Password string
}
