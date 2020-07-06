package sysinit

import (
	"fmt"
	. "github.com/Taoey/hot-search-back/pkg"
	"github.com/jinzhu/gorm"
)

// mysql 初始化
func InitMysql() {
	db, err := gorm.Open("mysql", GCF.UString("mysql.url"))
	if err != nil {
		fmt.Println(err)
	}
	db.DB().SetMaxIdleConns(GCF.UInt("mysql.max_idle", 10))  //空闲最大连接数
	db.DB().SetMaxOpenConns(GCF.UInt("mysql.max_open", 100)) //最大连接数

	db.LogMode(true)
	db.SetLogger(Logger{})
	MysqlSession = db

	//fmt.Println("mysql init end",db)
}

type Logger struct {
}

//level := values[0]    // 操作级别
//filePath := values[1] // 文件路径
//time := values[2]     // 执行时间
//sql := values[3]      // sql语句
//params := values[4]   // 查询的参数
//rows := values[5]     // 查询出来的数据条数
func (logger Logger) Print(values ...interface{}) {
	LOG.Debugf("%v %v %v %v", values[2], values[3], values[4], values[1]) // 这里输出的就是日志内容
}
