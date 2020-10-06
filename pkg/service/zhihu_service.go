package service

import (
	. "github.com/Taoey/hot-search-back/pkg"
	"github.com/Taoey/hot-search-back/pkg/utils"
	"time"
)

// 自增ID，创建时间，更新时间，状态，标题，地址，标题图片地址，热度，个人评价
// Zhihu ...
type ZhihuItem struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`

	// Status 状态：0-删除；1-正常
	Status uint32 `json:"status"`
	// Title 标题
	Title string `json:"title"`
	// URL 热榜问题url
	URL string `json:"url"`
	// Imgurl 热榜图片
	Imgurl string `json:"imgurl"`
	// Hotscore 热度
	Hotscore int32 `json:"hotscore"`
	// Remark 个人备注
	Remark string `json:"remark"`
}

//-----------------------------
// 数据库交互dao
//-----------------------------

// 插入新数据
func OnZhihuItemsAddDao(items []ZhihuItem) {
	db := MysqlSession
	//db.Table("zhihu").CreateTable(ZhihuItem{})
	table := db.Table("zhihu")

	tx := db.Begin()
	for _, item := range items {

		// 根据问题url 来判断item在数据库中是否存在
		selector := map[string]interface{}{"url": item.URL}

		itemCount := 0
		table.Where(selector).Count(&itemCount)
		LOG.Debug(itemCount)

		if itemCount == 0 {
			err := table.Create(&item).Error
			if err != nil {
				LOG.Error(err)
			}
			tx.Rollback()
		}
	}
	tx.Commit()
}

// 更新数据

// 查询数据
// 查询条件：时间
func OnZhiHuItemsQuery() (interface{}, error) {
	db := MysqlSession

	items := []ZhihuItem{}

	zhihuTable := db.Table("zhihu")
	//localTime := utils.DateTimeStrToTime("2020-01-02 15:04:05")
	todayMidnight := utils.GetMidNightObj(time.Now())
	tomorrowMidnight := todayMidnight.Add(time.Hour * 24)

	zhihuTable.Where("created_at > ? AND created_at < ?", todayMidnight, tomorrowMidnight).Find(&items)
	return items, nil
}
