package util

import "github.com/xormplus/xorm"

type XormSession interface {
	Session(sqlTagName string) *xorm.Session
}
