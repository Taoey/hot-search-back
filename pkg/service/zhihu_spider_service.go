package service

import (
	"encoding/json"
	"errors"
	"fmt"
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
	cookie := `_zap=c4c42a73-9ee4-4296-b3fe-7c99ced99e43; d_c0="AMCcopaSqxGPTiO7_7jWnVag9tR7wiIWbXg=|1596351819"; _ga=GA1.2.1947722025.1596351825; _xsrf=UUajsqu9b7d0cbRMLyx21hviuMOwToWZ; capsion_ticket="2|1:0|10:1599884140|14:capsion_ticket|44:ZGM3YWFmNTRiNDM2NDlmYWI2NDAyYTE4MDNmOThlMDA=|e74e28a69f9b5abd6266efbba27b8f65eb5563303ec08a00a8e503ddc89da2f7"; z_c0="2|1:0|10:1599884168|4:z_c0|92:Mi4xb3hHYkFnQUFBQUFBd0p5aWxwS3JFU1lBQUFCZ0FsVk5pSmxKWUFCQ3Zyb0RmeGF1aE0yV2RVU1FuNUthekxka0RB|2faa8ced12762ce373771165897e26db9adf8814b78e8652f148fdfdd088d508"; tst=h; tshl=; q_c1=7b92f9cb3b474f7f9db616e9b54fa0f9|1601439070000|1596882330000; _gid=GA1.2.1767664365.1601692030; SESSIONID=6NldmcoBOAGuf6FmUPBVOpiIkEMk5buet5Ku55Kbrco; JOID=U1gUBU_OnZRwAw9CDcmZzi9_oO0c-dL_FmttLEiE3dICZzwFYcF25igHD0QIqDlsoWIHIxAKjjhn70KZG3lmu0Q=; osd=VVwTBEvImZNxBwlGCsidyCt4oeka_dX-Em1pK0mA29YFZjgDZcZ34i4DCEUMrj1roGYBJxcLij5j6EOdHX1hukA=; Hm_lvt_98beee57fd2ef70ccdd5ca52b9740c49=1601692031,1601699537,1601720145,1601774312; Hm_lpvt_98beee57fd2ef70ccdd5ca52b9740c49=1601774312; KLBRSID=2177cbf908056c6654e972f5ddc96dc2|1601783383|1601783383; _gat_gtag_UA_149949619_1=1`
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
