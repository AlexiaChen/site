package site

import (
	"testing"
)

func TestSite_SendSiteMessage(t *testing.T) {
	site := Site{
		UserId:       14610,
		UserName:     "86326328",
		APIUriPrefix: "https://www.st.landui.cn",
	}
	resp, err := site.SendSiteMessage("title", "msg", 14610)

	if err != nil {
		t.Errorf("SendSiteMessage error: %s", err.Error())
	}
	t.Logf("SendSiteMessage resp: %s", resp)
}
