package entity

import "time"

// BaseEntity
type BaseEntity struct {
	Id 			int
	GmtCreate 	time.Time
	GmtUpdate 	time.Time
	CreateUser	string
	UpdateUser 	string
}
