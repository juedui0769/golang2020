package util

import (
	"fanyi.baidu/baidu"
	"fmt"
	"strings"
)

// 配置文件限定为当前目录下的文件
var filename string = GetCurrFileDir() + "/" + FANYI_BAIDU_CONFIG_FILE

func MakeBaiduFanyiInvokeUrl() {

	// 从配置文件读取用户信息
	userInfo, ok := getUserInfoFromConf()

	// 是否需要从键盘获取用户信息（1. 没有配置文件 2. 需要更新）
	ifGetUserInfoFromInputFlag := !ok || canYouUpdate()
	if ifGetUserInfoFromInputFlag {
		// 从键盘读取信息
		userInfo = getUserInfoFromInputV3()
		// 写入到文件
		saveUserInfo(filename, userInfo)
	}



	//userInfo := getUserInfoFromInputV3()
	//fmt.Printf("\n----------\nappId: %s, secretKey: %s", userInfo.AppId, userInfo.SecretKey)
}



// 是否更新信息
func canYouUpdate() bool {
	var inputStr string
	flag := false
	fmt.Print("是否更新信息(Y/N): ")
	fmt.Scanf("%s", &inputStr)
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
func saveUserInfo(filename string, userInfo baidu.UserInfo) {
	var userInfoLines []string

	userInfoLines = append(userInfoLines, fmt.Sprintf("%s:%s", baidu.APP_ID, userInfo.AppId))
	userInfoLines = append(userInfoLines, fmt.Sprintf("%s:%s", baidu.SECRET_KEY, userInfo.SecretKey))

	// 写入到文件
	WriteLines(filename, userInfoLines)
}


// 从配置文件中获取用户信息
// 从当前目录下的""文件读取配置信息
func getUserInfoFromConf() (baidu.UserInfo, bool) {

	lines := ReadLines(filename)

	var userInfo baidu.UserInfo
	ok := true

	if lines != nil {
		for _, line := range lines {
			arr := strings.Split(line, ":")
			if strings.TrimSpace(arr[0]) == baidu.APP_ID {
				userInfo.AppId = strings.TrimSpace(arr[1])
			}
			if strings.TrimSpace(arr[0]) == baidu.SECRET_KEY {
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
	var appId, secretKey string
	fmt.Print("请输入app_id: ")
	fmt.Scanf("%s", &appId)
	fmt.Print("请输入secret_key: ")
	fmt.Scanf("%s", &secretKey)
	return appId, secretKey
}

// 返回切片
func getUserInfoFromInputV2() []string {
	appId, secretKey := getUserInfoFromInput()
	return []string{appId, secretKey}
}

// 返回结构体
func getUserInfoFromInputV3() baidu.UserInfo {
	appId, secretKey := getUserInfoFromInput()
	return baidu.UserInfo{appId, secretKey}
}




