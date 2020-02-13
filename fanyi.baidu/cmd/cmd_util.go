package cmd

import (
	"fanyi.baidu/baidu"
	"fanyi.baidu/baidu/user"
	"fanyi.baidu/util"
	"fmt"
	"myhello/input"
	"strings"
)

// 配置文件限定为当前目录下的文件
var configFileName string = util.GetCurrFileDir() + "/" + util.FANYI_BAIDU_CONFIG_FILE

var logFileName string = util.GetCurrFileDir() + "/" + util.FANYI_BAIDU_LOG_FILE

func MakeBaiduFanyiInvokeUrl() {

	// 从配置文件读取用户信息
	userInfo, ok := getUserInfoFromConf()

	// 是否需要从键盘获取用户信息（1. 没有配置文件 2. 需要更新）
	ifGetUserInfoFromInputFlag := !ok || canYouUpdate()
	if ifGetUserInfoFromInputFlag {
		// 从键盘读取信息
		userInfo = getUserInfoFromInputV3()
		// 写入到文件
		saveUserInfo(configFileName, userInfo)
	}

	q, from, to := getWordInfoFromInput()
	input := baidu.NewBaiduFanyiInput(q, from, to, userInfo)
	genUrl := input.GenQueryUrl()

	//写入到文件
	util.AppendContent(logFileName, genUrl)
	//输出到控制台
	fmt.Println(genUrl)
}

// 从键盘获取单词信息（单词，源语言，目标语言）
func getWordInfoFromInput() (string, string, string){
	//var q, from, to string
	//fmt.Print("请输入单词(q): ")
	//fmt.Scanf("%s", &q)
	//fmt.Print("请输入源语言(from): ")
	//fmt.Scanf("%s", &from)
	//fmt.Print("请输入目标语言(to): ")
	//fmt.Scanf("%s", &to)
	q := input.ReadStdin("请输入单词(q)")
	from := input.ReadStdin("请输入源语言(from)")
	to := input.ReadStdin("请输入目标语言(to)")
	return q, from, to
}

// 是否更新信息
func canYouUpdate() bool {
	//var inputStr string
	flag := false
	//fmt.Print("是否更新信息(Y/N): ")
	//fmt.Scanf("%s", &inputStr)
	inputStr := input.ReadStdin("是否更新信息(Y/N)")
	inputStr = strings.ToUpper(inputStr)
	if inputStr == "Y" {
		flag = true
	}
	if inputStr == "N" {
		flag = false
	}
	return flag
}

// 保存信息到文件
func saveUserInfo(filename string, userInfo user.UserInfo) {
	var userInfoLines []string

	userInfoLines = append(userInfoLines, fmt.Sprintf("%s:%s", user.APP_ID, userInfo.AppId))
	userInfoLines = append(userInfoLines, fmt.Sprintf("%s:%s", user.SECRET_KEY, userInfo.SecretKey))

	// 写入到文件
	util.WriteLines(filename, userInfoLines)
}


// 从配置文件中获取用户信息
// 从当前目录下的""文件读取配置信息
func getUserInfoFromConf() (user.UserInfo, bool) {

	lines := util.ReadLines(configFileName)

	var userInfo user.UserInfo
	ok := true

	if lines != nil {
		for _, line := range lines {
			arr := strings.Split(line, ":")
			if strings.TrimSpace(arr[0]) == user.APP_ID {
				userInfo.AppId = strings.TrimSpace(arr[1])
			}
			if strings.TrimSpace(arr[0]) == user.SECRET_KEY {
				userInfo.SecretKey = strings.TrimSpace(arr[1])
			}
		}
	} else {
		ok = false
	}

	return userInfo, ok
}

// -----------------------------
// 从键盘输入获取用户信息（appId, secretKey）
func getUserInfoFromInput() (string, string) {
	//var appId, secretKey string
	//fmt.Print("请输入app_id: ")
	//fmt.Scanf("%s", &appId)
	//fmt.Print("请输入secret_key: ")
	//fmt.Scanf("%s", &secretKey)
	appId := input.ReadStdin("请输入app_id")
	secretKey := input.ReadStdin("请输入secret_key")
	return appId, secretKey
}

// 返回切片
func getUserInfoFromInputV2() []string {
	appId, secretKey := getUserInfoFromInput()
	return []string{appId, secretKey}
}

// 返回结构体
func getUserInfoFromInputV3() user.UserInfo {
	appId, secretKey := getUserInfoFromInput()
	return user.UserInfo{appId, secretKey}
}




