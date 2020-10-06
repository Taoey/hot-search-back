package service

import (
	"github.com/Taoey/hot-search-back/pkg/sysinit"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestOnZhihuItemsAdd(t *testing.T) {
	sysinit.InitConf()
	sysinit.InitMysql()

	items, _ := ZhuhuLoaclData2ZhihuItem()

	OnZhihuItemsAddDao(items)
}
