package oop

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employee struct {
	Id   string
	Name string
	Age  int
}

func (e Employee) String() string {
	fmt.Printf("Memory address(Employee.Name) is %x\n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("Id: %s, Name: %s, Age: %d", e.Id, e.Name, e.Age)
}

func Test01(t *testing.T) {
	e := Employee{"1001", "Bob", 20}
	t.Log(e)
	e1 := Employee{Name:"Mike", Age:30}
	t.Log(e1)
	e2 := new(Employee)
	e2.Id = "2"
	e2.Name="小明"
	e2.Age=35
	t.Log(e2)
}

func Test02(t *testing.T) {
	e := Employee{"0", "Bob", 30}
	fmt.Printf("Memory address(Employee.Name) is %x\n", unsafe.Pointer(&e.Name))
	t.Log(e.String())
}