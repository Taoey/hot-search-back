package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"testing"
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

func fetch(url, cookie string) string {
	fmt.Println("Fetch Url", url)
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	//req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36")
	req.Header.Set("cookie", cookie)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Http get err:", err)
		return ""
	}
	if resp.StatusCode != 200 {
		fmt.Println("Http status code:", resp.StatusCode)
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read error", err)
		return ""
	}
	return string(body)
}

// 获取知乎hot排行html并转换成json数据
func ZhihuHot2Json(zhCookie string) (interface{}, error) {
	url := "https://www.zhihu.com/hot"
	cookie := `_zap=71f3d34d-74bf-4215-b1ab-4c6402cbde8d; d_c0="AGBcgHSQDxGPTgen4xICn3YyiLymnHOebBM=|1585882277"; _ga=GA1.2.557933148.1585882278; tst=h; tshl=; _xsrf=b5GK3lXhic52y2iDT3Ma5m8njt8ef6HC; __utma=155987696.557933148.1585882278.1589207068.1589207068.1; __utmz=155987696.1589207068.1.1.utmcsr=(direct)|utmccn=(direct)|utmcmd=(none); q_c1=05942608f1604731b10ef507a75ae8d4|1591852293000|1585968770000; _gid=GA1.2.1247412657.1593233797; Hm_lvt_98beee57fd2ef70ccdd5ca52b9740c49=1593139711,1593140049,1593233796,1593242424; SESSIONID=bLK64z8KM1I3lLrsIrIVmTaVN8sxXEBHSdeskb7Ctfb; JOID=V1gdBEouG4SF3JuEWCwKXq8UpPhESV_AtezktBxnL8W5ttPzI-_ve9ndloNR1d3FSLBs70wm5WhsSx0RQXcK2PQ=; osd=W1wRA0giH4iC3peAVCsIUqsYo_pITVPHt-DguBtlI8G1sdH_J-PoedXZmoRT2dnJT7Jg60Ah52RoRxoTTXMG3_Y=; capsion_ticket="2|1:0|10:1593242611|14:capsion_ticket|44:N2RkOGJiNzBmNTg4NDFjMjk4ODEwZmVjZmY4OWI2OTQ=|1c10185d13d1f6683058c8cbadf747dd34b64c456956b93e6e49417a4ee718a2"; z_c0="2|1:0|10:1593242673|4:z_c0|92:Mi4xb3hHYkFnQUFBQUFBWUZ5QWRKQVBFU2NBQUFDRUFsVk5NSUVlWHdEZGxlY3hmcFh1UV9WS2p4OUlTSTB5R3ZKUV9B|94f20d7c1226e8f533e217857c4cfe128ef986ea4918b6e47a93ae1e7e7e46f2"; Hm_lpvt_98beee57fd2ef70ccdd5ca52b9740c49=1593242771; _gat_gtag_UA_149949619_1=1; KLBRSID=fe0fceb358d671fa6cc33898c8c48b48|1593242777|1593241801`
	s := fetch(url, cookie)

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
	return hotlist, nil
}

// 测试用：本地json，转为json对象，开发过程中多次请求知乎网站
func ZhuhuLoaclData2Json() (interface{}, error) {
	filePath := "data.json"
	bytes, err := ioutil.ReadFile(filePath)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	hotlist := HotList{}
	json.Unmarshal(bytes, &hotlist)

	return hotlist, nil
}

func Test01(t *testing.T) {
	data2Json, e := ZhuhuLoaclData2Json()
	fmt.Println(data2Json, e)
}
