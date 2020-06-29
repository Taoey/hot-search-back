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

	MysqlSession = db.DB()
}
