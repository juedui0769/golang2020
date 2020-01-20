package csp

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Microsecond * 50)
	return "1 Done"
}

func otherTask() {
	fmt.Println("2 working on something else")
	time.Sleep(time.Microsecond * 100)
	fmt.Println("2 Task is done.")
}

func TestService(t *testing.T) {
	fmt.Println(service())
	otherTask()
}

func AsyncService() chan string {
	//retCh := make(chan string)
	retCh := make(chan string, 1) // 使用这条语句将声明"buffer chan",是更高效的方式,不会导致阻塞
	go func() {
		ret := service()
		fmt.Println(">> returned result.")
		retCh <- ret
		fmt.Println(">> service exited.")
	}()
	return retCh
}

func TestAsyncService(t *testing.T) {
	retCh := AsyncService()
	otherTask()
	fmt.Println(<-retCh)
	time.Sleep(time.Second * 1)
}



