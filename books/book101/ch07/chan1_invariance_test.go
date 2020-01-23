package ch07

import (
	"fmt"
	"testing"
)

// 通道中的元素值丝毫不受外界的影响
// 这说明了，在发送过程中进行的元素值复制并非浅表复制，而属于完全复制。
// 这也保证了我们使用通道传递的值的不变性

type Person struct {
	Name    string
	Age     uint8
	Address Addr
}

type Addr struct {
	city     string
	district string // 区域
}

func TestChanInvariance01(t *testing.T) {
	var personChan = make(chan Person, 1)

	p1 := Person{"Harry", 32, Addr{"北京", "海淀区"}}
	fmt.Printf("p1 (1): %v\n", p1)
	personChan <- p1

	p1.Address.district = "石景山区"
	fmt.Printf("p1 (2): %v\n", p1)

	p1_copy := <-personChan
	fmt.Printf("p1_copy: %v\n", p1_copy)

}
