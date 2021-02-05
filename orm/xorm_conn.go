package orm

import (
	"github.com/go-xorm/xorm"
	"xorm.io/core"
)

/**
1、依赖
go get github.com/go-xorm/cmd/xorm
go get github.com/go-sql-driver/mysql
go get github.com/go-xorm/xorm

2、表结构反转为pojo
cd到模板所在目录%GO_PATH%/src/github.com/go-xorm/cmd/xorm
例：xorm reverse mysql "dba:987654321@tcp(172.16.18.6:3306)/test?charset=utf8" templates/goxorm /tmp

3、官方文档
https://gobook.io/read/gitea.com/xorm/manual-zh-CN/
*/

// 创建数据库连接
func NewXormMysqlEngine(conn string, maxIdle, maxConns int) *xorm.Engine {
	var engine *xorm.Engine
	var err error

	engine, err = xorm.NewEngine(core.MYSQL, conn)
	if err != nil {
		panic(err)
		return nil
	}

	// 连接测试
	if err := engine.Ping(); err != nil {
		panic(err)
		return nil
	}

	// 闲置连接数
	engine.SetMaxIdleConns(maxIdle)
	// 最大连接数
	engine.SetMaxOpenConns(maxConns)
	// 驼峰命名。也支持自定义模式，需实现core.IMapper接口
	engine.SetTableMapper(core.SnakeMapper{})

	// 设置日志级别
	engine.Logger().SetLevel(core.LOG_DEBUG)
	// 是否输出SQL日志
	engine.ShowSQL(true)

	return engine
}

// 创建数据库连接：主从切换。主库连接异常时，自动使用从库
func NewXormMysqlMasterSlave(master, slaves *xorm.Engine) *xorm.EngineGroup {
	eg, err := xorm.NewEngineGroup(master, slaves)
	if err != nil {
		panic(err)
		return nil
	}
	return eg
}
