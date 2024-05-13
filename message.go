package site

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"gitlab.landui.cn/gomod/logs"
	"strconv"
)

// SendSiteMessage 发送站内信
func (s *Site) SendSiteMessage(title, info string, userId uint) error {
	header := map[string]string{
		"User-Agent":  "api-landui-lan",
		"X-RequestAU": fmt.Sprintf("%d|\t|%s", s.UserId, s.UserName),
	}
	data := map[string]string{
		"title":  title,
		"info":   info,
		"userid": strconv.Itoa(int(userId)),
	}
	client := resty.New()
	url := s.APIUriPrefix + MessageAPI
	resp, err := client.R().SetHeaders(header).SetBody(data).Post(url)
	logs.New().
		SetAdditionalInfo("url", url).
		SetAdditionalInfo("body", data).
		SetAdditionalInfo("response", resp).
		Debug("发送站内信")
	return err
}
