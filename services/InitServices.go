package services

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

var (
	OrmerS   orm.Ormer
	SysDataS = new(SysDataService)
	ForwardS = new(ForwardService)
)

func init() {
	//数据库连接
	//_ "github.com/mattn/go-sqlite3"
	orm.RegisterDriver("sqlite3", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "file:data/data.db?cache=shared&loc=auto")
	//orm.RegisterDataBase("default", "sqlite3", "file::memory:?mode=memory&cache=shared&loc=auto")

	//_ "github.com/go-sql-driver/mysql"
	//dataSource := beego.AppConfig.String("mysql.url")
	//orm.RegisterDriver("mysql", orm.DRMySQL)
	//orm.RegisterDataBase("default", "mysql", dataSource)

	//开启DEBUG模式，输出SQL信息
	orm.Debug = true

	OrmerS = orm.NewOrm()
	OrmerS.Using("default")

}
