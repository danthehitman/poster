package model

import "database/sql"

type Session struct {
	sessionId sql.NullString
}
