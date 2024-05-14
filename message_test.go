package site

import (
	"testing"

	"gitlab.landui.cn/gomod/global"
	"go.uber.org/zap"
)

func TestSite_SendSiteMessage(t *testing.T) {
	initLog()
	site := Site{
		UserId:       14610,
		UserName:     "86326328",
		APIUriPrefix: "https://www.st.landui.cn",
	}
	_, _ = site.SendSiteMessage("title", "msg", 14610)
}

func initLog() {
	global.Logger = zap.NewExample() // 需要传入 zap.AddCaller() 才会显示打日志点的文件名和行数
}
