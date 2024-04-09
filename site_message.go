package site

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/util/guid"
	"gitlab.landui.cn/gomod/global"
	"gitlab.landui.cn/gomod/logs"
	"sort"
	"time"
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
func SendSiteMessage(title, info, userId string) {
	times := time.Now().Unix()
	randStr := guid.S()
	text := fmt.Sprintf("%d%s%s", times, randStr, global.Prepay)
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
	logs.New().SetAdditionalInfo("url", global.ApiUrl["send_mail"]).SetAdditionalInfo("body", data).Info("发送站内信")
	resp, err := client.R().SetBody(data).Post(global.ApiUrl["send_mail"])
	if err != nil {
		fmt.Println("发送错误")
	}
	fmt.Println("发送站内信结果", resp.String())
}
