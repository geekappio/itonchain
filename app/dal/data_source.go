package dal

import (
	. "github.com/geekappio/itonchain/app/config"
	"github.com/geekappio/itonchain/app/util"
	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
	. "github.com/geekappio/itonchain/app/enum"
)

// 数据源参数定义
var DB *xorm.Engine

/**
 * 初始化数据库连接
 */
func InitDataSource() error {
	var err error

	DB, err = xorm.NewMySQL(Config.Database.DriverName, Config.Database.DatasourceName)
	util.LogError(err)

	// 设置连接池
	DB.SetMaxIdleConns(Config.Database.MaxIdelConnections)
	DB.SetMaxOpenConns(Config.Database.MaxOpenConnections)

	// 配置模板文件
	err = DB.RegisterSqlMap(xorm.Xml(Config.XormPlus.XmlDirectory, ".xml"))
	if err != nil {
		util.LogError(err)
		return err
	}

	err = DB.RegisterSqlTemplate(xorm.Pongo2(Config.XormPlus.StplDirectory, ".stpl"))
	if err != nil {
		util.LogError(err)
		return err
	}

	// 模板文件更新监控
	err = DB.StartFSWatcher()
	if err != nil {
		util.LogError(err)
		return err
	}

	err = DB.Ping()
	if err != nil {
		util.LogError(err)
		return err
	}

	return nil
}

// FIXME
func Transaction(handler func(session *xorm.Session) ErrorCode) ErrorCode {
	// 创建新的session
	session := DB.NewSession()
	defer session.Close()

	// 捕获异常并rollback
	defer func() {
		if r := recover(); r != nil {
			err, ok := r.(error)
			if ok && nil != err {
				session.Rollback()
			}
		}
	}()

	err := session.Begin()
	if nil != err {
		return DB_TRANSACTION_ERROR
	}

	// 将session传入处理器，需要handler内部自行使用session实现数据库操作
	respCode := handler(session)

	// 基于error进行事务提交或回滚
	if respCode.IsSuccess() {
		session.Commit()
	} else {
		session.Rollback()
	}
	return respCode
}
