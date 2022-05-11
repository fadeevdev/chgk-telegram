package models

import "database/sql"

type User struct {
	ID        int
	Username  string
	FirstName string
	IsBot     bool
	CreatedAt sql.NullTime
}
