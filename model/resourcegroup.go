package model

type ResourceGroup struct {
	Uuid string `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	ParentResourceId string `sql:"type:uuid;`
	ResourceId string `sql:"type:uuid;`
	ResourceType string
}
