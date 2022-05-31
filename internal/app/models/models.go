package models

import "database/sql"

type User struct {
	ID        uint64
	Username  string
	FirstName string
	IsBot     bool
	CreatedAt sql.NullTime
}
