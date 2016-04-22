package model

type ResourceAction int; const (
	ReadResourceAction ResourceAction = 1 + iota
	WriteResourceAction
)

func (r ResourceAction) String() string { return resourceActions[r-1] }

var resourceActions = [...]string{
	"Read",
	"Write",
}

type ResourceAuthorization struct {
	Uuid string `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	UserId string `sql:"type:uuid REFERENCES users(uuid)"`
	ResourceId string `sql:"type:uuid;`
	Action ResourceAction `gorm:"not null;"`
	ResourceType string
}