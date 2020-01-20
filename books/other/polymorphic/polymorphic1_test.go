package polymorphic

import (
	"fmt"
	"testing"
)

type Code string

type Programmer interface {
	WriteHelloWorld() Code
}

type GoProgrammer struct {}

func (p *GoProgrammer) WriteHelloWorld() Code {
	return "fmt.Println(\"Hello World!\")"
}

type JavaProgrammer struct {}

func (p *JavaProgrammer) WriteHelloWorld() Code {
	return "System.out.println(\"Hello World!\")"
}

func writeFirstProgram(p Programmer) {
	fmt.Printf("%T %v\n", p, p.WriteHelloWorld())
}

func TestPolymorphic01(t *testing.T) {
	goProgrammer := new(GoProgrammer)
	javaProgrammer := new(JavaProgrammer)
	writeFirstProgram(goProgrammer)
	writeFirstProgram(javaProgrammer)
}

