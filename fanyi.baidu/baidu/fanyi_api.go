package baidu

import (
	"fanyi.baidu/baidu/user"
	"fanyi.baidu/util"
	"fmt"
	"math/rand"
	"net/url"
	"time"
)

const BAIDU_FANYI_PREFIX string = "http://api.fanyi.baidu.com/api/trans/vip/translate"

type BaiduFanyiInput struct {
	// Query+From+To+AppId+Salt+Sign
	Query string // 请求翻译query，UTF8编码
	From  string // 源语言，可设置为auto
	To    string // 目标语言，不可设置为auto
	AppId string //
	Salt  string // 随机数
	Sign  string // 签名 : md5(AppId+Query+Salt+SecretKey)
	//Tts string // 是否显示语音合成资源，0-显示，1-不显示
	//Dict string // 是否显示词典资源，0-显示，1-不显示
	//Action string // 是否需要使用自定义术语干预API，1-是，0-否（仅开通“我的术语库”用户需填写）
}

func NewBaiduFanyiInput(q string, from string, to string, info user.UserInfo) BaiduFanyiInput {
	rand.Seed(time.Now().Unix())
	n := rand.Intn(10000)*1000 + rand.Intn(1000)

	salt := fmt.Sprintf("%d", n)
	input := BaiduFanyiInput{q, from, to, "",
		salt, ""}

	input.GenSign(info)
	return input
}

// 生成签名
func (b *BaiduFanyiInput) GenSign(info user.UserInfo) string {
	b.AppId = info.AppId
	originStr := info.AppId + b.Query + b.Salt + info.SecretKey
	sign := util.GetMd5Str(originStr)
	b.Sign = sign
	return sign
}

// 组装请求信息(query)
func (b *BaiduFanyiInput) GenQueryStr() string {
	query := url.Values{}
	// Query+From+To+AppId+Salt+Sign
	// http://api.fanyi.baidu.com/api/trans/vip/translate?q=apple&from=en&to=zh&appid=2015063000000001&salt=1435660288&sign=f89f9594663708c1605f3d736d01d2d4
	query.Add("q", b.Query)
	query.Add("from", b.From)
	query.Add("to", b.To)
	query.Add("appid", b.AppId)
	query.Add("salt", b.Salt)
	query.Add("sign", b.Sign)
	return query.Encode()
}

// url
func (b *BaiduFanyiInput) GenQueryUrl() string {

	return BAIDU_FANYI_PREFIX + "?" + b.GenQueryStr()
}
