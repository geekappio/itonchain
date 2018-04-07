package entity

import "time"

// BaseEntity
type BaseEntity struct {
	Id         int64     `xorm:'id' pk autoincr`
	GmtCreate  time.Time `xorm:'gmt_create' datetime notnull created`
	GmtUpdate  time.Time `xorm:'gmt_update' datetime notnull updated`
	CreateUser string    `xorm:'create_user'`
	UpdateUser string    `xorm:'update_user'`
}
