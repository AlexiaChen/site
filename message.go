package site

import (
	"fmt"
	"sort"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/util/guid"
)

func sorts(text string) string {
	var array []string
	for _, v := range text {
		array = append(array, string(v))
	}
	sort.Strings(array)
	newText := ""
	for _, v := range array {
		newText += v
	}
	return newText
}

// SendSiteMessage 发送站内信
func (s *Site) SendSiteMessage(title, info, userId string) (string, error) {
	times := time.Now().Unix()
	randStr := guid.S()
	text := fmt.Sprintf("%d%s%s", times, randStr, s.APISignSecret)
	newText := sorts(text)
	ciphertext := gmd5.MustEncryptString(newText)
	data := map[string]string{
		"time_stamp": fmt.Sprintf("%d", times),
		"nonce_str":  randStr,
		"sign":       ciphertext,
		"title":      title,
		"info":       info,
		"userid":     userId,
	}
	client := resty.New()
	url := s.APIUriPrefix + MessageAPI
	resp, err := client.R().SetBody(data).Post(url)
	if err != nil {
		return "", err
	}
	return resp.String(), nil
}
