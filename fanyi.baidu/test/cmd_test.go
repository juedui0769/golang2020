package test

import (
	"fmt"
	"testing"
)

// 在测试代码中不能读取键盘输入, 下面的代码没有效果！
func TestCmd(t *testing.T) {
	appId := ""
	secretKey := ""
	fmt.Print("请输入app_id: ")
	fmt.Scanf("%s", &appId)
	fmt.Print("请输入secret_key: ")
	fmt.Scanf("%s", &secretKey)
	t.Logf("appId: %s\n", appId)
	t.Logf("secretKey: %s\n", secretKey)
}

