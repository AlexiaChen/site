package site

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/go-resty/resty/v2"
)

// SendSiteMessage 发送站内信
func (s *Site) SendSiteMessage(title, info string, userId uint) (string, error) {

	if userId == 0 || s.UserId == 0 || strings.TrimSpace(s.UserName) == "" {
		return "", fmt.Errorf("发送站内信失败，参数错误: userid=%d & %d, username=%s", userId, s.UserId, s.UserName)
	}

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

	if err != nil {
		return "", fmt.Errorf("发送站内信失败: %s statue code %d  Response: %s  ", err.Error(), resp.StatusCode(), resp.String())
	}

	return resp.String(), nil
}
