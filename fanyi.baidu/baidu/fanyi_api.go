package baidu

import (
	"fanyi.baidu/baidu/user"
	"fanyi.baidu/util"
)

type BaiduFanyiInput struct {
	// Query+From+To+AppId+Salt+Sign
	Query string // 请求翻译query，UTF8编码
	From string // 源语言，可设置为auto
	To string // 目标语言，不可设置为auto
	AppId string //
	Salt string // 随机数
	Sign string // 签名 : md5(AppId+Query+Salt+SecretKey)
	//Tts string // 是否显示语音合成资源，0-显示，1-不显示
	//Dict string // 是否显示词典资源，0-显示，1-不显示
	//Action string // 是否需要使用自定义术语干预API，1-是，0-否（仅开通“我的术语库”用户需填写）
}

// 生成签名
func (b *BaiduFanyiInput) GenSign(info user.UserInfo) string {
	originStr := info.AppId + b.Query + b.Salt + info.SecretKey
	return util.GetMd5Str(originStr)
}


