package sysinit

import (
	"fmt"
	. "github.com/Taoey/hot-search-back/pkg"
	. "github.com/Taoey/hot-search-back/pkg/service"
	"github.com/bamzi/jobrunner"
	"time"
)

func InitQuartz() {
	jobrunner.Start()
	jobrunner.Schedule("@every 24h", CronGetZhihuHotData{})
}

//-----打印时间任务-------
type PrintTime struct{}

func (p PrintTime) Run() {
	fmt.Println("time:", time.Now())
}

//-----定时获取知乎热榜内容----
type CronGetZhihuHotData struct{}

func (p CronGetZhihuHotData) Run() {
	LOG.Debug("定时获取知乎热榜 start ")
	OnSaveZhihuItems()
	LOG.Debug("定时获取知乎热榜 end ")
}
