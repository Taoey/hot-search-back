package sysinit

import (
	"database/sql"
	"fmt"
	. "github.com/Taoey/hot-search-back/pkg"
	"github.com/jinzhu/gorm"
)

func DbGetConnect() *sql.DB {
	db, err := gorm.Open("mysql", GCF.UString("mysql.url"))
	if err != nil {
		fmt.Println(err)
		return nil
	}
	db.DB().SetMaxIdleConns(GCF.UInt("mysql.max_idle", 10))  //空闲最大连接数
	db.DB().SetMaxOpenConns(GCF.UInt("mysql.max_open", 100)) //最大连接数
	return db.DB()
}
