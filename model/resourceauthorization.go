package model

type ResourceAction int; const (
	Read ResourceAction = 1 + iota
	Write
)

func (r ResourceAction) String() string { return resourceActions[r-1] }

var resourceActions = [...]string{
	"Read",
	"Write",
}

type ResourceAuthorization struct {
	UserId string `sql:"type:uuid REFERENCES users(uuid)"`
	ResourceId string `sql:"type:uuid;`
	Action ResourceAction `gorm:"not null;"`
	ResourceType string
}