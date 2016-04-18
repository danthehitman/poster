package model

type ResourceGroup struct {
	ParentResourceId string `sql:"type:uuid;`
	ResourceId string `sql:"type:uuid;`
	ResourceType string
}
