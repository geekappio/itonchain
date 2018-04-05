package dao

import (
	. "github.com/geekappio/itonchain/app/config"
	"github.com/geekappio/itonchain/app/util"
	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
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
