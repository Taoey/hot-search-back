package service

import (
	"encoding/json"
	"errors"
	"fmt"
	. "github.com/Taoey/hot-search-back/pkg"
	. "github.com/Taoey/hot-search-back/pkg/utils"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

const (
	LOCAL_DATA_PATH = `./pkg/service/data.json`
)

type HotList struct {
	HotList []HotItem `json:"hotList"`
}
type FeedSpecific struct {
	AnswerCount int `json:"answerCount"`
}
type TitleArea struct {
	Text string `json:"text"`
}
type ExcerptArea struct {
	Text string `json:"text"`
}
type ImageArea struct {
	URL string `json:"url"`
}
type MetricsArea struct {
	Text string `json:"text"`
}
type LabelArea struct {
	Type        string `json:"type"`
	Text        string `json:"text"`
	NightColor  string `json:"nightColor"`
	NormalColor string `json:"normalColor"`
}
type Link struct {
	URL string `json:"url"`
}
type Target struct {
	TitleArea   TitleArea   `json:"titleArea"`
	ExcerptArea ExcerptArea `json:"excerptArea"`
	ImageArea   ImageArea   `json:"imageArea"`
	MetricsArea MetricsArea `json:"metricsArea"`
	LabelArea   LabelArea   `json:"labelArea"`
	Link        Link        `json:"link"`
}
type HotItem struct {
	Type         string       `json:"type"`
	StyleType    string       `json:"styleType"`
	ID           string       `json:"id"`
	CardID       string       `json:"cardId"`
	FeedSpecific FeedSpecific `json:"feedSpecific"`
	Target       Target       `json:"target"`
	AttachedInfo string       `json:"attachedInfo"`
}

// 获取知乎hot排行html并转换成json数据
func ZhihuHot2Json(cookie string) ([]ZhihuItem, error) {
	url := "https://www.zhihu.com/hot"
	s := Fetch(url, cookie)

	reg := regexp.MustCompile(`"hotList":\[{.*}\]`)
	all := reg.FindAll([]byte(s), -1)

	if len(all) < 1 {
		return nil, errors.New("cookie is invalid")
	}

	// 替换所有url中的\u002F
	hotlist := HotList{}

	for _, item := range all {
		hotJson := strings.Replace(string(item), `\u002F`, "/", -1)
		hotJson = "{" + hotJson + "}"
		//转换为json对象
		json.Unmarshal([]byte(hotJson), &hotlist)
		break
	}

	result := []ZhihuItem{}
	for _, item := range hotlist.HotList {

		// 热度处理：string->num
		hotNumStrs := strings.SplitN(item.Target.MetricsArea.Text, " ", 2)
		hotsocre, _ := strconv.Atoi(hotNumStrs[0])

		result = append(result, ZhihuItem{
			Status:   0,
			Title:    item.Target.TitleArea.Text,
			URL:      item.Target.Link.URL,
			Imgurl:   item.Target.ImageArea.URL,
			Hotscore: int32(hotsocre),
			Remark:   "",
		})
	}
	return result, nil
}

// 测试用：本地json，转为json对象，开发过程中避免多次请求知乎网站
func ZhuhuLoaclData2ZhihuItem() ([]ZhihuItem, error) {
	filePath := LOCAL_DATA_PATH
	bytes, err := ioutil.ReadFile(filePath)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	hotlist := HotList{}
	json.Unmarshal(bytes, &hotlist)

	result := []ZhihuItem{}
	for _, item := range hotlist.HotList {

		// 热度处理：string->num
		hotNumStrs := strings.SplitN(item.Target.MetricsArea.Text, " ", 2)
		hotsocre, _ := strconv.Atoi(hotNumStrs[0])

		result = append(result, ZhihuItem{
			Status:   0,
			Title:    item.Target.TitleArea.Text,
			URL:      item.Target.Link.URL,
			Imgurl:   item.Target.ImageArea.URL,
			Hotscore: int32(hotsocre),
			Remark:   "",
		})
	}
	return result, nil
}

// 定时保存获取的热榜items
func OnSaveZhihuItems() {
	cookie := GCF.UString("cookie", "")
	zhihuItems, _ := ZhihuHot2Json(cookie)
	fmt.Print(zhihuItems)
	OnZhihuItemsAddDao(zhihuItems)
}

// 定时保存获取的热榜items
func OnSaveZhihuItemsLoacl() {
	zhihuItems, _ := ZhuhuLoaclData2ZhihuItem()
	fmt.Print(zhihuItems)
	OnZhihuItemsAddDao(zhihuItems)
}
