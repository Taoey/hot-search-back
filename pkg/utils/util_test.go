package utils

import (
	"fmt"
	"testing"
)

func TestFetch(t *testing.T) {
	cookie := `pgv_pvid=7184190702; pgv_pvi=5975110656; RK=KLpt4AkINI; ptcz=83aaf825afc46783c97291675ef506410c111a4dde6f632c186cb2baf68a9eb0; eas_sid=OUr3rNcux1O4zA4o7SALeSf56l; _ga=amp-T2TwTGiKL0lu-WEt6sfrwg; tvfe_boss_uuid=53f0336b4d1c1fbe; OUTFOX_SEARCH_USER_ID_NCOO=1407654004.5319939; XWINDEXGREY=0; o_cookie=741494582; pac_uid=1_741494582; iip=0; pgv_info=ssid=s1603787992; pgv_si=s2583370752; _qpsvr_localtk=0.2935633393368695; p_uin=o0741494582; traceid=8bd96ea996; enc_uin=5OtSuKrsA7uNSQjtWun5dQ; uin=o0741494582; skey=@d1LGGCWGg; pt4_token=kvlZynJ5pVJm6-7mAci8rKEShNq002a6UuMRX9iYyuM_; p_skey=c7VkKmgnaZjkEteeNMNTXwOcxITYZHRyHAHgLlGqUCE_`
	url := "https://qun.qq.com/member.html"
	fetch := Fetch(url, cookie)
	fmt.Println(fetch)
}
