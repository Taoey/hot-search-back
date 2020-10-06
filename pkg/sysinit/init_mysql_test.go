package sysinit

import (
	"database/sql"
	"fmt"
	"github.com/Taoey/hot-search-back/pkg/service"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"testing"
)

func Test01(t *testing.T) {
	//第⼀步：打开数据库,格式是 ⽤户名:密码@(IP:端口)/数据库名称？编码⽅式
	db, err := sql.Open("mysql", "root:123456@tcp(192.168.3.148:3306)/hotsearch?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}
	//关闭数据库
	defer db.Close()

	sql := "INSERT INTO zhihu (id, name) VALUES (?, ?)"
	args := []interface{}{12, "hong"}
	result, err := db.Exec(sql, args...)
	fmt.Println(result, err)

}

type zhihu struct {
	Id   int64
	Name string
}

// 连接测试
func Test02(t *testing.T) {
	db, err := gorm.Open("mysql", "root:123456@tcp(192.168.3.148:3306)/hotsearch?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(db)
	defer db.Close()
}

// 插入数据
func Test03(t *testing.T) {
	db, err := gorm.Open("mysql", "root:123456@tcp(192.168.3.148:3306)/hotsearch?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	item := zhihu{18, "tao"}
	db.LogMode(true)
	create := db.Table("zhihu").Create(item)
	fmt.Println(create)
}

// 查询数据
func Test04(t *testing.T) {
	db, err := gorm.Open("mysql", "root:123456@tcp(192.168.3.148:3306)/hotsearch?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	items := []service.ZhihuItem{}

	db.LogMode(true)
	db.CreateTable(service.ZhihuItem{})
	//db.Table("zhihu").Find(&items)
	fmt.Println(items)
}
