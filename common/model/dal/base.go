package dal

import "time"

type Base struct {
	Id 			int
	GmtCreate 	time.Time
	GmtUpdate 	time.Time
	CreateUser	string
	UpdateUser 	string
}
