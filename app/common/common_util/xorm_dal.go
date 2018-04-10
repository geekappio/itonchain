package common_util

import "github.com/xormplus/xorm"

type XormSession interface {
	getSqlTemplateClient(sqlTagName string, args ...interface{}) *xorm.Session
	//Session() *xorm.Session
}

//type XormSqlMapper struct {
//	XormSession
//	session *xorm.Session
//}