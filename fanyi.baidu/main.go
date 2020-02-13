// Copyright 2013 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fanyi.baidu/cmd"
	"fmt"
)

func main() {
	cmd.MakeBaiduFanyiInvokeUrl()
}




func testCmdInputV1() {
	appId := ""
	secretKey := ""
	fmt.Print("请输入app_id: ")
	fmt.Scanf("%s", &appId)
	fmt.Print("请输入secret_key: ")
	fmt.Scanf("%s", &secretKey)
	fmt.Printf("\n----------\nappId: %s, secretKey: %s", appId, secretKey)
}

// 错误的写法
func testCmdInputV0() {
	//appId := ""
	//secretKey := ""
	//fmt.Scanf("请输入app_id: %s", &appId)
	//fmt.Scanf("请输入secret_key: %s", &secretKey)
	//fmt.Printf("\n----------\nappId: %s, secretKey: %s", appId, secretKey)
}
