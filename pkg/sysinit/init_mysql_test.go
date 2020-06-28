package sysinit

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func Test01(t *testing.T) {
	//第⼀步：打开数据库,格式是 ⽤户名：密码@/数据库名称？编码⽅式
	db, err := sql.Open("mysql", "root:123456@/hotsearch?charset=utf8")
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
